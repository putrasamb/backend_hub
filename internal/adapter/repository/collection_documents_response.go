package repository

import (
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/logger"
	httprequest "backend_hub/pkg/common/http/request"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryCollectionDocument struct {
	Repository
	Logger *logger.Logger
}

// Constructor for MonitoringSalesOrderFreeRepository
func NewRepositoryCollectionDocument(l *logger.Logger, db *database.Kind[*gorm.DB]) *RepositoryCollectionDocument {
	repo := &RepositoryCollectionDocument{
		Logger: l,
	}
	repo.DB = db
	return repo
}

// logError logs an error with a specific message
func (r *RepositoryCollectionDocument) logError(err error, msg string) {
	if r.Logger != nil {
		r.Logger.Logger.WithError(err).Error(msg)
	}
}

// Ping checks database connectivity
func (r *RepositoryCollectionDocument) Ping() error {
	readDB, err := r.DB.Read.DB()
	if err != nil {
		return fmt.Errorf("failed to get read database instance: %w", err)
	}
	if err := readDB.Ping(); err != nil {
		return fmt.Errorf("read database is not available: %w", err)
	}
	return nil
}

func (r *RepositoryCollectionDocument) applyFilterScope(request *httprequest.ListRequest) func(db *gorm.DB) *gorm.DB {
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

func (r *RepositoryCollectionDocument) applySortScope(request httprequest.FilteredRequestInterface) func(db *gorm.DB) *gorm.DB {
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

func (r *RepositoryCollectionDocument) ListTotal(request *httprequest.ListRequest, count *int64) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()
	if err := trx.Model(&entity.CollectionDocument{}).
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

func (r *RepositoryCollectionDocument) List(request *httprequest.ListRequest, result *[]entity.CollectionDocument) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()

	if err := trx.Scopes(
		Paginate(request.Page, request.PerPage),
		r.applyFilterScope(request),
		r.applySortScope(request),
	).
		Find(result).Error; err != nil {
		r.logError(err, "failed while querying list of data")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing get list data")
		return err
	}

	return nil
}
