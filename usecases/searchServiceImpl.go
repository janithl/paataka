package usecases

// SearchServiceImpl is an implementation of the SearchService
type SearchServiceImpl struct {
	searchIndex []SearchObject
}

// NewSearchServiceImpl returns a new instance of SearchServiceImpl
func NewSearchServiceImpl() *SearchServiceImpl {
	return &SearchServiceImpl{
		searchIndex: make([]SearchObject, 0),
	}
}

// Index adds the given object to the search index
func (s *SearchServiceImpl) Index(obj SearchObject) {
	s.searchIndex = append(s.searchIndex, obj)
}

// Search returns the list of relevant results
func (s *SearchServiceImpl) Search(objtype string, query string) ([]SearchObject, error) {
	return []SearchObject{}, nil
}
