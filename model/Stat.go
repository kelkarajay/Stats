package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Stat - model for statistic store
type Stat struct {
	gorm.Model
	StatAppID    string    `json:"statAppID"`
	StatClientID string    `json:"statClientID"`
	StatType     string    `json:"statType"`
	StatCategory string    `json:"statCategory"`
	StatValue    string    `json:"statValue"`
	StatTime     time.Time `json:"statTime,omitempty"`
	StatData     string    `json:"statData,omitempty"`
}

// Stats - array of stat objects
type Stats []Stat
