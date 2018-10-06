package usecases_test

import (
	"reflect"
	"testing"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/entities"
	"github.com/janithl/paataka/usecases"
)

/** MockFeedReader */
type MockFeedReader struct {
	Posts map[string]entities.Post
}

func (m MockFeedReader) Read(url string) map[string]entities.Post {
	return m.Posts
}

func setupService(version string, reader usecases.FeedReader) *usecases.PublicationServiceImpl {
	repo := database.NewInMemoryPublicationRepository(version)
	return usecases.NewPublicationServiceImpl(repo, reader)
}

const version string = "Mock InMemoryRepository v1.0"

var mockReader = MockFeedReader{Posts: nil}

/*
	given PublicationService version is Mock InMemoryRepository v1.0
	when GetRepositoryVersion is called
	then Mock InMemoryRepository v1.0 is returned
*/
func TestPublicationVersion(t *testing.T) {
	service := setupService(version, mockReader)

	got := service.GetRepositoryVersion()
	want := version

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestPublicationAddAndList(t *testing.T) {
	service := setupService(version, mockReader)
	publications := make(map[string]entities.Publication)

	publications["pub-001"] = entities.Publication{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
	publications["pub-002"] = entities.Publication{ID: "pub-002", Title: "Ben's Thoughts", URL: "https://ben-bert.me"}
	publications["pub-003"] = entities.Publication{ID: "pub-003", Title: "Cambrian Technical Group", URL: "http://blog.cambrian.tech"}

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
}

/*
	given PublicationService with no publications
	when we try to find "pub-001"
	then PublicationNotFound error is thrown
*/
func TestPublicationFindFail(t *testing.T) {
	service := setupService(version, mockReader)
	if _, err := service.Find("pub-001"); err == nil {
		t.Errorf("got '%s' want '%s'", err, usecases.ErrPublicationNotFound)
	}
}

func TestPublicationFindAndUpdate(t *testing.T) {
	service := setupService(version, mockReader)

	// Initial add
	publication := entities.Publication{ID: "pub-010", Title: "Greenland Business Digest", URL: "https://gbd.org"}
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
}

/*
	given
		3 publications in the repository:
			pub 1 - fetched 1 hour ago
			pub 2 - fetched 30 minutes ago
			pub 3 - fetched 3 hours ago
	and
		fetch older than is 1 hour
	and
		fetch at a time limit is 10

	when GetFetchable is called

	then pub 1 and pub 3 should be returned
*/

/*
	given publication is in repository:
		{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
	and publication's feed has 3 posts on it:
		{ID: "100-001", Title: "Hello World", URL: "https://alberta.ca/blog/001/hello-world"}
		{ID: "100-002", Title: "Yesterday", URL: "https://alberta.ca/blog/002/yesterday"}
		{ID: "100-003", Title: "Another Day", URL: "https://alberta.ca/blog/003/another-day"}
	when AddPublicationPosts is called on it
	then the Posts should be added to the repository
*/
func TestFetchPublicationPostsAddAndListAll(t *testing.T) {
	mockFeedReader := MockFeedReader{}

	// create a map of posts
	posts := make(map[string]entities.Post)
	posts["100-001"] = entities.Post{ID: "100-001", Title: "Hello World", URL: "https://alberta.ca/blog/001/hello-world"}
	posts["100-002"] = entities.Post{ID: "100-002", Title: "Yesterday", URL: "https://alberta.ca/blog/002/yesterday"}
	posts["100-003"] = entities.Post{ID: "100-003", Title: "Another Day", URL: "https://alberta.ca/blog/003/another-day"}

	// assign it to the mock FeedReader
	mockFeedReader.Posts = posts

	// setup service
	service := setupService(version, mockFeedReader)

	// create new publication
	publication := entities.Publication{ID: "pub-001", Title: "Alberta Blog", URL: "https://alberta.ca/blog"}
	service.Add(publication)

	// fetch posts for the publication
	service.FetchPublicationPosts(publication)

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
}
