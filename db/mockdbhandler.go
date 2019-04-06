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

func (mi *MockInstance) GetApps() ([]model.App, error) {
	retVal := mi.Called()
	return retVal.Get(0).([]model.App), retVal.Error(1)
}

func (mi *MockInstance) Migrate() error {
	return nil
}
