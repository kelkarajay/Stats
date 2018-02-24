package main

import (
	"os"
	"testing"

	"github.com/Xivolkar/Stats/db"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	db.NewDB()
	os.Exit(m.Run())
}
