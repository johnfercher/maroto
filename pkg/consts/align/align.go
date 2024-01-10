// Package align contains all align types.
package align

// Type is arRepresentation of a column align.
type Type string

const (
	// Left represents a left horizontal align.
	Left Type = "L"
	// Right represents a right horizontal align.
	Right Type = "R"
	// Center represents a center horizontal and/or vertical align.
	Center Type = "C"
	// Top represents a top vertical align.
	Top Type = "T"
	// Bottom represents a bottom vertical align.
	Bottom Type = "B"
	// Middle represents a middle align (from gofpdf).
	Middle Type = "M"
)
