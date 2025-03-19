package props

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

// SubText represents properties from a SubText inside a cell.
type SubText struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink *string
}

func (s *SubText) MakeValid(font *Font) {
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = font.Family
	}

	if s.Style == "" {
		s.Style = font.Style
	}

	if s.Size == undefinedValue {
		s.Size = font.Size
	}

	if s.Color == nil {
		s.Color = font.Color
	}
}

func (s *SubText) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if s.Family != "" {
		m["prop_font_family"] = s.Family
	}

	if s.Style != "" {
		m["prop_font_style"] = s.Style
	}

	if s.Size != 0 {
		m["prop_font_size"] = s.Size
	}

	if s.Color != nil {
		m["prop_color"] = s.Color.ToString()
	}

	if s.Hyperlink != nil {
		m["prop_hyperlink"] = *s.Hyperlink
	}

	return m
}

func NewSubText(t *Text) SubText {
	return SubText{
		Family:    t.Family,
		Style:     t.Style,
		Size:      t.Size,
		Color:     t.Color,
		Hyperlink: t.Hyperlink,
	}
}
