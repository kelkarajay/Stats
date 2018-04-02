package db

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	_, err := NewDB()
	if err != nil {
		t.Error("Database not setup", err)
	}
}
