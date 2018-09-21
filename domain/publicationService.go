package domain

// PublicationService is an interface for service for working with Publications
type PublicationService interface {
	GetRepositoryVersion() string
	Add(Publication) string
	ListAll() map[string]Publication
}
