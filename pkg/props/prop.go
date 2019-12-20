package props

import "github.com/johnfercher/maroto/pkg/consts"

// Proportion represents a proportion from a rectangle, example: 16x9, 4x3...
type Proportion struct {
	// Width from the rectangle: Barcode, image and etc
	Width float64
	// Height from the rectangle: Barcode, image and etc
	Height float64
}

// Barcode represents properties from a barcode inside a cell
type Barcode struct {
	// Left is the space between the left cell boundary to the barcode, if center is false
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false
	Top float64
	// Percent is how much the barcode will occupy the cell,
	// ex 100%: The barcode will fulfill the entire cell
	// ex 50%: The greater side from the barcode will have half the size of the cell
	Percent float64
	// Proportion is the proportion between size of the barcode
	// Ex: 16x9, 4x3...
	Proportion Proportion
	// Center define that the barcode will be vertically and horizontally centralized
	Center bool
}

// Rect represents properties from a rectangle (Image, QrCode or Barcode) inside a cell
type Rect struct {
	// Left is the space between the left cell boundary to the rectangle, if center is false
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false
	Top float64
	// Percent is how much the rectangle will occupy the cell,
	// ex 100%: The rectangle will fulfill the entire cell
	// ex 50%: The greater side from the rectangle will have half the size of the cell
	Percent float64
	// Center define that the barcode will be vertically and horizontally centralized
	Center bool
}

// Text represents properties from a Text inside a cell
type Text struct {
	// Top is space between the upper cell limit to the barcode, if align is not center
	Top float64
	// Family of the text, ex: consts.Arial, helvetica and etc
	Family consts.Family
	// Style of the text, ex: consts.Normal, bold and etc
	Style consts.Style
	// Size of the text
	Size float64
	// Align of the text
	Align consts.Align
	// Extrapolate define if the text will automatically add a new line when
	// text reach the right cell boundary
	Extrapolate bool
	// VerticalPadding define an additional space between lines
	VerticalPadding float64
}

// Font represents properties from a text
type Font struct {
	// Family of the text, ex: consts.Arial, helvetica and etc
	Family consts.Family
	// Style of the text, ex: consts.Normal, bold and etc
	Style consts.Style
	// Size of the text
	Size float64
}

// TableList represents properties from a TableList
type TableList struct {
	// HeaderProp is the custom properties of the text inside
	// the headers
	HeaderProp Font
	// ContentProp is the custom properties of the text inside
	// the contents
	ContentProp Font
	// Align is the align of the text (header and content) inside the columns
	Align consts.Align
	// HeaderContentSpace is the space between the header and the contents
	HeaderContentSpace float64
}

// MakeValid from Rect means will make the properties from a rectangle reliable to fit inside a cell
// and define default values for a rectangle
func (s *Rect) MakeValid() {
	if s.Percent <= 0.0 || s.Percent > 100.0 {
		s.Percent = 100.0
	}

	if s.Center {
		s.Left = 0
		s.Top = 0
	}

	if s.Left < 0.0 {
		s.Left = 0.0
	}

	if s.Top < 0.0 {
		s.Top = 0
	}
}

// MakeValid from Barcode means will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode
func (s *Barcode) MakeValid() {
	if s.Percent <= 0.0 || s.Percent > 100.0 {
		s.Percent = 100.0
	}

	if s.Center {
		s.Left = 0
		s.Top = 0
	}

	if s.Left < 0.0 {
		s.Left = 0.0
	}

	if s.Top < 0.0 {
		s.Top = 0
	}

	if s.Proportion.Width <= 0 {
		s.Proportion.Width = 1
	}

	if s.Proportion.Height <= 0 {
		s.Proportion.Height = 1
	}

	if s.Proportion.Height > s.Proportion.Width*0.20 {
		s.Proportion.Height = s.Proportion.Width * 0.20
	} else if s.Proportion.Height < s.Proportion.Width*0.10 {
		s.Proportion.Height = s.Proportion.Width * 0.10
	}
}

// MakeValid from Text define default values for a Text
func (s *Text) MakeValid() {
	if s.Family == "" {
		s.Family = consts.Arial
	}

	if s.Style == "" {
		s.Style = consts.Normal
	}

	if s.Align == "" {
		s.Align = consts.Left
	}

	if s.Size == 0.0 {
		s.Size = 10.0
	}

	if s.Top < 0.0 {
		s.Top = 0.0
	}

	if s.VerticalPadding < 0 {
		s.VerticalPadding = 0
	}
}

// MakeValid from Font define default values for a Signature
func (s *Font) MakeValid() {
	if s.Family == "" {
		s.Family = consts.Arial
	}

	if s.Style == "" {
		s.Style = consts.Bold
	}

	if s.Size == 0.0 {
		s.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font
func (s *Font) ToTextProp(align consts.Align, top float64, extrapolate bool, verticalPadding float64) Text {
	textProp := Text{
		Family:          s.Family,
		Style:           s.Style,
		Size:            s.Size,
		Align:           align,
		Top:             top,
		Extrapolate:     extrapolate,
		VerticalPadding: verticalPadding,
	}

	textProp.MakeValid()

	return textProp
}

// MakeValid from TableList define default values for a TableList
func (s *TableList) MakeValid() {
	if s.HeaderProp.Size == 0.0 {
		s.HeaderProp.Size = 10.0
	}

	if s.HeaderProp.Family == "" {
		s.HeaderProp.Family = consts.Arial
	}

	if s.HeaderProp.Style == "" {
		s.HeaderProp.Style = consts.Bold
	}

	if s.Align == "" {
		s.Align = consts.Left
	}

	if s.ContentProp.Size == 0.0 {
		s.ContentProp.Size = 10.0
	}

	if s.ContentProp.Family == "" {
		s.ContentProp.Family = consts.Arial
	}

	if s.ContentProp.Style == "" {
		s.ContentProp.Style = consts.Normal
	}

	if s.HeaderContentSpace == 0.0 {
		s.HeaderContentSpace = 4.0
	}
}
