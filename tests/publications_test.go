package tests

import (
	"testing"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
)

func TestPublications(t *testing.T) {
	version := "Mock InMemoryRepository v1.0"
	repo := database.NewInMemoryPublicationRepository(version)
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

		publications["pub-001"] = domain.Publication{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
		publications["pub-002"] = domain.Publication{ID: "pub-002", Title: "Ben's Thoughts", URL: "https://ben-bert.me"}
		publications["pub-003"] = domain.Publication{ID: "pub-003", Title: "Cambrian Technical Group", URL: "http://blog.cambrian.tech"}

		service.Add(publications["pub-001"])
		service.Add(publications["pub-002"])
		service.Add(publications["pub-003"])

		got := service.ListAll()
		want := publications

		ids := [3]string{"pub-001", "pub-002", "pub-003"}
		for _, id := range ids {
			if got[id].ID != want[id].ID {
				t.Errorf("got '%s' want '%s'", got[id].ID, want[id].ID)
			}
			if got[id].Title != want[id].Title {
				t.Errorf("got '%s' want '%s'", got[id].Title, want[id].Title)
			}
			if got[id].URL != want[id].URL {
				t.Errorf("got '%s' want '%s'", got[id].URL, want[id].URL)
			}
		}
	})

	t.Run("Finding nonexistent publication should fail", func(t *testing.T) {
		if _, err := service.Find("not-there"); err == nil {
			t.Errorf("got '%s' want '%s'", err, domain.ErrorPublicationNotFound)
		}
	})

	t.Run("Finding existing publication and update it", func(t *testing.T) {
		// Initial add
		publication := domain.Publication{ID: "pub-010", Title: "Greenland Business Digest", URL: "https://gbd.org"}
		service.Add(publication)

		// Then find
		pub, err := service.Find("pub-010")
		if err != nil {
			t.Error(err)
		}

		// Verify
		got := pub.Title
		want := publication.Title
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		// Then update title, add (update) on service
		publication.Title = "Greenland Business Standard"
		service.Add(publication)

		// Find
		pub, err = service.Find("pub-010")
		if err != nil {
			t.Error(err)
		}

		// Verify
		got = pub.Title
		want = publication.Title
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
