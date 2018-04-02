package web

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/model"
)

type PageVariables struct {
	Date string
	Time string
}

func Index(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("./web/index.html") //parse the html file homepage.html
	if err != nil {                                   // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

// GetAllStats - Retrieves all stats
func GetAllStats(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var stats []model.Stat

	stats, _ = ctx.DB.GetStats()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&stats); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
}

// PostStat - Creates and stores stat
func PostStat(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var st model.Stat
	var err error
	var body []byte

	if r.Body != nil {
		body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			returnBadRequest(&w, err)
			return
		}
	} else {
		returnBadRequest(&w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		returnBadRequest(&w, err)
		return
	}

	if err := json.Unmarshal(body, &st); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	// TODO: Create

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&st); err != nil {
		panic(err)
	}
}

func GetAllApps(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var apps []model.App
	// TODO : QUERY
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&apps); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
}

func GetApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App

	// TODO : QUERY

	json.NewEncoder(w).Encode(&app)
}

func CreateApp(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var app model.App
	json.NewDecoder(r.Body).Decode(&app)

	// TODO : Create

	json.NewEncoder(w).Encode(&app)
}

func returnBadRequest(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder((*w)).Encode(err); err != nil {
		panic(err)
	}
}
