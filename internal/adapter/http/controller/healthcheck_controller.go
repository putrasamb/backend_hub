package controller

import (
	"backend_hub/internal/infrastructure/logger"
	"backend_hub/internal/usecase"

	"github.com/labstack/echo/v4"
)

// HealthCheckController represents healthcheck controller
type HealthCheckController struct {
	UseCase *usecase.HealthCheckUseCase
	Logger  *logger.Logger
}

// HealthCheckController instantiate new healthcheck controller
func NewHealthCheckController(l *logger.Logger, c *usecase.HealthCheckUseCase) *HealthCheckController {
	return &HealthCheckController{
		UseCase: c,
		Logger:  l,
	}
}

func (ctrl *HealthCheckController) Ping(c echo.Context) error {
	health := ctrl.UseCase.Ping()
	return c.JSON(health.Code, health)
}
