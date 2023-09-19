package linestyle

const (
	// DefaultLineWidth represents the default line width in gofpdf.
	DefaultLineWidth float64 = 0.1
)

// Type is a representation of a line style.
type Type string

const (
	// Solid represents a solid style.
	Solid Type = "solid"
	// Dashed represents a dashed style.
	Dashed Type = "dashed"
	// Dotted represents a dotted style.
	Dotted Type = "dotted"
)
