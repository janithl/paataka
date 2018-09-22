package database

import "github.com/janithl/paataka/domain"

// SQLPublicationRepository is an implementation of Posts repository currently using an in-memory store
type SQLPublicationRepository struct {
	version      string
	publications map[string]domain.Publication
}

// NewSQLPublicationRepository returns a new SQLPublicationRepository
func NewSQLPublicationRepository(version string) *SQLPublicationRepository {
	return &SQLPublicationRepository{
		version:      version,
		publications: make(map[string]domain.Publication),
	}
}

// GetVersion returns the version string of the repository
func (s *SQLPublicationRepository) GetVersion() string {
	return s.version
}

// Add adds a new Publication
func (s *SQLPublicationRepository) Add(pub domain.Publication) string {
	s.publications[pub.ID] = pub
	return pub.ID
}

// ListAll returns all the publications in a Map
func (s *SQLPublicationRepository) ListAll() map[string]domain.Publication {
	return s.publications
}
