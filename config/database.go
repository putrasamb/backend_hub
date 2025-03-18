package config

import (
	"backend_hub/internal/infrastructure/database"

	"github.com/spf13/viper"
)

func setDefaultDatabaseConfig(viper *viper.Viper) {
	viper.SetDefault(database.DB_LOG_LEVEL, "info")
	viper.SetDefault(database.DB_MAX_OPEN_CONNS, 10)
	viper.SetDefault(database.DB_MAX_IDLE_CONNS, 10)
	viper.SetDefault(database.DB_CONN_MAX_LIFETIME, 10)
	viper.SetDefault(database.DB_CONN_MAX_IDLE_TIME, 10)
}
