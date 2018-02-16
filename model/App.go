package model

// App - model for Apps
type App struct {
	AppID          string `json:"appID"`
	AppName        string `json:"appName, omitempty"`
	AppDescription string `json:"appDescription, omitempty"`
	AppIconURL     string `json:"appIconUrl, omitempty"`
}
