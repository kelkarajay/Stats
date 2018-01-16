package web

import (
	"time"

	"github.com/Xivolkar/Stats/model"
)

var stats model.Stats

// Give us some seed data
func init() {
	//RepoCreateStat(Stat{StatAppID: "TestApp", StatClientID: "Test Client", StatType: "Browser", StatCategory: "BrowserName", StatValue: "FF"})
}

// RepoCreateStat - Creates and records the stat in the repo
func RepoCreateStat(st model.Stat) model.Stat {
	st.StatTime = time.Now()
	stats = append(stats, st)
	return st
}
