package domain

import "errors"

// PublicationServiceImpl is an implementation of the PublicationService
type PublicationServiceImpl struct {
	repository Publications
}

// NewPublicationServiceImpl returns a new instance of PublicationServiceImpl
func NewPublicationServiceImpl(repository Publications) *PublicationServiceImpl {
	return &PublicationServiceImpl{
		repository: repository,
	}
}

// GetRepositoryVersion returns the version string of the repository used by the service
func (p *PublicationServiceImpl) GetRepositoryVersion() string {
	return p.repository.GetVersion()
}

// Add adds a new Publication
func (p *PublicationServiceImpl) Add(pub Publication) string {
	if len(pub.Posts) == 0 {
		pub.Posts = make(map[string]Post)
	}

	return p.repository.Add(pub)
}

// AddPost adds a new Post to the Publication
func (p *PublicationServiceImpl) AddPost(id string, post Post) error {
	if pub, err := p.Find(id); err == nil {
		pub.Posts[post.ID] = post
		p.Add(pub)

		return nil
	}

	return errors.New(ErrorPublicationNotFound)
}

// ListAll returns all the publications in a Map
func (p *PublicationServiceImpl) ListAll() map[string]Publication {
	return p.repository.ListAll()
}

// Find returns the Publication by the given ID
func (p *PublicationServiceImpl) Find(id string) (Publication, error) {
	if value, exists := p.repository.ListAll()[id]; exists {
		return value, nil
	}

	return Publication{}, errors.New(ErrorPublicationNotFound)
}
