package common

import (
	"encoding/json"
	"time"
)

// AppState struct to hold application state
type AppState struct {
	Timestamp    time.Time
	AppVersion   string
	Publications []Publication
}

// AddPublication adds a publication to the app
func (a *AppState) AddPublication(pub Publication) {
	a.Publications = append(a.Publications, pub)
}

// Store stores application state to persistance
func (a *AppState) Store() {
	jsonContent, _ := json.MarshalIndent(a, "", "    ")

	p := Persist{}
	p.WriteFile("./data/AppState.json", string(jsonContent))
}
