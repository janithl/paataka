package main

import (
	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/feedreader"
	"github.com/janithl/paataka/ui"
	"github.com/janithl/paataka/usecases"
)

func main() {
	paatakaVersion := "Paataka v1.0"
	repo := database.NewSQLPublicationRepository(paatakaVersion)
	defer repo.Persist()

	reader := feedreader.XMLFeedReader{}
	search := usecases.NewSearchServiceImpl()
	service := usecases.NewPublicationServiceImpl(search, repo, reader)

	cli := ui.CLI{PublicationService: service}
	cli.GetInput()
}
