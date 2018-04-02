package db

import (
	"github.com/Xivolkar/Stats/model"
	"github.com/stretchr/testify/mock"
)

// MockInstance - For tests
type MockInstance struct {
	mock.Mock
}

func (mi *MockInstance) Close() error {
	return nil
}

func (mi *MockInstance) GetStats() ([]model.Stat, error) {
	var stats []model.Stat
	return stats, nil
}
