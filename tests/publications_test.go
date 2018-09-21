package tests

import (
	"reflect"
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

	t.Run("Add 3 publications and then list them all", func(t *testing.T) {
		publications := make(map[string]domain.Publication)

		publications["pub-001"] = domain.Publication{}
		publications["pub-002"] = domain.Publication{}
		publications["pub-003"] = domain.Publication{}

		service.Add(publications["pub-001"])
		service.Add(publications["pub-002"])
		service.Add(publications["pub-003"])

		got := service.ListAll()
		want := publications

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
