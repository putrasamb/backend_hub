package repository

import (
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/logger"
	httprequest "backend_hub/pkg/common/http/request"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository
	Logger *logger.Logger
}

// Constructor for ProductRepository
func NewProductRepository(l *logger.Logger, db *database.Kind[*gorm.DB]) *ProductRepository {
	repo := &ProductRepository{
		Logger: l,
	}
	repo.DB = db
	return repo
}

// logError logs an error with a specific message
func (r *ProductRepository) logError(err error, msg string) {
	if r.Logger != nil {
		r.Logger.Logger.WithError(err).Error(msg)
	}
}

// Ping checks database connectivity
func (r *ProductRepository) Ping() error {
	readDB, err := r.DB.Read.DB()
	if err != nil {
		return fmt.Errorf("failed to get read database instance: %w", err)
	}
	if err := readDB.Ping(); err != nil {
		return fmt.Errorf("read database is not available: %w", err)
	}
	return nil
}

func (r *ProductRepository) applyFilterScope(request *httprequest.ListRequest) func(db *gorm.DB) *gorm.DB {
	return func(trx *gorm.DB) *gorm.DB {

		if request.Filters == nil {
			return trx
		}

		for _, filter := range *request.Filters {
			operator := r.ParseFilterOperator(filter.Operator)
			value := filter.Value

			if operator == "like" {
				value = fmt.Sprintf("%%%v%%", value)
			}
			trx = trx.Where(fmt.Sprintf("%s %s ?", filter.Field, operator), value)
		}

		return trx
	}
}

func (r *ProductRepository) applySortScope(request httprequest.FilteredRequestInterface) func(db *gorm.DB) *gorm.DB {
	return func(trx *gorm.DB) *gorm.DB {
		if request.GetSort() == nil {
			return trx
		}

		for _, sort := range *request.GetSort() {
			trx = trx.Order(fmt.Sprintf("%s %s", sort.Field, sort.Direction))
		}
		return trx
	}
}

func (r *ProductRepository) ListTotal(request *httprequest.ListRequest, count *int64) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()
	if err := trx.Model(&entity.Product{}).
		Scopes(
			r.applyFilterScope(request),
			r.applySortScope(request),
		).Count(count).Error; err != nil {
		return err
	}
	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing get list data")
		return err
	}
	return nil
}

// Create creates a new product
func (r *ProductRepository) Create(product *entity.Product) error {
	trx := r.getTransaction(r.DB.Write)
	defer trx.Rollback()

	if err := trx.Create(product).Error; err != nil {
		r.logError(err, "failed to create product")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing create product")
		return err
	}

	return nil
}

// Update updates an existing product
func (r *ProductRepository) Update(product *entity.Product) error {
	trx := r.getTransaction(r.DB.Write)
	defer trx.Rollback()

	if err := trx.Save(product).Error; err != nil {
		r.logError(err, "failed to update product")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing update product")
		return err
	}

	return nil
}

// Delete deletes a product by ID
func (r *ProductRepository) Delete(id int) error {
	trx := r.getTransaction(r.DB.Write)
	defer trx.Rollback()

	if err := trx.Delete(&entity.Product{}, id).Error; err != nil {
		r.logError(err, "failed to delete product")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing delete product")
		return err
	}

	return nil
}

// Get retrieves a product by ID
func (r *ProductRepository) Get(id int, product *entity.Product) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()

	if err := trx.First(product, id).Error; err != nil {
		r.logError(err, "failed to get product")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing get product")
		return err
	}

	return nil
}

// List retrieves a list of products with pagination
func (r *ProductRepository) List(request *httprequest.ListRequest, result *[]entity.Product) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()

	if err := trx.Scopes(
		Paginate(request.Page, request.PerPage),
		r.applyFilterScope(request),
		r.applySortScope(request),
	).Find(result).Error; err != nil {
		r.logError(err, "failed while querying list of products")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing get list of products")
		return err
	}

	return nil
}
