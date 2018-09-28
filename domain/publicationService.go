package domain

// PublicationService is an interface for service for working with Publications
type PublicationService interface {
	GetRepositoryVersion() string
	Add(Publication) string
	AddPost(string, Post)
	ListAll() map[string]Publication
	Find(string) (Publication, error)
}
