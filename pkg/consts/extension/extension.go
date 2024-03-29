// Package extension contains all image extensions.
package extension

// Type is a representation of a Image extension.
type Type string

const (
	// Jpg represents a jpg extension.
	Jpg Type = "jpg"
	// Jpeg represents a jpeg extension.
	Jpeg Type = "jpeg"
	// Png represents a png extension.
	Png Type = "png"
)

// IsValid checks if the extension is valid.
func (t Type) IsValid() bool {
	return t == Jpg || t == Jpeg || t == Png
}
