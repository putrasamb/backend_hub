package mysql

import (
	util "backend_hub/pkg/common/util/formatter"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectionQuery represents db connection query string
type ConnectionQuery struct {
	Charset   string `query:"charset"`
	ParseTime bool   `query:"parseTime"`
	Timezone  string `query:"loc"`
	TLS       bool   `query:"tls"`
}

// Config represents database connection and initialization configs
type Config struct {
	LogLevel                string
	MaxOpenConnection       uint8
	MaxIdleConnection       uint8
	MaxLifetimeConnection   uint8
	MaxIdleTimeConnection   time.Duration
	RetryConnectionInterval time.Duration
	MaxRetry                uint8
	PrepareStmt             bool
	SkipDefaultTransaction  bool
}

// Connection represents MySQL DSN and configs
type Connection struct {
	Host            string
	Port            string
	User            string
	Password        string
	Database        string
	ConnectionQuery ConnectionQuery
	Config          Config
	retry           uint8
	db              *gorm.DB
	sqlDB           *sql.DB
}

// Query represents ConnectionQuery query strings.
func (cq *ConnectionQuery) Query() *string {
	return util.StructToQuery(cq)
}

// String returns connection string.
func (c *Connection) String() string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	query := c.ConnectionQuery.Query()

	if *query != "" {
		connection = fmt.Sprintf("%s?%s", connection, *query)
	}
	return connection
}

// GetGormLogLevel returns gorm.Logger based on the given string literal.
func (c *Config) GetGormLogLevel(loglevel *string) *logger.LogLevel {

	var logLevel logger.LogLevel

	switch *loglevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		fallthrough
	default:
		logLevel = logger.Info
	}

	return &logLevel
}

// GetRetryInterval returns retry interval after failed connections
func (c *Connection) GetRetryInterval() time.Duration {
	if c.Config.RetryConnectionInterval > 0 {
		return c.Config.RetryConnectionInterval * time.Second
	}
	return time.Duration(5) * time.Second
}

// Connect is a robust connect to mysql server
func (c *Connection) Connect(options ...func(db *gorm.DB)) (*gorm.DB, error) {
	c.retry = 0

	if c.Config.MaxRetry == 0 {
		c.Config.MaxRetry = 3
	}

	var err error

	c.db, err = reconnect(c)

	if err != nil {
		return nil, err
	}

	c.sqlDB, err = c.db.DB()

	if err != nil {
		return nil, err
	}

	// connection pool configuration

	// SetMaxOpenConns sets the maximum number of open connections to the database
	if c.Config.MaxOpenConnection > 0 {
		c.sqlDB.SetMaxOpenConns(int(c.Config.MaxOpenConnection))
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	if c.Config.MaxIdleConnection > 0 {
		c.sqlDB.SetMaxIdleConns(int(c.Config.MaxIdleConnection))
	}

	if c.Config.MaxLifetimeConnection > 0 {
		c.sqlDB.SetConnMaxLifetime(time.Duration(c.Config.MaxLifetimeConnection) * time.Second)
	}

	if c.Config.MaxIdleTimeConnection > 0 {
		c.sqlDB.SetConnMaxIdleTime(time.Duration(c.Config.MaxIdleTimeConnection) * time.Second)
	}

	for _, option := range options {
		option(c.db)
	}

	return c.db, nil
}

// Stats returns sql.DBStats
func (c *Connection) Stats() *sql.DBStats {
	stats := c.sqlDB.Stats()
	return &stats
}

// Ping is pong
func (c *Connection) Ping() error {
	return c.sqlDB.Ping()
}

// connect connect to mysql server using gorm
func connect(c *Connection) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.String()), &gorm.Config{
		PrepareStmt:            c.Config.PrepareStmt,
		SkipDefaultTransaction: c.Config.SkipDefaultTransaction,
		Logger: logger.Default.LogMode(
			*(c.Config.GetGormLogLevel(&c.Config.LogLevel))),
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}

// reconnect up to n retries of failed connection.
func reconnect(c *Connection) (*gorm.DB, error) {
	interval := c.GetRetryInterval()

	for {
		if c.retry >= c.Config.MaxRetry {
			break
		}
		db, err := connect(c)
		if err == nil {
			return db, nil
		}
		c.retry++
		time.Sleep(interval)
	}

	return nil, fmt.Errorf("cannot connect to mysql server")
}
