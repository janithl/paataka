package domain

// Publications exposes the repository interface expected by the business domain
type Publications interface {
	CreatePublication(*Publication) int
	RetrievePublication(int) *Publication
	UpdatePublication(*Publication) bool
	DeletePublication(int) bool
}
