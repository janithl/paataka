package entities

import (
	"fmt"
	"time"
)

// Publication domain entity type
type Publication struct {
	ID        string
	Title     string
	URL       string
	Posts     map[string]Post
	AddedAt   time.Time
	FetchedAt time.Time
	UpdatedAt time.Time
}

// AddPost adds a Post to the Publication
func (p *Publication) AddPost(post Post) {
	if p.Posts == nil {
		p.Posts = map[string]Post{}
	}

	p.Posts[post.ID] = post
}

// String method returns the Publication details as a string
func (p *Publication) String() string {
	return fmt.Sprintf("%-20s %-48s %4d posts", p.Title, Truncate(p.URL, 48), len(p.Posts))
}
