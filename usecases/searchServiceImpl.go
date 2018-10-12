package usecases

import (
	"regexp"
	"strings"
)

// SearchServiceImpl is an implementation of the SearchService
type SearchServiceImpl struct {
	searchIndex map[string][]SearchObject
	wordRegexp  *regexp.Regexp
}

// NewSearchServiceImpl returns a new instance of SearchServiceImpl
func NewSearchServiceImpl() *SearchServiceImpl {
	return &SearchServiceImpl{
		searchIndex: make(map[string][]SearchObject),
		wordRegexp:  regexp.MustCompile(`\w+`),
	}
}

// Index adds the given object to the search index
func (s *SearchServiceImpl) Index(obj SearchObject) {
	matches := s.wordRegexp.FindAllString(strings.ToLower(obj.Content), -1)

	for _, match := range matches {
		s.searchIndex[match] = append(s.searchIndex[match], obj)
	}
}

// Search returns the list of relevant results
func (s *SearchServiceImpl) Search(objtype string, query string) ([]SearchObject, error) {
	matches := s.wordRegexp.FindAllString(strings.ToLower(query), -1)

	results := make([]SearchObject, 0)
	for _, match := range matches {
		for _, objmatch := range s.searchIndex[match] {
			if objmatch.Type == objtype {
				results = append(results, objmatch)
			}
		}
	}

	if len(results) == 0 {
		return results, ErrNoSearchResults
	}

	return results, nil
}
