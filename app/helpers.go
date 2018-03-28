package app

import (
	"github.com/unrolled/render"
	"github.com/Xivolkar/Stats/db"
)

// AppContext holds application configuration data
type AppContext struct {
	Render  *render.Render
	Version string
	Env     string
	Port    string
	DB      db.DataStorer
}

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() AppContext {
	testVersion := "0.0.0"
	db := CreateMockDatabase()
	ctx := AppContext{
		Render:  render.New(),
		Version: testVersion,
		Env:     "local",
		Port:    "3001",
		DB:      db,
	}
	return ctx
}

func CreateMockDatabase() *db.MockInstance {
	return new(db.MockInstance)
}
