package main

import (
	"fmt"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
)

func main() {
	paatakaVersion := "Paataka v1.0"
	repo := database.NewSQLPublicationRepository(paatakaVersion)
	service := domain.NewPublicationServiceImpl(repo)

	fmt.Println(service.GetRepositoryVersion())
}
