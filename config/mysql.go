package config

import (
	"time"

	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/database/mysql"
	"backend_hub/pkg/constant"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// setMySQLDefaultConfig sets default configuration values for MySQL connections
func setMySQLDefaultConfig(v *viper.Viper) {
	v.SetDefault(mysql.MYSQL_READ_HOST, "localhost")
	v.SetDefault(mysql.MYSQL_READ_PORT, "3306")
	v.SetDefault(mysql.MYSQL_WRITE_HOST, "localhost")
	v.SetDefault(mysql.MYSQL_WRITE_PORT, "3306")
	v.SetDefault(mysql.MASTERDB_MYSQL_READ_HOST, "localhost")
	v.SetDefault(mysql.MASTERDB_MYSQL_READ_HOST, "3306")
}

// NewMySQLReadMasterDB establishes a read-only database "Master" connection
func NewMySQLReadMasterDB(v *viper.Viper, opts ...func(*gorm.DB)) (*gorm.DB, error) {
	conn := &mysql.Connection{
		Host:     v.GetString(mysql.MASTERDB_MYSQL_READ_HOST),
		Port:     v.GetString(mysql.MASTERDB_MYSQL_READ_PORT),
		User:     v.GetString(mysql.MASTERDB_MYSQL_READ_USER),
		Password: v.GetString(mysql.MASTERDB_MYSQL_READ_PASSWORD),
		Database: v.GetString(mysql.MASTERDB_MYSQL_READ_DATABASE),
		ConnectionQuery: mysql.ConnectionQuery{
			Charset:   mysql.MASTERDB_MYSQL_DEFAULT_CHARSET,
			ParseTime: true,
			Timezone:  v.GetString(constant.TIMEZONE),
			TLS:       v.GetBool(mysql.MASTERDB_MYSQL_READ_USE_TLS),
		},
		Config: mysql.Config{
			LogLevel:                v.GetString(database.DB_LOG_LEVEL),
			MaxOpenConnection:       uint8(v.GetUint(database.DB_MAX_OPEN_CONNS)),
			MaxIdleConnection:       uint8(v.GetUint(database.DB_MAX_IDLE_CONNS)),
			MaxLifetimeConnection:   uint8(v.GetUint(database.DB_CONN_MAX_LIFETIME)),
			MaxIdleTimeConnection:   v.GetDuration(database.DB_CONN_MAX_IDLE_TIME),
			RetryConnectionInterval: 3 * time.Second,
			MaxRetry:                3,
			PrepareStmt:             true,
			SkipDefaultTransaction:  true,
		},
	}
	return conn.Connect(opts...)
}

// NewMySQLReadDB establishes a read-only database connection
func NewMySQLReadDB(v *viper.Viper, opts ...func(*gorm.DB)) (*gorm.DB, error) {
	conn := &mysql.Connection{
		Host:     v.GetString(mysql.MYSQL_READ_HOST),
		Port:     v.GetString(mysql.MYSQL_READ_PORT),
		User:     v.GetString(mysql.MYSQL_READ_USER),
		Password: v.GetString(mysql.MYSQL_READ_PASSWORD),
		Database: v.GetString(mysql.MYSQL_READ_DATABASE),
		ConnectionQuery: mysql.ConnectionQuery{
			Charset:   mysql.MYSQL_DEFAULT_CHARSET,
			ParseTime: true,
			Timezone:  v.GetString(constant.TIMEZONE),
			TLS:       v.GetBool(mysql.MYSQL_READ_USE_TLS),
		},
		Config: mysql.Config{
			LogLevel:                v.GetString(database.DB_LOG_LEVEL),
			MaxOpenConnection:       uint8(v.GetUint(database.DB_MAX_OPEN_CONNS)),
			MaxIdleConnection:       uint8(v.GetUint(database.DB_MAX_IDLE_CONNS)),
			MaxLifetimeConnection:   uint8(v.GetUint(database.DB_CONN_MAX_LIFETIME)),
			MaxIdleTimeConnection:   v.GetDuration(database.DB_CONN_MAX_IDLE_TIME),
			RetryConnectionInterval: 3 * time.Second,
			MaxRetry:                3,
			PrepareStmt:             true,
			SkipDefaultTransaction:  true,
		},
	}
	return conn.Connect(opts...)
}

// NewMySQLWriteDB establishes a write-enabled database connection
func NewMySQLWriteDB(v *viper.Viper, opts ...func(*gorm.DB)) (*gorm.DB, error) {
	conn := &mysql.Connection{
		Host:     v.GetString(mysql.MYSQL_WRITE_HOST),
		Port:     v.GetString(mysql.MYSQL_WRITE_PORT),
		User:     v.GetString(mysql.MYSQL_WRITE_USER),
		Password: v.GetString(mysql.MYSQL_WRITE_PASSWORD),
		Database: v.GetString(mysql.MYSQL_WRITE_DATABASE),
		ConnectionQuery: mysql.ConnectionQuery{
			Charset:   mysql.MYSQL_DEFAULT_CHARSET,
			ParseTime: true,
			Timezone:  v.GetString(constant.TIMEZONE),
			TLS:       v.GetBool(mysql.MYSQL_READ_USE_TLS),
		},
		Config: mysql.Config{
			LogLevel:                v.GetString(database.DB_LOG_LEVEL),
			MaxOpenConnection:       uint8(v.GetUint(database.DB_MAX_OPEN_CONNS)),
			MaxIdleConnection:       uint8(v.GetUint(database.DB_MAX_IDLE_CONNS)),
			MaxLifetimeConnection:   uint8(v.GetUint(database.DB_CONN_MAX_LIFETIME)),
			MaxIdleTimeConnection:   v.GetDuration(database.DB_CONN_MAX_IDLE_TIME),
			RetryConnectionInterval: 3 * time.Second,
			MaxRetry:                3,
			PrepareStmt:             true,
			SkipDefaultTransaction:  true,
		},
	}
	return conn.Connect(opts...)
}
