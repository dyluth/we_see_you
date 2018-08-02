package analyser

// fancy word for a bag of regexs..

// Analyser is the front end, you give it a message, and it returns with the category its in.
type Analyser interface {
	// GetCategory takes the string, and returns a category name, combined with its confidence (0-100)
	GetCategory(message string) (cat string, confidence int)
}
