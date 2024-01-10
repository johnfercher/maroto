package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

// Signature represents properties from a signature.
type Signature struct {
	// FontFamily of the text, ex: consts.Arial, helvetica and etc.
	FontFamily string
	// FontStyle of the text, ex: consts.Normal, bold and etc.
	FontStyle fontstyle.Type
	// FontSize of the text.
	FontSize float64
	// FontColor define the font color.
	FontColor *Color
	// LineColor define the line color.
	LineColor *Color
	// LineStyle define the line style (solid or dashed).
	LineStyle linestyle.Type
	// LineThickness define the line thickness.
	LineThickness float64
}

// ToMap returns a map with the Signature fields.
func (s *Signature) ToMap() map[string]interface{} {
	if s == nil {
		return nil
	}

	m := make(map[string]interface{})

	if s.FontFamily != "" {
		m["prop_font_family"] = s.FontFamily
	}

	if s.FontStyle != "" {
		m["prop_font_style"] = s.FontStyle
	}

	if s.FontSize != 0 {
		m["prop_font_size"] = s.FontSize
	}

	if s.LineStyle != "" {
		m["prop_line_style"] = s.LineStyle
	}

	if s.LineThickness != 0 {
		m["prop_line_thickness"] = s.LineThickness
	}

	if s.FontColor != nil {
		m["prop_font_color"] = s.FontColor.ToString()
	}

	if s.LineColor != nil {
		m["prop_line_color"] = s.LineColor.ToString()
	}

	return m
}

// MakeValid from Signature define default values for a Signature.
func (s *Signature) MakeValid(defaultFontFamily string) {
	if s.FontFamily == "" {
		s.FontFamily = defaultFontFamily
	}

	if s.FontStyle == "" {
		s.FontStyle = fontstyle.Bold
	}

	if s.FontSize == 0.0 {
		s.FontSize = 8.0
	}

	if s.LineStyle == "" {
		s.LineStyle = linestyle.Solid
	}

	if s.LineThickness == 0 {
		s.LineThickness = linestyle.DefaultLineThickness
	}
}

// ToLineProp from Signature return a Line based on Signature.
func (s *Signature) ToLineProp(offsetPercent float64) *Line {
	line := &Line{
		Color:         s.LineColor,
		Style:         s.LineStyle,
		Thickness:     s.LineThickness,
		OffsetPercent: offsetPercent,
	}
	line.MakeValid()
	return line
}

// ToFontProp from Signature return a Font based on Signature.
func (s *Signature) ToFontProp() *Font {
	font := &Font{
		Family: s.FontFamily,
		Style:  s.FontStyle,
		Size:   s.FontSize,
		Color:  s.FontColor,
	}
	font.MakeValid(s.FontFamily)
	return font
}

// ToTextProp from Signature return a Text based on Signature.
func (s *Signature) ToTextProp(align align.Type, top float64, verticalPadding float64) *Text {
	font := s.ToFontProp()
	text := &Text{
		Family:          font.Family,
		Style:           font.Style,
		Size:            font.Size,
		Align:           align,
		Top:             top,
		VerticalPadding: verticalPadding,
		Color:           font.Color,
	}
	text.MakeValid(font)
	return text
}
