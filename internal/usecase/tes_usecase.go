package usecase

import (
	"service-collection/internal/adapter/repository"
	"service-collection/internal/adapter/validator"
	"service-collection/internal/domain/model/entity"
	"service-collection/internal/infrastructure/logger"
)

type TesUseCase struct {
	Logger        *logger.Logger
	Validator     *validator.CustomValidator
	RepositoryTes *repository.RepositoryTes
}

func NewUseCaseTes(
	l *logger.Logger,
	v *validator.CustomValidator,
	r *repository.RepositoryTes,
) *TesUseCase {
	return &TesUseCase{
		Logger:        l,
		Validator:     v,
		RepositoryTes: r,
	}
}

func (u *TesUseCase) logError(err error, msg string) {
	u.Logger.Logger.WithError(err).Error(msg)
}

func (u *TesUseCase) List() (*[]entity.Tes, error) {

	var data []entity.Tes
	if err := u.RepositoryTes.ListTes(&data); err != nil {
		errMessage := "Failed to list monitoring sales order free"
		u.logError(err, errMessage)
		return nil, err
	}

	return &data, nil
}
