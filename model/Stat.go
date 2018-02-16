package model

import "time"

// Stat - model for statistic store
type Stat struct {
	StatAppID    App       `json:"app"`
	StatClientID string    `json:"statClientID"`
	StatType     string    `json:"statType"`
	StatCategory string    `json:"statCategory"`
	StatValue    string    `json:"statValue"`
	StatTime     time.Time `json:"statTime,omitempty"`
	StatData     string    `json:"statData,omitempty"`
}

// Stats - array of stat objects
type Stats []Stat
