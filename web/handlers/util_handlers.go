package handlers

import (
	"encoding/json"
	"net/http"
)

func returnBadRequest(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder((*w)).Encode(err); err != nil {
		panic(err)
	}
}

func returnInternalServerError(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder((*w)).Encode(err); err != nil {
		panic(err)
	}
}
