package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/janithl/paataka/entities"
)

// SQLPublicationRepository is an implementation of Posts repository currently using an in-memory store
type SQLPublicationRepository struct {
	version      string
	publications map[string]entities.Publication
	filename     string
}

// NewSQLPublicationRepository returns a new SQLPublicationRepository
func NewSQLPublicationRepository(version string) *SQLPublicationRepository {
	repo := &SQLPublicationRepository{
		version:      version,
		publications: make(map[string]entities.Publication),
		filename:     "./database/appstate.json",
	}
	repo.readFromPersistance()
	return repo
}

// GetVersion returns the version string of the repository
func (s *SQLPublicationRepository) GetVersion() string {
	return s.version
}

// Add adds a new Publication
func (s *SQLPublicationRepository) Add(pub entities.Publication) string {
	s.publications[pub.ID] = pub
	return pub.ID
}

// ListAll returns all the publications in a Map
func (s *SQLPublicationRepository) ListAll() map[string]entities.Publication {
	return s.publications
}

// Persist persists repository state to filesystem
func (s *SQLPublicationRepository) Persist() {
	content, _ := json.MarshalIndent(s.publications, "", "  ")
	if err := ioutil.WriteFile(s.filename, content, 0644); err != nil {
		fmt.Println("Error: ", err)
	}
}

// readFromPersistance reads the repository state from the filesystem
func (s *SQLPublicationRepository) readFromPersistance() {
	if content, err := ioutil.ReadFile(s.filename); err != nil {
		fmt.Println("Error: ", err)
	} else {
		var data map[string]entities.Publication
		if err := json.Unmarshal(content, &data); err != nil {
			fmt.Println("Error: ", err)
		} else {
			s.publications = data
		}
	}
}
