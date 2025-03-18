package repository

import (
	"backend_hub/internal/domain/model/entity"
	"backend_hub/internal/infrastructure/database"
	"backend_hub/internal/infrastructure/logger"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryTes struct {
	Repository
	Logger *logger.Logger
}

// Constructor for MonitoringSalesOrderFreeRepository
func NewRepositoryTes(l *logger.Logger, db *database.Kind[*gorm.DB]) *RepositoryTes {
	repo := &RepositoryTes{
		Logger: l,
	}
	repo.DB = db
	return repo
}

// logError logs an error with a specific message
func (r *RepositoryTes) logError(err error, msg string) {
	if r.Logger != nil {
		r.Logger.Logger.WithError(err).Error(msg)
	}
}

// Ping checks database connectivity
func (r *RepositoryTes) Ping() error {
	readDB, err := r.DB.Read.DB()
	if err != nil {
		return fmt.Errorf("failed to get read database instance: %w", err)
	}
	if err := readDB.Ping(); err != nil {
		return fmt.Errorf("read database is not available: %w", err)
	}
	return nil
}

func (r *RepositoryTes) ListTes(result *[]entity.Tes) error {
	trx := r.getTransaction(r.DB.Read)
	defer trx.Rollback()

	if err := trx.Find(result).Error; err != nil {
		r.logError(err, "failed while querying list of monitoring")
		return err
	}

	if err := trx.Commit().Error; err != nil {
		r.logError(err, "failed while committing get list of monitoring")
		return err
	}

	return nil
}
