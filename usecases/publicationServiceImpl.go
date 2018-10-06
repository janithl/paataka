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
	return p.repository.Add(pub)
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

// FetchPublicationPosts fetches Publication posts from the feed and adds it to the publication
func (p *PublicationServiceImpl) FetchPublicationPosts(pub entities.Publication) error {
	posts := p.reader.Read(pub.URL)
	for _, post := range posts {
		pub.AddPost(post)
	}

	p.Add(pub)
	return nil
}
