package handlers

import (
	"encoding/json"
	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/model"
	"io"
	"io/ioutil"
	"net/http"
)

// GetAllStats - Retrieves all stats
func GetAllStats(w http.ResponseWriter, r *http.Request, ctx app.AppContext) {
	var stats []model.Stat

	stats, err := ctx.DB.GetStats()

	if err != nil {
		returnInternalServerError(&w, err)
		return
	}

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
