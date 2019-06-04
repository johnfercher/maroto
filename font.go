package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

// Abstraction of Font configuration used in Maroto
type Font interface {
	SetFamily(family Family)
	SetStyle(style Style)
	SetSize(size float64)
	SetFont(family Family, style Style, size float64)
	GetFamily() Family
	GetStyle() Style
	GetSize() float64
	GetFont() (Family, Style, float64)
}

type font struct {
	pdf    gofpdf.Pdf
	size   float64
	family Family
	style  Style
}

// Create a Font configurator used in Maroto
func NewFont(pdf gofpdf.Pdf, size float64, family Family, style Style) Font {
	return &font{
		pdf,
		size,
		family,
		style,
	}
}

// Get the currently Font family configured
func (f *font) GetFamily() Family {
	return f.family
}

// Get the currently Font style configured
func (f *font) GetStyle() Style {
	return f.style
}

// Get the currently Font size configured
func (f *font) GetSize() float64 {
	return f.size
}

// Get all the currently Font properties configured
func (f *font) GetFont() (Family, Style, float64) {
	return f.family, f.style, f.size
}

// Set the Font family
func (f *font) SetFamily(family Family) {
	f.family = family

	f.pdf.SetFont(string(f.family), string(f.style), f.size)
}

// Set the Font style
func (f *font) SetStyle(style Style) {
	f.style = style

	f.pdf.SetFontStyle(string(f.style))
}

// Set the Font size
func (f *font) SetSize(size float64) {
	f.size = size
	f.pdf.SetFontSize(f.size)
}

// Set all the Font properties
func (f *font) SetFont(family Family, style Style, size float64) {
	f.family = family
	f.style = style
	f.size = size

	f.pdf.SetFont(string(f.family), string(f.style), f.size)
}
