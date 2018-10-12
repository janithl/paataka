package usecases_test

import (
	"reflect"
	"testing"

	"github.com/janithl/paataka/usecases"
)

func TestSearch(t *testing.T) {
	service := usecases.NewSearchServiceImpl()

	objects := []usecases.SearchObject{
		usecases.SearchObject{ID: "100", Type: "Post", Content: "Hello"},
		usecases.SearchObject{ID: "101", Type: "Post", Content: "World"},
		usecases.SearchObject{ID: "102", Type: "Post", Content: "Hello World"},
		usecases.SearchObject{ID: "AAE", Type: "Publication", Content: "Hello Japan"},
		usecases.SearchObject{ID: "AAF", Type: "Publication", Content: "World Police"},
	}

	for _, obj := range objects {
		service.Index(obj)
	}

	t.Run("Search for all Posts containing 'Hello'", func(t *testing.T) {
		got, _ := service.Search("Post", "Hello")
		want := []usecases.SearchObject{objects[0], objects[2]}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q mathces, want %q", got, want)
		}
	})

	t.Run("Search for all Publications containing 'World'", func(t *testing.T) {
		got, _ := service.Search("Publication", "World")
		want := []usecases.SearchObject{objects[4]}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q mathces, want %q", got, want)
		}
	})

	t.Run("Search for a nonexistent word -- should throw error", func(t *testing.T) {
		if _, err := service.Search("Publication", "Not There"); err == nil {
			t.Errorf("got %q want %q", err, usecases.ErrNoSearchResults)
		}
	})

}
