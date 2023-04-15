package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kelkarajay/Stats/pkg/event"
	"go.uber.org/zap"
)

type StatHandlerOperations interface {
	GetAllStats(w http.ResponseWriter, r *http.Request)
	PostStat(w http.ResponseWriter, r *http.Request)
}

type statHandler struct {
	eventRepository event.RepositoryOperations
	logger          *zap.Logger
	httpUtil        *Util
}

func NewStatHandler(repository event.RepositoryOperations, logger *zap.Logger) *statHandler {
	return &statHandler{eventRepository: repository, logger: logger, httpUtil: NewHttpUtil(logger)}
}

// GetAllStats - Retrieves all stats
func (h *statHandler) GetAllStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.eventRepository.GetStats(r.Context())

	if err != nil {
		h.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	data, err := json.Marshal(stats)
	if err != nil {
		h.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	h.httpUtil.writeJsonResponse(w, http.StatusOK, JSON_CONTENT_TYPE, data)
}

// PostStat - Creates and stores event
func (h *statHandler) PostStat(w http.ResponseWriter, r *http.Request) {
	var st event.Event
	var err error
	var body []byte

	if r.Body != nil {
		body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			h.httpUtil.WriteJSONBadRequestResponse(w, err)
		}
	} else {
		h.httpUtil.WriteJSONBadRequestResponse(w, err)
	}

	if err := r.Body.Close(); err != nil {
		h.httpUtil.WriteJSONBadRequestResponse(w, err)
	}

	if err := json.Unmarshal(body, &st); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			h.logger.Error("Writing response failed", zap.Error(err))
		}
		return
	}

	err = h.eventRepository.CreateEvent(r.Context(), st)
	if err != nil {
		h.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	data, err := json.Marshal(st)
	if err != nil {
		h.httpUtil.WriteJSONInternalServerErrorResponse(w, err)
	}

	h.httpUtil.writeJsonResponse(w, http.StatusCreated, JSON_CONTENT_TYPE, data)
}
