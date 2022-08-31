package event

import (
	"time"
)

// Event - model for statistic store
type Event struct {
	AppID             string `json:"app_id"`
	ClientID          string `json:"client_id"`
	AnonymousClientID string `json:"anonymous_client_id"`
	Domain            string `json:"domain"`
	Name              string `json:"name"`
	Category          string `json:"category"`
	Time              time.Time
	Data              string `json:"data,omitempty"`
}
