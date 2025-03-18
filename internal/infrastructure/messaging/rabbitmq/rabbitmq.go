package rabbitmq

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"backend_hub/internal/infrastructure/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitmuttex sync.Mutex

// Credential represents rabbitmq Credentials
type Credential struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// DSN represents RabbitMQ data source name
type DSN struct {
	Host string `json:"host" validate:"required"`
	Port string `json:"port" validate:"required"`
}

// Config represents rabbitmq connection configuration
type Config struct {
	Vhost                  string        `json:"vhost" `
	Heartbeat              time.Duration `json:"heartbeat"` // in Seconds
	TLSClientConfig        *tls.Config   `json:"tls_config"`
	RetryConnectInterval   time.Duration `json:"connect_retry_interval"`
	RetryConnectMaxAttempt uint8         `json:"connect_retry_max_attempt"`
	connectAttempt         uint8
}

type Connection struct {
	DSN        DSN        `json:"dsn"`
	Credential Credential `json:"credential"`
	Config     Config     `json:"config"`
}

// RabbitMQ represents RabbitMQ and simplify implementation
type RabbitMQ struct {
	Connection Connection `json:"connection" validate:"required"`
	Logger     *logger.Logger
}

// String returns DSN string
func (c *Connection) String() *string {
	d := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		c.Credential.User,
		c.Credential.Password,
		c.DSN.Host,
		c.DSN.Port)
	return &d
}

// setDefaultConfigs set default dsn configs
func (r *Connection) setDefaultConfigs() {
	if r.Config.Vhost == "" {
		r.Config.Vhost = "/"
	}

	if r.Config.RetryConnectInterval == 0 {
		r.Config.RetryConnectInterval = 3
	}

	if r.Config.RetryConnectMaxAttempt == 0 {
		r.Config.RetryConnectMaxAttempt = 3
	}

	if r.Config.Heartbeat == 0 {
		r.Config.Heartbeat = time.Duration(10)
	}
}

// Connect open connection to RabbitMQ
func (r *RabbitMQ) Connect() (*amqp.Connection, error) {
	rabbitmuttex.Lock()
	defer rabbitmuttex.Unlock()

	r.Connection.setDefaultConfigs()
	interval := r.Connection.Config.RetryConnectInterval * time.Second
	for {
		conn, err := amqp.DialConfig(*r.Connection.String(), amqp.Config{
			Vhost:           r.Connection.Config.Vhost,
			Heartbeat:       r.Connection.Config.Heartbeat * time.Second,
			TLSClientConfig: r.Connection.Config.TLSClientConfig,
		})

		if err == nil {
			r.Logger.Logger.Infof("AMQP connection to %v establised", r.Connection.DSN.Host)
			return conn, nil
		}

		r.Connection.Config.connectAttempt++
		if r.Connection.Config.connectAttempt == r.Connection.Config.RetryConnectMaxAttempt {
			return nil, fmt.Errorf("failed to connect to rabbitmq: max connection retry exceed")
		}

		r.Logger.Logger.WithError(err).Errorf(
			"[ATTEMPT %d/%d] Failed to connect to RabbitMQ, retrying in %f second(s)",
			r.Connection.Config.connectAttempt, r.Connection.Config.RetryConnectMaxAttempt,
			interval.Seconds())

		time.Sleep(interval)
	}
}
