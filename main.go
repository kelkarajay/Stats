package main

import (
	"log"
	"net/http"

	"github.com/Xivolkar/Stats/db"

	"github.com/Xivolkar/Stats/model"
	"github.com/Xivolkar/Stats/web"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Xivolkar/Stats/app"
)

func main() {
	d, err := db.NewProdDB()
	if err != nil {
		log.Fatalln("No DB")
	}

	ins := &db.Instance{d}

	ctx := app.AppContext{
		DB: ins,
	}
	log.Println(ctx)

	log.Println("Database up and running")
	db.CurrentInstance.AutoMigrate(&model.Stat{}, &model.App{})
	defer db.CurrentInstance.Close()

	router := web.NewRouter()
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
