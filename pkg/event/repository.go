package event

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryOperations interface {
	GetStats(context context.Context) ([]Event, error)
	CreateEvent(context context.Context, event Event) error
}

func NewStatRepository(db *gorm.DB) RepositoryOperations {
	return Repository{db: db}
}

func (r Repository) GetStats(context context.Context) ([]Event, error) {
	var stats []Event

	result := r.db.WithContext(context).Find(&stats)
	if result.Error != nil {
		return stats, result.Error
	}

	return stats, nil
}

func (r Repository) CreateEvent(context context.Context, event Event) error {
	result := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&event)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
