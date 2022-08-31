package handlers

import (
	"encoding/json"
	"github.com/kelkarajay/Stats/pkg/app"
	"go.uber.org/zap"
	"net/http"

	"github.com/Xivolkar/Stats/model"
)

type AppHandlerOperations interface {
	GetAllStats(w http.ResponseWriter, r *http.Request)
	PostStat(w http.ResponseWriter, r *http.Request)
}

type appHandler struct {
	appRepository app.RepositoryOperations
	logger        *zap.Logger
	httpUtil      *Util
}

func NewAppHandler(repository app.RepositoryOperations, logger *zap.Logger) *appHandler {
	return &appHandler{appRepository: repository, logger: logger, httpUtil: NewHttpUtil(logger)}
}

// GetAllApps - statHandler to query and return all Applications
func (a *appHandler) GetAllApps(w http.ResponseWriter, r *http.Request) {
	// TODO : QUERY
	apps, err := a.appRepository.GetApps(r.Context())

	if err != nil {
		a.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
		return
	}

	jsonR, err := json.Marshal(apps)
	if err != nil {
		a.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	a.httpUtil.WriteJSONSuccessResponse(w, jsonR)
}

// GetApp - Handler to return specific app
func (a *appHandler) GetApp(w http.ResponseWriter, r *http.Request) {
	var app model.App

	// TODO : QUERY

	jsonR, err := json.Marshal(app)
	if err != nil {
		a.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	a.httpUtil.WriteJSONSuccessResponse(w, jsonR)
}

// CreateApp - Handler to create new application for Stats
func (a *appHandler) CreateApp(w http.ResponseWriter, r *http.Request) {
	var app model.App
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		a.httpUtil.WriteJSONBadRequestResponse(w, err)
		return
	}

	// TODO : Create

	jsonR, err := json.Marshal(app)
	if err != nil {
		a.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	a.httpUtil.WriteJSONSuccessResponse(w, jsonR)
}
