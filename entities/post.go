package entities

import (
	"fmt"
	"time"
)

// Post domain entity type
type Post struct {
	ID        string
	Title     string
	URL       string
	AddedAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// String method returns the Post details as a string
func (p *Post) String() string {
	return fmt.Sprintf("%-60s %19s", Truncate(p.Title, 60), p.CreatedAt.Format("2006-01-02 03:04PM"))
}
