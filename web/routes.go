package web

import "github.com/Xivolkar/Stats/web/handlers"

// Route type def
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes - collection of routes
type Routes []Route

const apiV1 string = "/api/v1"

var routes = Routes{
	// v1 Routes
	Route{
		"GetAllStats",
		"GET",
		apiV1 + "/Stats",
		handlers.GetAllStats,
	},
	Route{
		"PostStat",
		"POST",
		apiV1 + "/Stats",
		handlers.PostStat,
	},
	Route{
		"GetAllApps",
		"GET",
		apiV1 + "/Apps",
		handlers.GetAllApps,
	},
	Route{
		"GetApp",
		"GET",
		apiV1 + "/Apps/{appID}",
		handlers.GetApp,
	},
	Route{
		"CreateApp",
		"POST",
		apiV1 + "/Apps",
		handlers.CreateApp,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
