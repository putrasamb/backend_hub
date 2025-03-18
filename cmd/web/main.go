package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"service-collection/config"
	"service-collection/internal/adapter/validator"
	"service-collection/internal/infrastructure/database"
	"service-collection/pkg/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
		FileName: "logs",
	}

	logging, err := config.NewLogger(loggerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Logger initialization failed"))
	}

	// Connect to message broker
	rabbitmq := config.NewRabbitMQ(viperConfig, logging, &tls.Config{})
	amqpConnection, err := rabbitmq.Connect()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Message broker connection failed"))
	}

	// Configure logging channels
	logging.ActivityLogger = config.NewActivityLogger(amqpConnection, "", "activity_logs")
	logging.AuditLogger = config.NewAuditLogger(amqpConnection, "", "audit_logs")

	// Initialize components
	validate := validator.NewValidator()

	// Connect to read and write databases
	dbRead, err := config.NewMySQLReadDB(viperConfig)
	if err != nil {
		log.Fatalf("Read database connection error: %v", err)
	}

	dbWrite, err := config.NewMySQLWriteDB(viperConfig)
	if err != nil {
		log.Fatalf("Write database connection error: %v", err)
	}
	// Create web server
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// Prepare bootstrap configuration
	bootstrapConfig := &config.BootstrapConfig{
		DB: &database.Kind[*gorm.DB]{
			Read:  dbRead,
			Write: dbWrite,
		},
		App:       e,
		Logger:    logging,
		Validator: validate,
		Config:    viperConfig,
	}

	// Bootstrap application
	if err = config.Bootstrap(bootstrapConfig); err != nil {
		log.Fatal(errors.Wrap(err, "Application bootstrap failed"))
	}

	// Manage server lifecycle
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	port := viperConfig.GetInt(constant.PORT)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Server startup error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		viperConfig.GetDuration(constant.TIMEOUT_GRACEFUL_SHUTDOWN)*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		e.Logger.Fatalf("Server shutdown error: %v", err)
	}
}
