package usecases

import "github.com/janithl/paataka/entities"

// PublicationServiceImpl is an implementation of the PublicationService
type PublicationServiceImpl struct {
	repository Publications
	reader     FeedReader
}

// NewPublicationServiceImpl returns a new instance of PublicationServiceImpl
func NewPublicationServiceImpl(repository Publications, reader FeedReader) *PublicationServiceImpl {
	return &PublicationServiceImpl{
		repository: repository,
		reader:     reader,
	}
}

// GetRepositoryVersion returns the version string of the repository used by the service
func (p *PublicationServiceImpl) GetRepositoryVersion() string {
	return p.repository.GetVersion()
}

// Add adds a new Publication
func (p *PublicationServiceImpl) Add(pub entities.Publication) string {
	if len(pub.Posts) == 0 {
		pub.Posts = make(map[string]entities.Post)
	}

	return p.repository.Add(pub)
}

// AddPost adds a new Post to the Publication
func (p *PublicationServiceImpl) AddPost(id string, post entities.Post) error {
	if pub, err := p.Find(id); err == nil {
		pub.Posts[post.ID] = post
		p.Add(pub)

		return nil
	}

	return ErrPublicationNotFound
}

// ListAll returns all the publications in a Map
func (p *PublicationServiceImpl) ListAll() map[string]entities.Publication {
	return p.repository.ListAll()
}

// Find returns the Publication by the given ID
func (p *PublicationServiceImpl) Find(id string) (entities.Publication, error) {
	if value, exists := p.repository.ListAll()[id]; exists {
		return value, nil
	}

	return entities.Publication{}, ErrPublicationNotFound
}
