package domain

import "time"

// Post is a struct that holds info about blog/site posts
type Post struct {
	id          int
	Title       string
	URL         string
	Content     string
	publishedAt time.Time
	addedAt     time.Time
	updatedAt   time.Time
}
