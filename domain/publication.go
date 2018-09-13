package domain

import "time"

// Publication is a struct that holds info about websites and blogs
type Publication struct {
	id         int
	Name       string
	URL        string
	addedAt    time.Time
	accessedAt time.Time
	updatedAt  time.Time
	Posts      []Post
}
