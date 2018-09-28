package domain_test

import (
	"reflect"
	"testing"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
)

func setupService(version string) *domain.PublicationServiceImpl {
	repo := database.NewInMemoryPublicationRepository(version)
	return domain.NewPublicationServiceImpl(repo)
}

const version string = "Mock InMemoryRepository v1.0"

func TestPublicationVersion(t *testing.T) {
	service := setupService(version)

	got := service.GetRepositoryVersion()
	want := version

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestPublicationAddAndList(t *testing.T) {
	service := setupService(version)
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
}

func TestPublicationFindFail(t *testing.T) {
	service := setupService(version)
	if _, err := service.Find("not-there"); err == nil {
		t.Errorf("got '%s' want '%s'", err, domain.ErrorPublicationNotFound)
	}
}

func TestPublicationFindAndUpdate(t *testing.T) {
	service := setupService(version)

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
}

func TestPublicationPostAddAndListAll(t *testing.T) {
	service := setupService(version)

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
}

func TestPublicationPostFindFail(t *testing.T) {
	service := setupService(version)

	post := domain.Post{ID: "100-008", Title: "Heyya Failure", URL: "https://bigwig.com/blog/heyya"}
	if err := service.AddPost("pub-404", post); err == nil {
		t.Errorf("got '%s' want '%s'", err, domain.ErrorPublicationNotFound)
	}
}
