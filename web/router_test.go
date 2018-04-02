package web

import (
	"testing"

	"github.com/Xivolkar/Stats/app"
)

func TestNewRouterInit(t *testing.T) {
	ctx := app.CreateContextForTestSetup()
	router := NewRouter(ctx)
	if router == nil {
		t.Error("Router not setup")
	}
}
