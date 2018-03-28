package db

import (
	"database/sql"
	"log"

	"github.com/stretchr/testify/mock"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // Database driver
)

// CurrentInstance - Pointer to the database session
var CurrentInstance *gorm.DB

// DataStorer defines all the database operations
type DataStorer interface {
	//ListUsers() ([]User, error)
	//GetUser(i int) (User, error)
	//AddUser(u User) (User, error)
	//UpdateUser(u User) (User, error)
	DeleteUser(i int) error
}

// MockInstance - For tests
type MockInstance struct {
	mock.Mock
}

// Instance - DB Instance
type Instance struct {
	db *sql.DB
}

func (mi *MockInstance) DeleteUser(i int) error {
	returnVals := mi.Called(i)
	// return the values which we define
	return returnVals.Error(1)
}

func (ci *Instance) DeleteUser(i int) error {
	return nil
}

// NewDB - Initializes the DB and returns if there is an error in the process
func NewDB() (err error) {
	log.Println("Connecting to the Database")
	CurrentInstance, err = gorm.Open("sqlite3", "./foo.db")
	return err
}

func NewProdDB() (*sql.DB, error) {
db, err := sql.Open("sqlite3", "./foo.db")
return db, err
}
