package model

import (
	"github.com/jinzhu/gorm"
)

// App - model for Apps
type App struct {
	gorm.Model
	AppID          string `json:"appID"`
	AppName        string `json:"appName, omitempty"`
	AppDescription string `json:"appDescription, omitempty"`
	AppIconURL     string `json:"appIconUrl, omitempty"`
}

type Apps []App
