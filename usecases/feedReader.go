package usecases

import "github.com/janithl/paataka/entities"

// FeedReader is an interface for reading Publication Posts through feeds
type FeedReader interface {
	Read(string) []entities.Post
}
