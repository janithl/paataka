package main

import (
	"fmt"
	"time"
)

func main() {
	// need to move this to .env
	appVersion := "v1.0.0"

	fmt.Printf("Paataka %s\n", appVersion)

	appState := AppState{Timestamp: time.Now(), AppVersion: appVersion}
	appState.Store()
}
