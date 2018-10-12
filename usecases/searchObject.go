package usecases

import "fmt"

// SearchObject entity type
type SearchObject struct {
	ID      string
	Type    string
	Content string
}

func (s SearchObject) String() string {
	return fmt.Sprintf("[%s] %s", s.Type, s.Content)
}
