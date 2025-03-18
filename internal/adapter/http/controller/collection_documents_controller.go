package controller

import (
	"service-collection/internal/adapter/validator"
	"service-collection/internal/infrastructure/logger"
	"service-collection/internal/usecase"
)

type CollectionDocumentController struct {
	Logger                    *logger.Logger
	Validator                 *validator.CustomValidator
	CollectionDocumentUseCase *usecase.CollectionDocumentUseCase
}

func NewCollectionDocumentController(
	l *logger.Logger,
	v *validator.CustomValidator,
	c *usecase.CollectionDocumentUseCase,
) *CollectionDocumentController {
	return &CollectionDocumentController{
		Logger:                    l,
		Validator:                 v,
		CollectionDocumentUseCase: c,
	}
}

func (ctrl *CollectionDocumentController) logError(err error, msg string) {
	ctrl.Logger.Logger.WithError(err).Error(msg)
}
