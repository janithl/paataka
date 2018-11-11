// Package web contains implementations of the API, and belong to
// the infrastructure layer
package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/janithl/paataka/usecases"
)

// Server struct implements the REST API Server
type Server struct {
	PublicationService usecases.PublicationService
}

func (s *Server) outputJSON(w http.ResponseWriter, output interface{}) {
	outputJSON, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}

func (s *Server) version() http.HandlerFunc {
	version := make(map[string]string)
	version["Version"] = s.PublicationService.GetRepositoryVersion()

	return func(w http.ResponseWriter, r *http.Request) {
		s.outputJSON(w, version)
	}
}

func (s *Server) listPublications() http.HandlerFunc {
	pubs := s.PublicationService.ListAll()
	return func(w http.ResponseWriter, r *http.Request) {
		pubList := PublicationList{Page: 1}
		for _, pub := range pubs {
			pubList.Publications = append(pubList.Publications, Publication{
				ID:        pub.ID,
				Title:     pub.Title,
				URL:       pub.URL,
				PostCount: len(pub.Posts),
			})
		}

		s.outputJSON(w, pubList)
	}
}

func (s *Server) listPublicationDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/publication/"):]
		pub, err := s.PublicationService.Find(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		s.outputJSON(w, Publication{ID: pub.ID, Title: pub.Title, URL: pub.URL, PostCount: len(pub.Posts)})
	}
}

func (s *Server) listSearchResults() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, size, search := 0, 10, ""
		query := r.URL.Query()
		pageQuery, ok := query["page"]
		if ok {
			page, _ = strconv.Atoi(pageQuery[0])
		}

		sizeQuery, ok := query["size"]
		if ok {
			size, _ = strconv.Atoi(sizeQuery[0])
		}

		searchQuery, ok := query["q"]
		if ok {
			search = searchQuery[0]
		}

		posts := PostList{Page: page, PageSize: size}
		if results := s.PublicationService.Search("Post", search); len(results) > 0 {
			ids := make([]string, 0)
			for _, res := range results {
				ids = append(ids, res.ID)
			}

			postObjects := s.PublicationService.GetPosts(ids)

			start := page * size
			end := (page + 1) * size

			if start > len(postObjects) || page < 0 || size < 1 {
				http.Error(w, "Incorrect Page or Size values", http.StatusBadRequest)
				return
			} else if end > len(postObjects) {
				end = len(postObjects)
			}

			for _, post := range postObjects[start:end] {
				posts.Posts = append(posts.Posts, Post{ID: post.ID, Title: post.Title, URL: post.URL})
			}
			posts.TotalSize = len(results)
		}

		s.outputJSON(w, posts)
	}
}

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "index.html")
	http.ServeFile(w, r, fp)
}

// Serve serves HTTP
func (s *Server) Serve() {
	// define the routes
	http.HandleFunc("/publications", s.listPublications())
	http.HandleFunc("/publication/", s.listPublicationDetails())
	http.HandleFunc("/search", s.listSearchResults())
	http.HandleFunc("/version", s.version())
	http.HandleFunc("/", s.defaultHandler)

	// serve on 9000
	fmt.Println("Serving HTTP on localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
