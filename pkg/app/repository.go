package app

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryOperations interface {
	GetApps(context context.Context) ([]App, error)
}

func NewAppRepository(db *gorm.DB) RepositoryOperations {
	return Repository{db: db}
}

func (r Repository) GetApps(context context.Context) ([]App, error) {
	var apps []App

	result := r.db.WithContext(context).Find(&apps)
	if result.Error != nil {
		return apps, result.Error
	}

	return apps, nil
}
