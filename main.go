package main

import (
	"log"
	"net/http"

	"github.com/Xivolkar/Stats/db"

	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/web"
)

func main() {
	log.Println("Connecting to the Database")
	d, err := db.NewDB()
	if err != nil {
		log.Fatalln("Database Init failed. Stopping boot")
	}
	log.Println("Database up and running")

	instance := &db.Instance{DB: d}

	ctx := app.AppContext{
		DB: instance,
	}

	defer ctx.DB.Close()

	router := web.NewRouter(ctx)
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
