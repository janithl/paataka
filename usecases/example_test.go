package usecases_test

import (
	"fmt"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/usecases"
)

func ExampleNewPublicationServiceImpl() {
	repo := database.NewInMemoryPublicationRepository("Mock InMemoryRepository v1.0")
	search := usecases.NewSearchServiceImpl()
	reader := MockFeedReader{}
	pubService := usecases.NewPublicationServiceImpl(search, repo, reader)

	fmt.Println(pubService.GetRepositoryVersion())
	// Output: Mock InMemoryRepository v1.0
}
