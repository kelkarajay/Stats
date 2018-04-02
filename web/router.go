package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Xivolkar/Stats/app"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, app.AppContext)

// makeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func makeHandler(ctx app.AppContext, fn func(http.ResponseWriter, *http.Request, app.AppContext)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	}
}

// NewRouter - constructs a router with all possible routes to the app
func NewRouter(ctx app.AppContext) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := makeHandler(ctx, route.HandlerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
