package main

import (
	"log"
	"net/http"

	"github.com/Xivolkar/Stats/web"
)

func main() {
	router := web.NewRouter()
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
