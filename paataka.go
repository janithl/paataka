package main

import (
	"fmt"
	"time"
)

func main() {
	// need to move this to .env
	appVersion := "v1.0.0"

	fmt.Printf("Paataka %s\n", appVersion)

	p := persist{}
	p.OpenFile()
	p.WriteFile(appState{Timestamp: time.Now(), AppVersion: appVersion})
	defer p.CloseFile()
}
