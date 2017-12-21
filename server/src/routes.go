package main

import (
	"net/http"
)

// Route type def
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - collection of routes
type Routes []Route

const API string = "/api"

var routes = Routes{
	// v1 Routes
	Route{
		"GetAllStats",
		"GET",
		API + "/v1" + "/AllStats",
		GetAllStats,
	},
	Route{
		"PostStat",
		"POST",
		API + "/v1" + "/PostStat",
		PostStat,
	},
}
