package main

import (
	"encoding/json"
	"time"

	"github.com/janithl/paataka/common"
)

// AppState struct to hold application state
type AppState struct {
	Timestamp    time.Time
	AppVersion   string
	Publications []common.Publication
}

// AddPublication adds a publication to the app
func (a *AppState) AddPublication(pub common.Publication) {
	a.Publications = append(a.Publications, pub)
}

// Store stores application state to persistance
func (a *AppState) Store() {
	jsonContent, _ := json.Marshal(a)

	p := persist{}
	p.WriteFile("./data/appState.json", string(jsonContent))
}
