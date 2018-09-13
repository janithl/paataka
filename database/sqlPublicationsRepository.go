package database

import "github.com/janithl/paataka/domain"

// SQLPublicationsRepository is a struct that implements Publications
type SQLPublicationsRepository struct {
	db string
}

// NewSQLPublicationsRepository returns a new instance of SQLPublicationsRepository
func NewSQLPublicationsRepository(database string) *SQLPublicationsRepository {
	return &SQLPublicationsRepository{
		db: database,
	}
}

// CreatePublication creates a new publication
func (sql *SQLPublicationsRepository) CreatePublication(pub *domain.Publication) int {
	return 1
}

// RetrievePublication returns a publication with the given ID
func (sql *SQLPublicationsRepository) RetrievePublication(id int) *domain.Publication {
	return &domain.Publication{}
}

// UpdatePublication updates the publication to db
func (sql *SQLPublicationsRepository) UpdatePublication(pub *domain.Publication) bool {
	return true
}

// DeletePublication removes the publication from db
func (sql *SQLPublicationsRepository) DeletePublication(id int) bool {
	return true
}
