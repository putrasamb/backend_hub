package controller

import (
	"backend_hub/internal/infrastructure/logger"
	"backend_hub/internal/usecase"
	httprequest "backend_hub/pkg/common/http/request"
	httpresponse "backend_hub/pkg/common/http/response"
	"net/http"

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

func (c *CollectionDocumentController) List(ctx echo.Context) error {
	request := &httprequest.ListRequest{}
	if err := ctx.Bind(request); err != nil {
		errResponse := &httpresponse.ErrorResponse{}
		errResponse.Code = http.StatusBadRequest
		errResponse.Message = "failed to parse list request body"
		errResponse.Error = err.Error()
		c.logError(err, errResponse.Message)
		return errResponse.EchoJsonResponse(ctx)
	}

	resp, err := c.CollectionDocumentUseCase.List(request)
	if err != nil {
		errResponse := &httpresponse.ErrorResponse{}
		errResponse.Code = http.StatusBadRequest
		errResponse.Message = "failed to list customer"
		errResponse.Error = err.Error()
		c.logError(err, errResponse.Message)
		return errResponse.EchoJsonResponse(ctx)
	}
	return ctx.JSON(http.StatusOK, resp)
}
