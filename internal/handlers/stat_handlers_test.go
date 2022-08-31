package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllStatsHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/AllStats", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := app.CreateContextForTestSetup()

	resp := httptest.NewRecorder()
	handler := makeHandler(ctx, GetAllStats)

	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("statHandler returned wrong status code: got %v want %v",
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

	ctx := app.CreateContextForTestSetup()

	resp := httptest.NewRecorder()
	handler := makeHandler(ctx, PostStat)

	handler.ServeHTTP(resp, req)

	var status int

	if resp != nil {
		status = resp.Code
	}

	if status != http.StatusCreated {
		t.Errorf("statHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		b, _ := ioutil.ReadAll(resp.Body)
		t.Error(string(b))
	}
}
