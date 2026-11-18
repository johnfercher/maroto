package gopdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gopdf/gopdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	pointsPerInch      = 72.0
	millimetersPerInch = 25.4
)

type Font struct {
	pdf         gopdfwrapper.Gopdf
	size        float64
	family      string
	style       fontstyle.Type
	scaleFactor float64
	fontColor   *props.Color
}

// NewFont create a Font.
func NewFont(pdf gopdfwrapper.Gopdf, size float64, family string, style fontstyle.Type) *Font {
	pdf.SetFont(family, style, size)

	return &Font{
		pdf:         pdf,
		size:        size,
		family:      family,
		style:       style,
		scaleFactor: pointsPerInch / millimetersPerInch, // font sizes are in points, the document unit is millimeters
		fontColor:   &props.Color{Red: 0, Green: 0, Blue: 0},
	}
}

// GetFamily return the currently Font family configured.
func (s *Font) GetFamily() string {
	return s.family
}

// GetStyle return the currently Font style configured.
func (s *Font) GetStyle() fontstyle.Type {
	return s.style
}

// GetSize return the currently Font size configured.
func (s *Font) GetSize() float64 {
	return s.size
}

// GetFont return all the currently Font properties configured.
func (s *Font) GetFont() (string, fontstyle.Type, float64) {
	return s.family, s.style, s.size
}

func (s *Font) GetHeight(family string, style fontstyle.Type, size float64) float64 {
	s.SetFont(family, style, size)
	_, _, fontSize := s.GetFont()
	return fontSize / s.scaleFactor
}

// SetFamily defines a new Font family.
func (s *Font) SetFamily(family string) {
	s.family = family

	s.pdf.SetFont(s.family, s.style, s.size)
}

// SetStyle defines a new Font style.
func (s *Font) SetStyle(style fontstyle.Type) {
	s.style = style

	s.pdf.SetFont(s.family, s.style, s.size)
}

// SetSize defines a new Font size.
func (s *Font) SetSize(size float64) {
	s.size = size
	s.pdf.SetFont(s.family, s.style, s.size)
}

// SetFont defines all new Font properties.
func (s *Font) SetFont(family string, style fontstyle.Type, size float64) {
	s.family = family
	s.style = style
	s.size = size

	s.pdf.SetFont(s.family, s.style, s.size)
}

func (s *Font) SetColor(color *props.Color) {
	if color == nil {
		return
	}

	s.fontColor = color
	s.pdf.SetTextColor(color.Red, color.Green, color.Blue)
}

func (s *Font) GetColor() *props.Color {
	return s.fontColor
}
