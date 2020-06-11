package internal

import (
	"github.com/Vale-sail/maroto/pkg/consts"
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
	GetScaleFactor() (scaleFactor float64)
}

type font struct {
	pdf         gofpdf.Pdf
	size        float64
	family      consts.Family
	style       consts.Style
	scaleFactor float64
}

// NewFont create a Font
func NewFont(pdf gofpdf.Pdf, size float64, family consts.Family, style consts.Style) *font {
	return &font{
		pdf,
		size,
		family,
		style,
		72.0 / 25.4, // Value defined inside gofpdf constructor
	}
}

// GetFamily return the currently Font family configured
func (s *font) GetFamily() consts.Family {
	return s.family
}

// GetStyle return the currently Font style configured
func (s *font) GetStyle() consts.Style {
	return s.style
}

// GetSize return the currently Font size configured
func (s *font) GetSize() float64 {
	return s.size
}

// GetFont return all the currently Font properties configured
func (s *font) GetFont() (consts.Family, consts.Style, float64) {
	return s.family, s.style, s.size
}

// SetFamily defines a new Font family
func (s *font) SetFamily(family consts.Family) {
	s.family = family

	s.pdf.SetFont(string(s.family), string(s.style), s.size)
}

// SetStyle defines a new Font style
func (s *font) SetStyle(style consts.Style) {
	s.style = style

	s.pdf.SetFontStyle(string(s.style))
}

// SetSize defines a new Font size
func (s *font) SetSize(size float64) {
	s.size = size
	s.pdf.SetFontSize(s.size)
}

// SetFont defines all new Font properties
func (s *font) SetFont(family consts.Family, style consts.Style, size float64) {
	s.family = family
	s.style = style
	s.size = size

	s.pdf.SetFont(string(s.family), string(s.style), s.size)
}

// GetScaleFactor retrieve the scale factor defined in the instantiation of gofpdf
func (s *font) GetScaleFactor() (scaleFactor float64) {
	return s.scaleFactor
}
