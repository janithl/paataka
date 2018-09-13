package domain

// PublicationService exposes the Publication business logic to outside the domain
type PublicationService interface {
	CreatePublication(*Publication) int
	RetrievePublication(int) *Publication
	// UpdatePublicationName(*Publication, string) bool
	// UpdatePublicationURL(*Publication, string) bool
	// AddPublicationPost(*Publication, *Post) bool
	// DeletePublication(int) bool
}
