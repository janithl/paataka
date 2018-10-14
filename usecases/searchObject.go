package usecases

// SearchObject entity type
type SearchObject struct {
	ID      string
	Type    string
	Content string
	Terms   map[string]float64
	Score   float64
}
