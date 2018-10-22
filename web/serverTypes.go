package web

// Publication type for communicating JSON of Publications
type Publication struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	PostCount int    `json:"post_count"`
}

// PublicationList type for managing lists of Publications
type PublicationList struct {
	Publications []Publication `json:"publications"`
	Page         int           `json:"page"`
}
