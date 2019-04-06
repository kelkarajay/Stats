package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/model"
)

// GetAllApps - handler to query and return all Applications
func GetAllApps(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var apps []model.App
	// TODO : QUERY

	apps, err := ctx.DB.GetApps()

	if err != nil {
		returnInternalServerError(&w, err)
		return
	}

	returnBuiltResponse(&w, apps)
}

// GetApp - Handler to return specific app
func GetApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App

	// TODO : QUERY

	json.NewEncoder(w).Encode(&app)
}

// CreateApp - Handler to create new application for Stats
func CreateApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App
	json.NewDecoder(r.Body).Decode(&app)

	// TODO : Create

	json.NewEncoder(w).Encode(&app)
}
