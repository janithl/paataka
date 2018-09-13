package domain

// PublicationServiceImpl is a struct that implements a PublicationService
type PublicationServiceImpl struct {
	repo Publications
}

// NewPublicationServiceImpl returns a new instance of NewPublicationServiceImpl
func NewPublicationServiceImpl(publications Publications) *PublicationServiceImpl {
	return &PublicationServiceImpl{
		repo: publications,
	}
}

// CreatePublication creates new publication
func (service *PublicationServiceImpl) CreatePublication(p *Publication) int {
	return service.repo.CreatePublication(p)
}

// RetrievePublication gets publication by ID
func (service *PublicationServiceImpl) RetrievePublication(id int) *Publication {
	return service.repo.RetrievePublication(id)
}
