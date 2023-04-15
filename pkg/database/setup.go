package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()),
		&gorm.Config{
			NowFunc: func() time.Time { return time.Now().UTC() },
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)

	return db, nil
}
