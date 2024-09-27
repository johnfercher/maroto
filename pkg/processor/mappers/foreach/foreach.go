package foreach

// ForEach is responsible for generating a list of Pages for each group of information
type ForEach[T interface{}] struct {
	// SourceKey defines the word that will be used to search the page content.
	SourceKey string

	// Pages define the group of pages that should be created for each content group found
	Pages []T
}
