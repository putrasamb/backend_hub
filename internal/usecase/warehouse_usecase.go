package usecase

import (
	dto "backend_hub/internal/adapter/dto/request"
	dtores "backend_hub/internal/adapter/dto/response"
	"backend_hub/internal/adapter/repository"
	"backend_hub/internal/adapter/validator"
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/logger"
	httprequest "backend_hub/pkg/common/http/request"
	httpresponse "backend_hub/pkg/common/http/response"

	"github.com/pkg/errors"
)

type WarehouseUseCase struct {
	Logger            *logger.Logger
	Validator         *validator.CustomValidator
	ProductRepository *repository.ProductRepository
}

func NewWarehouseUseCase(
	l *logger.Logger,
	v *validator.CustomValidator,
	r *repository.ProductRepository,
) *WarehouseUseCase {
	return &WarehouseUseCase{
		Logger:            l,
		Validator:         v,
		ProductRepository: r,
	}
}

func (u *WarehouseUseCase) logError(err error, msg string) {
	u.Logger.Logger.WithError(err).Error(msg)
}

func (u *WarehouseUseCase) List(req *httprequest.ListRequest) (*httpresponse.PaginatedReponse[*dtores.ProductResponse], error) {
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

	var data []entity.Product
	if err := u.ProductRepository.List(req, &data); err != nil {
		errMessage := "failed to list products"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	res := make([]*dtores.ProductResponse, len(data))
	for i, product := range data {
		res[i] = &dtores.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Status:      product.Status,
			CreatedBy:   product.CreatedBy,
			UpdatedBy:   product.UpdatedBy,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}
	total := int64(0)
	if err := u.ProductRepository.ListTotal(req, &total); err != nil {
		errMessage := "failed to count list products"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	paginated := &httpresponse.PaginatedReponse[*dtores.ProductResponse]{
		Data:        res,
		Total:       total,
		CurrentPage: req.Page,
		PerPage:     req.PerPage,
	}
	return paginated, nil
}

func (u *WarehouseUseCase) Create(req *dto.CreateProductRequest) (*dtores.ProductResponse, error) {
	if err := u.Validator.Validate(req); err != nil {
		errMessage := "failed to validate create product request"
		validationErrors := u.Validator.ParseValidationErrors(err)
		u.logError(validationErrors, errMessage)
		return nil, validationErrors
	}

	product := &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
	}

	if err := u.ProductRepository.Create(product); err != nil {
		errMessage := "failed to create product"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	return &dtores.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (u *WarehouseUseCase) Update(req *dto.UpdateProductRequest) (*dtores.ProductResponse, error) {
	if err := u.Validator.Validate(req); err != nil {
		errMessage := "failed to validate update product request"
		validationErrors := u.Validator.ParseValidationErrors(err)
		u.logError(validationErrors, errMessage)
		return nil, validationErrors
	}

	product := &entity.Product{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
	}

	if err := u.ProductRepository.Update(product); err != nil {
		errMessage := "failed to update product"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	return &dtores.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (u *WarehouseUseCase) Delete(req *dto.DeleteProductRequest) error {
	if err := u.Validator.Validate(req); err != nil {
		errMessage := "failed to validate delete product request"
		validationErrors := u.Validator.ParseValidationErrors(err)
		u.logError(validationErrors, errMessage)
		return validationErrors
	}

	if err := u.ProductRepository.Delete(req.ID); err != nil {
		errMessage := "failed to delete product"
		u.logError(err, errMessage)
		return errors.Wrap(err, errMessage)
	}

	return nil
}

func (u *WarehouseUseCase) Get(req *dto.GetProductRequest) (*dtores.ProductResponse, error) {
	if err := u.Validator.Validate(req); err != nil {
		errMessage := "failed to validate get product request"
		validationErrors := u.Validator.ParseValidationErrors(err)
		u.logError(validationErrors, errMessage)
		return nil, validationErrors
	}

	var product entity.Product
	if err := u.ProductRepository.Get(req.ID, &product); err != nil {
		errMessage := "failed to get product"
		u.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}

	return &dtores.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
