package handlers

import (
	"errors"
	"github.com/Xivolkar/Stats/app"
	"github.com/Xivolkar/Stats/db"
	"github.com/Xivolkar/Stats/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllApps(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/Apps", nil)
	if err != nil {
		log.Fatal(err)
	}

	db := new(db.MockInstance)

	db.On("GetApps").Return([]model.App{{AppID: "123", AppName: "App1", AppDescription: "Test description"}}, nil)

	ctx := app.CreateContextForTestSetup()
	ctx.DB = db

	resp := httptest.NewRecorder()
	handler := makeHandler(ctx, GetAllApps)

	handler.ServeHTTP(resp, req)

	log.Print(resp.Body)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("statHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetAllApps_DatabaseFetchError(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/Apps", nil)
	if err != nil {
		log.Fatal(err)
	}

	db := new(db.MockInstance)

	db.On("GetApps").Return([]model.App{}, errors.New("Database query failed"))

	ctx := app.CreateContextForTestSetup()
	ctx.DB = db

	resp := httptest.NewRecorder()
	handler := makeHandler(ctx, GetAllApps)

	handler.ServeHTTP(resp, req)

	log.Print(resp.Body)

	if status := resp.Code; status != http.StatusInternalServerError {
		t.Errorf("statHandler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
