package db

import (
	"database/sql"
	"log"

	"github.com/Xivolkar/Stats/model"
	_ "github.com/mattn/go-sqlite3" // Database driver
)

// DataStorer defines all the database operations
type DataStorer interface {
	Migrate() error
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

	stmt, err := ci.DB.Prepare("select * from Stat")
	if err != nil {
		log.Println("DB Query failed")
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error during query")
		return nil, err
	}

	defer rows.Close()

	var stat model.Stat
	for rows.Next() {
		err := rows.Scan(&stat)
		stats = append(stats, stat)
		if err != nil {
			log.Println("Error during rows scan")
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return stats, nil
}

func (ci *Instance) Migrate() error {
	_, err := ci.DB.Exec("CREATE TABLE IF NOT EXISTS Stat(StatAppID VARCHAR(255), StatClientID VARCHAR(255), StatType VARCHAR(255), StatCategory VARCHAR(255), StatValue VARCHAR(255), StatTime VARCHAR(255), StatData VARCHAR(255))")

	return err
}

// NewDB - Initializes the DB and returns if there is an error in the process
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	return db, err
}
