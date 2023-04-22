package event

import (
	"github.com/kelkarajay/Stats/pkg/database"
)

// Event - model for statistic store
type Event struct {
	database.BaseEntityAttributes
	AppID     string `json:"app_id"`
	ClientID  string `json:"client_id"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Data      string `json:"data,omitempty"`
}
