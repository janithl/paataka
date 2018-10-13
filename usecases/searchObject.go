package usecases

import (
	"fmt"

	"github.com/janithl/paataka/entities"
)

// SearchObject entity type
type SearchObject struct {
	ID      string
	Type    string
	Content string
	Terms   map[string]float64
	Score   float64
}

// String method returns the SearchObject as a string
func (s *SearchObject) String() string {
	return fmt.Sprintf("%-60s [%s] %1.2f", entities.Truncate(s.Content, 60), entities.Truncate(s.Type, 14), s.Score)
}
