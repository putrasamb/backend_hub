package config

import (
	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/logger"

	"github.com/go-co-op/gocron/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapWorkerConfig struct {
	Config    *viper.Viper
	Scheduler gocron.Scheduler
	DB        *database.Kind[*gorm.DB]
	Logger    *logger.Logger
}

// Bootstrap your worker here
func (c *BootstrapWorkerConfig) Bootstrap() error {

	c.Scheduler.NewJob(
		gocron.CronJob("* * * * * *", true),
		gocron.NewTask(func() {
			c.Logger.Logger.Info("this crontab job is running every second")
		}),
	)
	return nil
}
