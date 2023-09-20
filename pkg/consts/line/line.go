package line

const (
	// DefaultLineWidth represents the default line width in gofpdf.
	DefaultLineWidth float64 = 0.1
)

// Style is a representation of a line style.
type Style string

const (
	// Solid represents a solid style.
	Solid Style = "solid"
	// Dashed represents a dashed style.
	Dashed Style = "dashed"
)
