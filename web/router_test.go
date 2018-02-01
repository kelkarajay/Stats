package web

import "testing"

func TestNewRouterInit(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("Router not setup")
	}
}
