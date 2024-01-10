// Package pagesize contains all default page sizes.
package pagesize

// Type is a representation of a page size.
type Type string

const (
	// A1 represents an A1 page size.
	A1 Type = "a1"
	// A2 represents an A2 page size.
	A2 Type = "a2"
	// A3 represents an A3 page size.
	A3 Type = "a3"
	// A4 represents an A4 page size.
	A4 Type = "a4"
	// A5 represents an A5 page size.
	A5 Type = "a5"
	// A6 represents an A6 page size.
	A6 Type = "a6"
	// Letter represents a letter page size.
	Letter Type = "letter"
	// Legal represents a legal page size.
	Legal Type = "legal"
	// Tabloid represents a tabloid page size.
	Tabloid Type = "tabloid"
	// DefaultTopMargin represents the default top margin in page size.
	DefaultTopMargin = 10.0
	// DefaultLeftMargin represents the default left margin in page size.
	DefaultLeftMargin = 10.0
	// DefaultRightMargin represents the default right margin in page size.
	DefaultRightMargin = 10.0
	// DefaultBottomMargin represents the default bottom margin in page size.
	DefaultBottomMargin = MinBottomMargin
	// MinTopMargin represents the minimum top margin in page size.
	MinTopMargin = 0.0
	// MinLeftMargin represents the minimum left margin in page size.
	MinLeftMargin = 0.0
	// MinRightMargin represents the minimum right margin in page size.
	MinRightMargin = 0.0
	// MinBottomMargin represents the minimum bottom margin in page size.
	MinBottomMargin = 20.0025
	// DefaultFontSize represents the default font size in page size.
	DefaultFontSize = 10.0
	// DefaultMaxGridSum represents the default max grid sum in page size.
	DefaultMaxGridSum = 12.0
)

// GetDimensions returns the width and height of the page size.
func GetDimensions(pageSize Type) (float64, float64) {
	switch pageSize {
	case A1:
		return 594.0, 841.0
	case A2:
		return 419.9, 594.0
	case A3:
		return 297.0, 419.9
	case A5:
		return 148.4, 210.0
	case A6:
		return 105.0, 148.5
	case Letter:
		return 215.9, 279.4
	case Legal:
		return 215.9, 355.6
	case Tabloid:
		return 279.4, 431.8
	default: // A4
		return 210.0, 297.0
	}
}
