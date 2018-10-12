package usecases

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/janithl/paataka/entities"
)

// PublicationServiceImpl is an implementation of the PublicationService
type PublicationServiceImpl struct {
	repository Publications
	reader     FeedReader

	random *rand.Rand
	search SearchService
}

// NewPublicationServiceImpl returns a new instance of PublicationServiceImpl
func NewPublicationServiceImpl(search SearchService, repository Publications, reader FeedReader) *PublicationServiceImpl {
	source := rand.NewSource(time.Now().UnixNano())
	return &PublicationServiceImpl{
		repository: repository,
		reader:     reader,
		random:     rand.New(source),
		search:     search,
	}
}

// GetRepositoryVersion returns the version string of the repository used by the service
func (p *PublicationServiceImpl) GetRepositoryVersion() string {
	return p.repository.GetVersion()
}

// Add adds a new Publication
func (p *PublicationServiceImpl) Add(pub entities.Publication) string {
	if pub.ID == "" {
		pub.ID = p.generateID()
	}
	if pub.AddedAt.IsZero() {
		pub.AddedAt = time.Now()
	}
	pub.UpdatedAt = time.Now()
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
		pub = p.addUniquePost(pub, post)
	}

	pub.FetchedAt = time.Now()
	p.Add(pub)
	return nil
}

func (p *PublicationServiceImpl) generateID() string {
	return fmt.Sprintf("%x-%x-%x", p.random.Uint32(), p.random.Uint32(), p.random.Uint32())
}

func (p *PublicationServiceImpl) addUniquePost(pub entities.Publication, post entities.Post) entities.Publication {
	for _, item := range pub.Posts {
		if item.URL == post.URL {
			return pub
		}
	}

	if post.ID == "" {
		post.ID = p.generateID()
	}
	post.AddedAt = time.Now()
	post.UpdatedAt = time.Now()

	pub.AddPost(post)

	return pub
}
