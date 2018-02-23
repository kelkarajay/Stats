package db

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	err := NewDB()
	if err != nil {
		t.Error("Database not setup", err)
	}
}
