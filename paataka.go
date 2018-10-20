package main

import (
	"flag"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/feedreader"
	"github.com/janithl/paataka/ui"
	"github.com/janithl/paataka/usecases"
	"github.com/janithl/paataka/web"
)

func main() {
	paatakaVersion := "Paataka v1.2"
	repo := database.NewSQLPublicationRepository(paatakaVersion)
	defer repo.Persist()

	reader := feedreader.XMLFeedReader{}
	search := usecases.NewSearchServiceImpl()
	service := usecases.NewPublicationServiceImpl(search, repo, reader)

	serveHTTP := flag.Bool("http", false, "serve http")
	flag.Parse()

	if *serveHTTP {
		server := web.Server{PublicationService: service}
		server.Serve()
	} else {
		cli := ui.CLI{PublicationService: service}
		cli.GetInput()
	}
}
