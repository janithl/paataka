package domain

import "time"

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
