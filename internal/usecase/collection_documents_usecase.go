package usecase

import (
	"service-collection/internal/adapter/converter"
	"service-collection/internal/adapter/dto/response"
	"service-collection/internal/adapter/repository"
	"service-collection/internal/adapter/validator"
	"service-collection/internal/domain/model/entity"
	"service-collection/internal/infrastructure/logger"
	httprequest "service-collection/pkg/common/http/request"
	httpresponse "service-collection/pkg/common/http/response"

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
