package handlers

import (
	"net/http"

	"github.com/Xivolkar/Stats/app"
)

func makeHandler(ctx app.AppContext, fn func(http.ResponseWriter, *http.Request, app.AppContext)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	}
}
