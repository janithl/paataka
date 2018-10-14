package entities

// Truncate truncates a given string to a given length
func Truncate(text string, length int) string {
	if len(text) <= length || length < 4 {
		return text
	}

	return text[:length-3] + "..."
}
