package db

import (
	"database/sql"

	"github.com/Xivolkar/Stats/model"
	_ "github.com/mattn/go-sqlite3" // Database driver
)

// DataStorer defines all the database operations
type DataStorer interface {
	GetStats() ([]model.Stat, error)
	Close() error
}

// Instance - DB Instance
type Instance struct {
	DB *sql.DB
}

// Close - Shutdown DB
func (ci *Instance) Close() error {
	return ci.DB.Close()
}

func (ci *Instance) GetStats() ([]model.Stat, error) {
	var stats []model.Stat
	return stats, nil
}

// NewDB - Initializes the DB and returns if there is an error in the process
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	return db, err
}
