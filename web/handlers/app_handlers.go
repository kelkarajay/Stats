package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/model"
)

func GetAllApps(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var apps []model.App
	// TODO : QUERY
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&apps); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)

	}
}

func GetApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App

	// TODO : QUERY

	json.NewEncoder(w).Encode(&app)
}

func CreateApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App
	json.NewDecoder(r.Body).Decode(&app)

	// TODO : Create

	json.NewEncoder(w).Encode(&app)
}
