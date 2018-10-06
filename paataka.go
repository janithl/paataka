package main

import (
	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/ui"
	"github.com/janithl/paataka/usecases"
)

func main() {
	paatakaVersion := "Paataka v1.0"
	repo := database.NewSQLPublicationRepository(paatakaVersion)
	service := usecases.NewPublicationServiceImpl(repo, nil)

	cli := ui.CLI{PublicationService: service}
	cli.GetInput()
}
