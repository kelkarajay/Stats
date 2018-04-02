package main

import (
	"log"
	"net/http"

	"github.com/Xivolkar/Stats/db"

	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	d, err := db.NewDB()
	if err != nil {
		log.Fatalln("No DB")
	}

	instance := &db.Instance{DB: d}

	ctx := app.AppContext{
		DB: instance,
	}
	log.Println(ctx)

	log.Println("Database up and running")
	defer ctx.DB.Close()

	router := web.NewRouter(ctx)
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
