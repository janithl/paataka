package usecases

// SearchService is an interface for a service that searches text content
type SearchService interface {
	Index(SearchObject)
	Search(string, string) ([]SearchObject, error)
}
