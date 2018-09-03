package main

import (
	"encoding/json"
	"time"
)

// AppState struct to hold application state
type AppState struct {
	Timestamp  time.Time
	AppVersion string
}

// Store stores application state to persistance
func (a *AppState) Store() {
	jsonContent, _ := json.Marshal(a)

	p := persist{}
	p.WriteFile("./data/appState.json", string(jsonContent))
}
