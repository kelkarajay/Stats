package main

import (
	"log"
	"net/http"

	"github.com/Xivolkar/Stats/db"

	"github.com/Xivolkar/Stats/model"
	"github.com/Xivolkar/Stats/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := db.NewDB(); err != nil {
		log.Fatalln("No DB")
	}
	log.Println("Database up and running")
	db.CurrentInstance.AutoMigrate(&model.Stat{}, &model.App{})
	defer db.CurrentInstance.Close()

	router := web.NewRouter()
	log.Println("Starting Stats Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
