package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/gorm"
)

func RunMigrations(instance *gorm.DB, config Config) error {
	db, err := instance.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		MultiStatementEnabled: true,
		DatabaseName:          config.Name,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil {
		return err
	}

	return m.Up()
}
