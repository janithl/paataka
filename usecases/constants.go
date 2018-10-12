package usecases

import "errors"

// ErrPublicationNotFound is thrown when the given publication ID cannot be found
var ErrPublicationNotFound = errors.New("Publication not found")

// ErrNoSearchResults is thrown when the given query returns no results
var ErrNoSearchResults = errors.New("No results for given search")
