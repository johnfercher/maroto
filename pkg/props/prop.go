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
	// HeaderHeight is the height of the cell with headers
	HeaderHeight float64
	// HeaderProp is the custom properties of the text inside
	// the headers
	HeaderProp Font
	// ContentHeight is the height of the cells with contents
	ContentHeight float64
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
func (r *Rect) MakeValid() {
	if r.Percent <= 0.0 || r.Percent > 100.0 {
		r.Percent = 100.0
	}

	if r.Center {
		r.Left = 0
		r.Top = 0
	}

	if r.Left < 0.0 {
		r.Left = 0.0
	}

	if r.Top < 0.0 {
		r.Top = 0
	}
}

// MakeValid from Barcode means will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode
func (r *Barcode) MakeValid() {
	if r.Percent <= 0.0 || r.Percent > 100.0 {
		r.Percent = 100.0
	}

	if r.Center {
		r.Left = 0
		r.Top = 0
	}

	if r.Left < 0.0 {
		r.Left = 0.0
	}

	if r.Top < 0.0 {
		r.Top = 0
	}

	if r.Proportion.Width <= 0 {
		r.Proportion.Width = 1
	}

	if r.Proportion.Height <= 0 {
		r.Proportion.Height = 1
	}

	if r.Proportion.Height > r.Proportion.Width*0.33 {
		r.Proportion.Height = r.Proportion.Width * 0.33
	}
}

// MakeValid from Text define default values for a Text
func (f *Text) MakeValid() {
	if f.Family == "" {
		f.Family = consts.Arial
	}

	if f.Style == "" {
		f.Style = consts.Normal
	}

	if f.Align == "" {
		f.Align = consts.Left
	}

	if f.Size == 0.0 {
		f.Size = 10.0
	}

	if f.Top < 0.0 {
		f.Top = 0.0
	}
}

// MakeValid from Font define default values for a Signature
func (f *Font) MakeValid() {
	if f.Family == "" {
		f.Family = consts.Arial
	}

	if f.Style == "" {
		f.Style = consts.Bold
	}

	if f.Size == 0.0 {
		f.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font
func (f *Font) ToTextProp(align consts.Align, top float64) Text {
	textProp := Text{
		Family: f.Family,
		Style:  f.Style,
		Size:   f.Size,
		Align:  align,
		Top:    top,
	}

	textProp.MakeValid()

	return textProp
}

// MakeValid from TableList define default values for a TableList
func (t *TableList) MakeValid() {
	if t.HeaderProp.Size == 0.0 {
		t.HeaderProp.Size = 10.0
	}

	if t.HeaderProp.Family == "" {
		t.HeaderProp.Family = consts.Arial
	}

	if t.HeaderProp.Style == "" {
		t.HeaderProp.Style = consts.Bold
	}

	if t.HeaderHeight == 0.0 {
		t.HeaderHeight = 7.0
	}

	if t.Align == "" {
		t.Align = consts.Left
	}

	if t.ContentProp.Size == 0.0 {
		t.ContentProp.Size = 10.0
	}

	if t.ContentProp.Family == "" {
		t.ContentProp.Family = consts.Arial
	}

	if t.ContentProp.Style == "" {
		t.ContentProp.Style = consts.Normal
	}

	if t.ContentHeight == 0.0 {
		t.ContentHeight = 5.0
	}

	if t.HeaderContentSpace == 0.0 {
		t.HeaderContentSpace = 4.0
	}
}
