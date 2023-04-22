package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xivolkar/Stats/model"
	"github.com/kelkarajay/Stats/pkg/database"
	"github.com/kelkarajay/Stats/pkg/event"
	"github.com/kelkarajay/Stats/pkg/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Define the suite
type StatTestSuite struct {
	suite.Suite
	db              *gorm.DB
	logger          *zap.Logger
	eventRepository event.RepositoryOperations
}

func (suite *StatTestSuite) SetupTest() {
	logger := zap.NewNop()
	databaseConfig, err := database.LoadConfig()
	if err != nil {
		logger.Fatal("Could not load database configuration", zap.Error(err))
	}

	db, err := database.SetupDatabase(databaseConfig)
	if err != nil {
		logger.Fatal("Database init failed", zap.Error(err))
	}

	suite.db = db
	suite.logger = logger
	suite.eventRepository = event.NewStatRepository(db)
}

func (suite *StatTestSuite) TestGetAllStatsHandler() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/AllStats", nil)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewStatHandler(suite.eventRepository, suite.logger)
	resp := httptest.NewRecorder()

	handler.GetAllStats(resp, req)

	if status := resp.Code; status != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(suite.T(), http.StatusOK, status, fmt.Sprintf("body %s", string(b)))
	}
}

func (suite *StatTestSuite) TestPostStatHandler() {
	j := &model.Stat{}
	payload, _ := json.Marshal(j)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/PostStat", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewStatHandler(suite.eventRepository, suite.logger)
	resp := httptest.NewRecorder()

	handler.PostStat(resp, req)

	var status int

	if resp != nil {
		status = resp.Code
	}

	if status != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(suite.T(), http.StatusCreated, status, fmt.Sprintf("body %s", string(b)))
	}
}

func TestStatTestSuite(t *testing.T) {
	suite.Run(t, new(StatTestSuite))
}
