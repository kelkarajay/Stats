package model

import "time"

// Stat - model for statistic store
type Stat struct {
	StatAppID    string    `json:"statAppID"`
	StatClientID string    `json:"statClientID"`
	StatType     string    `json:"statType"`
	StatCategory string    `json:"statCategory"`
	StatValue    string    `json:"statValue"`
	StatTime     time.Time `json:"statTime"`
	StatData     string    `json:"statData"`
}

// Stats - array of stat objects
type Stats []Stat
