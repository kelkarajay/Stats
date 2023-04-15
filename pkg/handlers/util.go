package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const JSON_CONTENT_TYPE = "application/json; charset=UTF-8"

type Util struct {
	logger *zap.Logger
}

func NewHttpUtil(logger *zap.Logger) *Util {
	util := &Util{logger: logger}
	return util
}

func (u *Util) WriteJSONSuccessResponse(w http.ResponseWriter, body []byte) {
	u.writeJsonResponse(w, http.StatusOK, JSON_CONTENT_TYPE, body)
}

func (u *Util) WriteJSONBadRequestResponse(w http.ResponseWriter, err error) {
	body := map[string]string{
		"code":  "BAD_REQUEST",
		"error": err.Error(),
	}

	data, _ := json.Marshal(body)
	u.writeJsonResponse(w, http.StatusBadRequest, JSON_CONTENT_TYPE, data)
}

func (u *Util) WriteJSONInternalServerErrorResponse(w http.ResponseWriter, err error) {
	body := map[string]string{
		"code":  "INTERNAL_SERVER_ERROR",
		"error": err.Error(),
	}

	data, _ := json.Marshal(body)
	u.writeJsonResponse(w, http.StatusBadRequest, JSON_CONTENT_TYPE, data)
}

func (u *Util) writeJsonResponse(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err := w.Write(body)

	if err != nil {
		u.logger.Error("Error writing response.", zap.Error(err))
	}
}
