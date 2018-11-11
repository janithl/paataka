package web

// Post type for communicating JSON of Posts
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// PostList type for managing lists of Posts
type PostList struct {
	Posts     []Post `json:"posts"`
	TotalSize int    `json:"total_size"`
	PageSize  int    `json:"page_size"`
	Page      int    `json:"page"`
}

// Publication type for communicating JSON of Publications
type Publication struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	PostCount int    `json:"post_count"`
	Posts     []Post `json:"posts"`
}

// PublicationList type for managing lists of Publications
type PublicationList struct {
	Publications []Publication `json:"publications"`
	Page         int           `json:"page"`
}
