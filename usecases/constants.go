package usecases

import "errors"

// ErrPublicationNotFound is thrown when the given publication ID cannot be found
var ErrPublicationNotFound = errors.New("Publication not found")

// ScoreCutoff is the minimum score needed to be added to a search result
const ScoreCutoff = 0.04
