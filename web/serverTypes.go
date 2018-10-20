package web

// Publication type for communicating JSON of Publications
type Publication struct {
	ID, Title, URL string
	Posts          int
}

// PublicationList type for managing lists of Publications
type PublicationList struct {
	Publications []Publication
	Page         int
}
