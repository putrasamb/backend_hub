package config

import (
	"crypto/tls"

	"service-collection/internal/infrastructure/logger"
	"service-collection/internal/infrastructure/messaging/rabbitmq"

	"github.com/spf13/viper"
)

// NewRabbitMQ returns rabbitmq configuration
func NewRabbitMQ(v *viper.Viper, l *logger.Logger, tlsConfig *tls.Config) *rabbitmq.RabbitMQ {
	return &rabbitmq.RabbitMQ{
		Connection: rabbitmq.Connection{
			DSN: rabbitmq.DSN{
				Host: v.GetString(rabbitmq.RABBITMQ_HOST),
				Port: v.GetString(rabbitmq.RABBITMQ_PORT),
			},
			Credential: rabbitmq.Credential{
				User:     v.GetString(rabbitmq.RABBITMQ_USER),
				Password: v.GetString(rabbitmq.RABBITMQ_PASSWORD),
			},
			Config: rabbitmq.Config{
				Vhost:                  v.GetString(rabbitmq.RABBITMQ_VHOST),
				Heartbeat:              v.GetDuration(rabbitmq.RABBITMQ_HEARTBEAT_INTERVAL),
				TLSClientConfig:        tlsConfig,
				RetryConnectInterval:   v.GetDuration(rabbitmq.RABBITMQ_RETRY_CONNECT_INTERVAL),
				RetryConnectMaxAttempt: uint8(v.GetUint(rabbitmq.RABBITMQ_MAX_RETRY_CONNECT)),
			},
		},

		Logger: l,
	}
}
