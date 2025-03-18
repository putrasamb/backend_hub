package controller

import (
	"net/http"
	"service-collection/internal/adapter/validator"
	"service-collection/internal/infrastructure/logger"
	"service-collection/internal/usecase"
	httpresponse "service-collection/pkg/common/http/response"

	"github.com/labstack/echo/v4"
)

type TesController struct {
	Logger     *logger.Logger
	Validator  *validator.CustomValidator
	TesUseCase *usecase.TesUseCase
}

func NewTesController(
	l *logger.Logger,
	v *validator.CustomValidator,
	c *usecase.TesUseCase,
) *TesController {
	return &TesController{
		Logger:     l,
		Validator:  v,
		TesUseCase: c,
	}
}

func (ctrl *TesController) logError(err error, msg string) {
	ctrl.Logger.Logger.WithError(err).Error(msg)
}

func (ctrl *TesController) List(c echo.Context) error {

	resp, err := ctrl.TesUseCase.List()
	if err != nil {
		errResponse := &httpresponse.ErrorResponse{}
		errResponse.Code = http.StatusBadRequest
		errResponse.Message = "failed to list monitoring sales order free"
		errResponse.Error = err.Error()
		ctrl.logError(err, errResponse.Message)
		return errResponse.EchoJsonResponse(c)
	}

	// Jika data ditemukan, kirimkan respons sukses
	successResponse := &httpresponse.DataResponse[any]{
		Response: httpresponse.Response{
			Code:    http.StatusOK,
			Message: "Successfully updated Document IDs",
		},
		Data: resp,
	}

	return successResponse.EchoJsonResponse(c)
}
