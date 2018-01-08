package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
