package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

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

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "index.html")
	http.ServeFile(w, r, fp)
}

// Serve serves HTTP
func (s *Server) Serve() {
	// define the routes
	http.HandleFunc("/publications", s.listPublications())
	http.HandleFunc("/publication/", s.listPublicationDetails())
	http.HandleFunc("/version", s.version())
	http.HandleFunc("/", s.defaultHandler)

	// serve on 9000
	fmt.Println("Serving HTTP on localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
