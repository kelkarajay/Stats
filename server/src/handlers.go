package main

import (
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func PostEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test", "Responding to post")
}
