package database

// SQLPublicationRepository is an implementation of Posts repository using SQLite
type SQLPublicationRepository struct {
	version string
}

// NewSQLPublicationRepository returns a new SQLPublicationRepository
func NewSQLPublicationRepository(version string) *SQLPublicationRepository {
	return &SQLPublicationRepository{
		version: version,
	}
}

// GetVersion returns the version string of the repository
func (s *SQLPublicationRepository) GetVersion() string {
	return s.version
}
