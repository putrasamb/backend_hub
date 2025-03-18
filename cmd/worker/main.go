package main

import (
	"backend_hub/config"
	"backend_hub/internal/infrastructure/scheduler"
	"backend_hub/pkg/constant"
	"context"
	"os"
	"os/signal"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize configuration
	viperConfig, err := config.NewViper(".env", "env", ".")
	if err != nil {
		log.Warnf("Configuration load error: %v", err)
	}

	// Setup logging
	loggerConfig := &config.LoggerConfig{
		Logger:   logrus.New(),
		Output:   "./public/logs",
		FileName: "scheduler",
	}

	logging, err := config.NewLogger(loggerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "scheduler logger initialization failed"))
	}

	schedulerLogger := &scheduler.SchedulerLogger{
		Logger: logging,
	}

	schedulerOptions := &scheduler.SchedulerOption{
		Timezone: viperConfig.GetString(constant.TIMEZONE),
		Logger:   schedulerLogger,
	}

	scheduler, err := scheduler.NewScheduler(schedulerOptions)
	if err != nil {
		log.Fatalf("cannot initialize scheduler: %v", err)
	}

	config := &config.BootstrapWorkerConfig{
		Scheduler: scheduler,
		Config:    viperConfig,
		Logger:    logging,
	}
	if err := config.Bootstrap(); err != nil {
		log.Fatal(errors.Wrap(err, "worker bootstrap failed"))
	}

	// manage worker lifecycle
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		scheduler.Start()
	}()

	// Wait for shutdown signal
	<-ctx.Done()

	if err := scheduler.Shutdown(); err != nil {
		logging.Logger.Fatalf("error while shutting down the scheduler: %v", err)
	}
}
