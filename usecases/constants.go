package usecases

import "errors"

// ErrPublicationNotFound is thrown when the given publication ID cannot be found
var ErrPublicationNotFound = errors.New("Publication not found")