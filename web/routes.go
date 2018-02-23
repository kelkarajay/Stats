package web

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

const apiV1 string = "/api/v1"

var routes = Routes{
	// v1 Routes
	Route{
		"GetAllStats",
		"GET",
		apiV1 + "/AllStats",
		GetAllStats,
	},
	Route{
		"PostStat",
		"POST",
		apiV1 + "/PostStat",
		PostStat,
	},
	Route{
		"GetAllApps",
		"GET",
		apiV1 + "/Apps",
		GetAllApps,
	},
	Route{
		"GetApp",
		"GET",
		apiV1 + "/Apps/{appID}",
		GetApp,
	},
	Route{
		"CreateApp",
		"POST",
		apiV1 + "/Apps",
		CreateApp,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
