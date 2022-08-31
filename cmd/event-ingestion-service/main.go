package main

import (
	"github.com/gorilla/mux"
	"github.com/kelkarajay/Stats/internal/handlers"
	"github.com/kelkarajay/Stats/pkg/database"
	"github.com/kelkarajay/Stats/pkg/event"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HandlerRegistry struct {
	StatHandler handlers.StatHandlerOperations
}

// NewRouter - constructs a router with all possible routes to the app
func NewRouter(registry HandlerRegistry) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/event", registry.StatHandler.GetAllStats).Methods("GET")
	router.HandleFunc("/event", registry.StatHandler.PostStat).Methods("POST")

	return router
}

func main() {
	logger := zap.L()

	logger.Info("Connecting to the database")
	databaseConfig, err := database.LoadConfig()
	if err != nil {
		logger.Fatal("Could not load database configuration", zap.Error(err))
	}

	db, err := database.SetupDatabase(databaseConfig)
	if err != nil {
		logger.Fatal("Database init failed", zap.Error(err))
	}

	eventRepository := event.NewStatRepository(db)
	statHandler := handlers.NewStatHandler(eventRepository, logger)

	router := NewRouter(HandlerRegistry{StatHandler: statHandler})

	logger.Info("Starting server")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Fatal("Failed starting the server", zap.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
