package usecases

import "github.com/janithl/paataka/entities"

// Publications is an interface for repositories storing Publications
type Publications interface {
	GetVersion() string
	Add(entities.Publication) string
	ListAll() map[string]entities.Publication
}
