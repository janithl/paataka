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
		publication := domain.Publication{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
		service.Add(publication)

		post1 := domain.Post{ID: "100-001", Title: "Hello World", URL: "https://alberta.ca/blog/001/hello-world"}
		post2 := domain.Post{ID: "100-002", Title: "Yesterday", URL: "https://alberta.ca/blog/002/yesterday"}
		post3 := domain.Post{ID: "100-003", Title: "Another Day", URL: "https://alberta.ca/blog/003/another-day"}

		service.AddPost("pub-001", post1)
		service.AddPost("pub-001", post2)
		service.AddPost("pub-001", post3)

		posts := make([]domain.Post, 3)
		posts = append(posts, post1)
		posts = append(posts, post2)
		posts = append(posts, post3)

		got := service.Find(publication.ID)
		want := posts

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
