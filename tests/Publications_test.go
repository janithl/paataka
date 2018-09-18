package tests

import (
	"testing"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
)

func TestPublications(t *testing.T) {
	version := "Mock SQLRepository v1.0"
	repo := database.NewSQLPublicationRepository(version)
	service := domain.NewPublicationServiceImpl(repo)

	t.Run("Check Version", func(t *testing.T) {
		got := service.GetRepositoryVersion()
		want := version

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
