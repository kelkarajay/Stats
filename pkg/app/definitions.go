package app

import "github.com/kelkarajay/Stats/pkg/database"

// App - model for Apps
type App struct {
	database.BaseEntityAttributes
	Name        string `json:"name, omitempty"`
	Description string `json:"description, omitempty"`
}
