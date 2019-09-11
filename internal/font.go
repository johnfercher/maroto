package internal

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/jung-kurt/gofpdf"
)

// Font is the abstraction which deals of how to set font configurations
type Font interface {
	SetFamily(family consts.Family)
	SetStyle(style consts.Style)
	SetSize(size float64)
	SetFont(family consts.Family, style consts.Style, size float64)
	GetFamily() consts.Family
	GetStyle() consts.Style
	GetSize() float64
	GetFont() (consts.Family, consts.Style, float64)
}

type font struct {
	pdf    gofpdf.Pdf
	size   float64
	family consts.Family
	style  consts.Style
}

// NewFont create a Font
func NewFont(pdf gofpdf.Pdf, size float64, family consts.Family, style consts.Style) *font {
	return &font{
		pdf,
		size,
		family,
		style,
	}
}

// GetFamily return the currently Font family configured
func (self *font) GetFamily() consts.Family {
	return self.family
}

// GetStyle return the currently Font style configured
func (self *font) GetStyle() consts.Style {
	return self.style
}

// GetSize return the currently Font size configured
func (self *font) GetSize() float64 {
	return self.size
}

// GetFont return all the currently Font properties configured
func (self *font) GetFont() (consts.Family, consts.Style, float64) {
	return self.family, self.style, self.size
}

// SetFamily defines a new Font family
func (self *font) SetFamily(family consts.Family) {
	self.family = family

	self.pdf.SetFont(string(self.family), string(self.style), self.size)
}

// SetStyle defines a new Font style
func (self *font) SetStyle(style consts.Style) {
	self.style = style

	self.pdf.SetFontStyle(string(self.style))
}

// SetSize defines a new Font size
func (self *font) SetSize(size float64) {
	self.size = size
	self.pdf.SetFontSize(self.size)
}

// SetFont defines all new Font properties
func (self *font) SetFont(family consts.Family, style consts.Style, size float64) {
	self.family = family
	self.style = style
	self.size = size

	self.pdf.SetFont(string(self.family), string(self.style), self.size)
}
