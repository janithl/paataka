package domain

// Publications is an interface for repositories storing Publications
type Publications interface {
	GetVersion() string
}
