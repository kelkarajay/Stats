package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xivolkar/Stats/model"
)

func TestGetAllStatsHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/AllStats", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllStats)

	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostStatHandler(t *testing.T) {
	j := &model.Stat{}
	payload, _ := json.Marshal(j)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/PostStat", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(PostStat)

	handler.ServeHTTP(resp, req)

	var status int

	if resp != nil {
		status = resp.Code
	}

	if status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		b, _ := ioutil.ReadAll(resp.Body)
		t.Error(string(b))
	}
}
