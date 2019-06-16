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
func (self *font) GetFamily() Family {
	return self.family
}

// Get the currently Font style configured
func (self *font) GetStyle() Style {
	return self.style
}

// Get the currently Font size configured
func (self *font) GetSize() float64 {
	return self.size
}

// Get all the currently Font properties configured
func (self *font) GetFont() (Family, Style, float64) {
	return self.family, self.style, self.size
}

// Set the Font family
func (self *font) SetFamily(family Family) {
	self.family = family

	self.pdf.SetFont(string(self.family), string(self.style), self.size)
}

// Set the Font style
func (self *font) SetStyle(style Style) {
	self.style = style

	self.pdf.SetFontStyle(string(self.style))
}

// Set the Font size
func (self *font) SetSize(size float64) {
	self.size = size
	self.pdf.SetFontSize(self.size)
}

// Set all the Font properties
func (self *font) SetFont(family Family, style Style, size float64) {
	self.family = family
	self.style = style
	self.size = size

	self.pdf.SetFont(string(self.family), string(self.style), self.size)
}
