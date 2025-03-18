package usecase

import (
	"backend_hub/internal/adapter/repository"
	"backend_hub/internal/adapter/validator"
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/logger"
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
