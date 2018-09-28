package tests

import (
	"reflect"
	"testing"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
)

func TestPosts(t *testing.T) {
	version := "Mock InMemoryRepository v1.0"
	repo := database.NewInMemoryPublicationRepository(version)
	service := domain.NewPublicationServiceImpl(repo)

	t.Run("Add 3 posts to publication and then list them all", func(t *testing.T) {
		// create new publication
		publication := domain.Publication{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
		service.Add(publication)

		// create a map of posts
		posts := make(map[string]domain.Post)
		posts["100-001"] = domain.Post{ID: "100-001", Title: "Hello World", URL: "https://alberta.ca/blog/001/hello-world"}
		posts["100-002"] = domain.Post{ID: "100-002", Title: "Yesterday", URL: "https://alberta.ca/blog/002/yesterday"}
		posts["100-003"] = domain.Post{ID: "100-003", Title: "Another Day", URL: "https://alberta.ca/blog/003/another-day"}

		// add them one-by-one to publication, checking for errors
		if err := service.AddPost("pub-001", posts["100-001"]); err != nil {
			t.Error(err)
		}

		if err := service.AddPost("pub-001", posts["100-002"]); err != nil {
			t.Error(err)
		}

		if err := service.AddPost("pub-001", posts["100-003"]); err != nil {
			t.Error(err)
		}

		// find the publication, checking for errors
		if pub, err := service.Find(publication.ID); err != nil {
			t.Error(err)
		} else {
			// make sure all the posts have been added in
			got := pub.Posts
			want := posts

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got '%s' want '%s'", got, want)
			}
		}
	})

	t.Run("Adding a post to a nonexistent publication should fail", func(t *testing.T) {
		post := domain.Post{ID: "100-008", Title: "Heyya Failure", URL: "https://bigwig.com/blog/heyya"}
		if err := service.AddPost("pub-404", post); err == nil {
			t.Errorf("got '%s' want '%s'", err, domain.ErrorPublicationNotFound)
		}
	})
}
