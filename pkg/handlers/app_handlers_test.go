package handlers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kelkarajay/Stats/pkg/app"
	"github.com/kelkarajay/Stats/pkg/database"
	"github.com/kelkarajay/Stats/pkg/handlers"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Define the suite
type AppTestSuite struct {
	suite.Suite
	db            *gorm.DB
	logger        *zap.Logger
	appRepository app.RepositoryOperations
}

func (suite *AppTestSuite) SetupTest() {
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
	suite.appRepository = app.NewAppRepository(db)
}

func (suite *AppTestSuite) TestGetAllApps(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/Apps", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := handlers.NewAppHandler(suite.appRepository, suite.logger)

	handler.GetAllApps(resp, req)

	log.Print(resp.Body)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("statHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
