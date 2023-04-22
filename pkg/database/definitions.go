package database

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type BaseEntityAttributes struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
