// Package contains all line styles.
package linestyle

const (
	// DefaultLineThickness represents the default linestyle width in gofpdf.
	DefaultLineThickness float64 = 0.2
)

// Type is a representation of a linestyle style.
type Type string

const (
	// Solid represents a solid style.
	Solid Type = "solid"
	// Dashed represents a dashed style.
	Dashed Type = "dashed"
)
