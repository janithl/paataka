package domain

import "time"

// Post domain entity type
type Post struct {
	ID        string
	Title     string
	URL       string
	AddedAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
