package database

import "github.com/janithl/paataka/entities"

// InMemoryPublicationRepository is an implementation of Posts repository using an in-memory store
type InMemoryPublicationRepository struct {
	version      string
	publications map[string]entities.Publication
}

// NewInMemoryPublicationRepository returns a new InMemoryPublicationRepository
func NewInMemoryPublicationRepository(version string) *InMemoryPublicationRepository {
	return &InMemoryPublicationRepository{
		version:      version,
		publications: make(map[string]entities.Publication),
	}
}

// GetVersion returns the version string of the repository
func (s *InMemoryPublicationRepository) GetVersion() string {
	return s.version
}

// Add adds a new Publication
func (s *InMemoryPublicationRepository) Add(pub entities.Publication) string {
	s.publications[pub.ID] = pub
	return pub.ID
}

// ListAll returns all the publications in a Map
func (s *InMemoryPublicationRepository) ListAll() map[string]entities.Publication {
	return s.publications
}
