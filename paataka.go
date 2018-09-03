package main

import (
	"fmt"
	"time"

	"github.com/janithl/paataka/common"
)

func main() {
	// need to move this to .env
	appVersion := "v1.0.0"

	fmt.Printf("Paataka %s\n", appVersion)

	pub := common.Publication{Name: "Daily Mirror"}
	pub.SetLink("http://www.dailymirror.lk/RSS_Feeds/breaking-news")
	pub.Fetch()

	appState := AppState{Timestamp: time.Now(), AppVersion: appVersion}
	appState.AddPublication(pub)
	appState.Store()
}
