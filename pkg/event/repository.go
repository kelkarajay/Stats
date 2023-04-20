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
	GetEvents(context context.Context) ([]Event, error)
	CreateEvent(context context.Context, event Event) error
}

func NewStatRepository(db *gorm.DB) RepositoryOperations {
	return Repository{db: db}
}

func (r Repository) GetEvents(context context.Context) ([]Event, error) {
	var events []Event

	result := r.db.WithContext(context).Find(&events)
	if result.Error != nil {
		return events, result.Error
	}

	return events, nil
}

func (r Repository) CreateEvent(context context.Context, event Event) error {
	result := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&event)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
