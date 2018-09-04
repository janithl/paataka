package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/janithl/paataka/common"
)

func initLog() {
	file, err := os.OpenFile("./data/application.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)
}

func main() {
	// need to move this to .env
	appVersion := "v1.0.0"

	// init the logfile
	initLog()

	fmt.Printf("Paataka %s\n", appVersion)

	pub := common.Publication{Name: "Daily Mirror"}
	pub.SetLink("http://www.dailymirror.lk/RSS_Feeds/breaking-news")
	pub.Fetch()

	appState := common.AppState{Timestamp: time.Now(), AppVersion: appVersion}
	appState.AddPublication(pub)
	appState.Store()
}
