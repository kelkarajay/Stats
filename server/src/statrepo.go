package main

import (
	"time"
)

var stats Stats

// Give us some seed data
func init() {
	//RepoCreateStat(Stat{StatAppID: "TestApp", StatClientID: "Test Client", StatType: "Browser", StatCategory: "BrowserName", StatValue: "FF"})
}

// RepoCreateStat - Creates and records the stat in the repo
func RepoCreateStat(st Stat) Stat {
	st.StatTime = time.Now()
	stats = append(stats, st)
	return st
}
