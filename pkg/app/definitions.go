package app

// App - model for Apps
type App struct {
	ID          string `json:"id"`
	Name        string `json:"name, omitempty"`
	Description string `json:"description, omitempty"`
}
