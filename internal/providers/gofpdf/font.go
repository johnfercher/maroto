package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	gofpdfFontScale1 = 72.0
	gofpdfFontScale2 = 25.4
)

type font struct {
	pdf         gofpdfwrapper.Fpdf
	size        float64
	family      string
	style       fontstyle.Type
	scaleFactor float64
	fontColor   *props.Color
}

// NewFont create a Font.
func NewFont(pdf gofpdfwrapper.Fpdf, size float64, family string, style fontstyle.Type) *font {
	pdf.SetFont(family, string(style), size)

	return &font{
		pdf:         pdf,
		size:        size,
		family:      family,
		style:       style,
		scaleFactor: gofpdfFontScale1 / gofpdfFontScale2, // Bytes defined inside gofpdf constructor,
		fontColor:   &props.Color{Red: 0, Green: 0, Blue: 0},
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

func (s *font) GetHeight(family string, style fontstyle.Type, size float64) float64 {
	s.SetFont(family, style, size)
	_, _, fontSize := s.GetFont()
	return fontSize / s.scaleFactor
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

func (s *font) SetColor(color *props.Color) {
	if color == nil {
		return
	}

	s.fontColor = color
	s.pdf.SetTextColor(color.Red, color.Green, color.Blue)
}

func (s *font) GetColor() *props.Color {
	return s.fontColor
}
