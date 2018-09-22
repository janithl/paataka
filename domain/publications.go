package domain

// Publications is an interface for repositories storing Publications
type Publications interface {
	GetVersion() string
	Add(Publication) string
	ListAll() map[string]Publication
}
