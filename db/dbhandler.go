package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // Database driver
)

// CurrentInstance - Pointer to the database session
var CurrentInstance *gorm.DB

// NewDB - Initializes the DB and returns if there is an error in the process
func NewDB() (err error) {
	log.Println("Connecting to the Database")
	CurrentInstance, err = gorm.Open("sqlite3", "./foo.db")
	return err
}
