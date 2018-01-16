package model

// App - model for Apps
type App struct {
	AppID          string `json:"appID"`
	AppName        string `json:"appName"`
	AppDescription string `json:"appDescription"`
	AppIconURL     string `json:"appIconUrl"`
}
