package usecase

import (
	"backend_hub/internal/adapter/converter"
	"backend_hub/internal/adapter/dto/response"
	"backend_hub/internal/adapter/repository"
	"backend_hub/internal/adapter/validator"
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/logger"
	httprequest "backend_hub/pkg/common/http/request"
	httpresponse "backend_hub/pkg/common/http/response"

	"github.com/pkg/errors"
)

type CollectionDocumentUseCase struct {
	Logger                       *logger.Logger
	Validator                    *validator.CustomValidator
	RepositoryCollectionDocument *repository.RepositoryCollectionDocument
}

func NewUseCaseCollectionDocuments(
	l *logger.Logger,
	v *validator.CustomValidator,
	r *repository.RepositoryCollectionDocument,
) *CollectionDocumentUseCase {
	return &CollectionDocumentUseCase{
		Logger:                       l,
		Validator:                    v,
		RepositoryCollectionDocument: r,
	}
}

func (u *CollectionDocumentUseCase) logError(err error, msg string) {
	u.Logger.Logger.WithError(err).Error(msg)
}

func (u *CollectionDocumentUseCase) List(req *httprequest.ListRequest) (*httpresponse.PaginatedReponse[*response.CollectionDocument], error) {
	if err := u.Validator.Validate(req); err != nil {
		errMessage := "failed to validate list request"
		validationErrors := u.Validator.ParseValidationErrors(err)
		u.logError(validationErrors, errMessage)
		return nil, validationErrors
	}

	if err := req.DecodeFilters(); err != nil {
		errMessage := "failed to decode list filters"
		u.logError(err, errMessage)
		return nil, err
	}

	if err := req.DecodeSort(); err != nil {
		errMessage := "failed to decode list sort"
		u.logError(err, errMessage)
		return nil, err
	}

	var data []entity.CollectionDocument
	if err := u.RepositoryCollectionDocument.List(req, &data); err != nil {
		errMessage := "failed to list objects"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	res := converter.ConvertCollectionDocumentListResponse(&data)
	total := int64(0)
	if err := u.RepositoryCollectionDocument.ListTotal(req, &total); err != nil {
		errMessage := "failed to count list objects"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	paginated := &httpresponse.PaginatedReponse[*response.CollectionDocument]{
		Data:        *res,
		Total:       total,
		CurrentPage: req.Page,
		PerPage:     req.PerPage,
	}
	return paginated, nil
}
