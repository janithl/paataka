package usecases

import "github.com/janithl/paataka/entities"

// PublicationService is an interface for service for working with Publications
type PublicationService interface {
	GetRepositoryVersion() string
	Add(entities.Publication) string
	ListAll() map[string]entities.Publication
	Find(string) (entities.Publication, error)
	FetchPublicationPosts(entities.Publication) error
}
