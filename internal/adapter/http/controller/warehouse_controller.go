package controller

import (
	"errors"
	"net/http"

	dto "backend_hub/internal/adapter/dto/request"
	"backend_hub/internal/infrastructure/logger"
	"backend_hub/internal/usecase"
	httprequest "backend_hub/pkg/common/http/request"
	httpresponse "backend_hub/pkg/common/http/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WarehouseController struct {
	Logger  *logger.Logger
	UseCase *usecase.WarehouseUseCase
}

func NewWarehouseController(l *logger.Logger, u *usecase.WarehouseUseCase) *WarehouseController {
	return &WarehouseController{
		Logger:  l,
		UseCase: u,
	}
}

func (ctrl *WarehouseController) List(c echo.Context) error {
	request := &httprequest.ListRequest{}
	if err := c.Bind(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusBadRequest, "failed to parse list request", err).EchoJsonResponse(c)
	}

	resp, err := ctrl.UseCase.List(request)
	if err != nil {
		return httpresponse.NewErrorResponse(http.StatusInternalServerError, "failed to list products", err).EchoJsonResponse(c)
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl *WarehouseController) Get(c echo.Context) error {
	request := &dto.GetProductRequest{}
	if err := c.Bind(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusBadRequest, "failed to parse get request", err).EchoJsonResponse(c)
	}

	product, err := ctrl.UseCase.Get(request)
	if err != nil {
		code := http.StatusInternalServerError
		message := "failed to get product"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "NOT FOUND"
		}
		return httpresponse.NewErrorResponse(code, message, err).EchoJsonResponse(c)
	}

	return c.JSON(http.StatusOK, product)
}

func (ctrl *WarehouseController) Create(c echo.Context) error {
	request := &dto.CreateProductRequest{}
	if err := c.Bind(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusBadRequest, "failed to parse create request", err).EchoJsonResponse(c)
	}

	resp, err := ctrl.UseCase.Create(request)
	if err != nil {
		return httpresponse.NewErrorResponse(http.StatusInternalServerError, "failed to create product", err).EchoJsonResponse(c)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (ctrl *WarehouseController) Update(c echo.Context) error {
	request := &dto.UpdateProductRequest{}
	if err := c.Bind(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusBadRequest, "failed to parse update request", err).EchoJsonResponse(c)
	}

	resp, err := ctrl.UseCase.Update(request)
	if err != nil {
		return httpresponse.NewErrorResponse(http.StatusInternalServerError, "failed to update product", err).EchoJsonResponse(c)
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctrl *WarehouseController) Delete(c echo.Context) error {
	request := &dto.DeleteProductRequest{}
	if err := c.Bind(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusBadRequest, "failed to parse delete request", err).EchoJsonResponse(c)
	}

	if err := ctrl.UseCase.Delete(request); err != nil {
		return httpresponse.NewErrorResponse(http.StatusInternalServerError, "failed to delete product", err).EchoJsonResponse(c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "product deleted successfully"})
}
