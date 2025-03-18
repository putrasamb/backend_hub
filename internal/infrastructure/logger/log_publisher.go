package logger

import (
	"time"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CustomLog interface {
	Json() ([]byte, error)
	String() (*string, error)
}

type LogPublisher[T CustomLog] struct {
	AMQPConnection *amqp.Connection
	Exchange       string // Exchange name
	Route          string // Routing/Binding key
}

// Publish publish activity log into message broker
func (l *LogPublisher[T]) Publish(log T) error {
	ch, err := l.AMQPConnection.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open amqp channel")
	}
	defer ch.Close()

	payload, err := log.Json()
	if err != nil {
		return err
	}

	if err := ch.Publish(l.Exchange, l.Route, false, false, amqp.Publishing{
		ContentType:     "application/json",
		ContentEncoding: "UTF-8",
		Timestamp:       time.Now(),
		Body:            payload,
	}); err != nil {
		return errors.Wrap(err, "failed to publish activity log")
	}

	return nil
}
