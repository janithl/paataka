package domain

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
	return "pub-uuid"
}

// ListAll returns all the publications in a Map
func (p *PublicationServiceImpl) ListAll() map[string]Publication {
	return make(map[string]Publication)
}
