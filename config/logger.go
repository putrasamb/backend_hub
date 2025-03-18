package config

import (
	"backend_hub/internal/infrastructure/logger"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// LoggerConfig represents logger config
type LoggerConfig struct {
	Logger         *logrus.Logger
	ActivityLogger *logger.ActivityLogger
	AuditLogger    *logger.AuditLogger
	Output         string
	FileName       string
}

// NewLogger returns logger instance
func NewLogger(c *LoggerConfig) (*logger.Logger, error) {
	logger := &logger.Logger{
		Logger:         c.Logger,
		ActivityLogger: c.ActivityLogger,
		AuditLogger:    c.AuditLogger,
		Output:         c.Output,
		FileName:       c.FileName,
	}
	if err := logger.Setup(); err != nil {
		return nil, err
	}
	return logger, nil
}

// NewActivityLogger returns new activity logger instance
func NewActivityLogger(conn *amqp.Connection, exchange, route string) *logger.ActivityLogger {
	l := &logger.ActivityLogger{}
	l.AMQPConnection = conn
	l.Exchange = exchange
	l.Route = route
	return l
}

// NewAuditLogger returns new audit logger instance
func NewAuditLogger(conn *amqp.Connection, exchange, route string) *logger.AuditLogger {
	l := &logger.AuditLogger{}
	l.AMQPConnection = conn
	l.Exchange = exchange
	l.Route = route
	return l
}
