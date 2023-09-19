package internal

import (
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

const (
	gofpdfFontScale1 = 72.0
	gofpdfFontScale2 = 25.4
)

// Font is the abstraction which deals of how to set fontstyle configurations.
type Font interface {
	SetFamily(family string)
	SetStyle(style fontstyle.Type)
	SetSize(size float64)
	SetFont(family string, style fontstyle.Type, size float64)
	GetFamily() string
	GetStyle() fontstyle.Type
	GetSize() float64
	GetFont() (string, fontstyle.Type, float64)
	GetScaleFactor() (scaleFactor float64)
	SetColor(color *color.Color)
	GetColor() *color.Color
}

type font struct {
	pdf         fpdf.Fpdf
	size        float64
	family      string
	style       fontstyle.Type
	scaleFactor float64
	fontColor   *color.Color
}

// NewFont create a Font.
func NewFont(pdf fpdf.Fpdf, size float64, family string, style fontstyle.Type) *font {
	pdf.SetFont(family, string(style), size)

	return &font{
		pdf:         pdf,
		size:        size,
		family:      family,
		style:       style,
		scaleFactor: gofpdfFontScale1 / gofpdfFontScale2, // Value defined inside gofpdf constructor,
		fontColor:   &color.Color{Red: 0, Green: 0, Blue: 0},
	}
}

// GetFamily return the currently Font family configured.
func (s *font) GetFamily() string {
	return s.family
}

// GetStyle return the currently Font style configured.
func (s *font) GetStyle() fontstyle.Type {
	return s.style
}

// GetSize return the currently Font size configured.
func (s *font) GetSize() float64 {
	return s.size
}

// GetFont return all the currently Font properties configured.
func (s *font) GetFont() (string, fontstyle.Type, float64) {
	return s.family, s.style, s.size
}

// SetFamily defines a new Font family.
func (s *font) SetFamily(family string) {
	s.family = family

	s.pdf.SetFont(s.family, string(s.style), s.size)
}

// SetStyle defines a new Font style.
func (s *font) SetStyle(style fontstyle.Type) {
	s.style = style

	s.pdf.SetFontStyle(string(s.style))
}

// SetSize defines a new Font size.
func (s *font) SetSize(size float64) {
	s.size = size
	s.pdf.SetFontSize(s.size)
}

// SetFont defines all new Font properties.
func (s *font) SetFont(family string, style fontstyle.Type, size float64) {
	s.family = family
	s.style = style
	s.size = size

	s.pdf.SetFont(s.family, string(s.style), s.size)
}

// GetScaleFactor retrieve the scale factor defined in the instantiation of gofpdf.
func (s *font) GetScaleFactor() (scaleFactor float64) {
	return s.scaleFactor
}

func (s *font) SetColor(color *color.Color) {
	if color == nil {
		return
	}

	s.fontColor = color
	s.pdf.SetTextColor(color.Red, color.Green, color.Blue)
}

func (s *font) GetColor() *color.Color {
	return s.fontColor
}
