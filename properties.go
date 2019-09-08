package maroto

// Proportion represents a proportion from a rectangle, example: 16x9, 4x3...
type Proportion struct {
	// Width from the rectangle: Barcode, image and etc
	Width float64
	// Height from the rectangle: Barcode, image and etc
	Height float64
}

// BarcodeProp represents properties from a barcode inside a cell
type BarcodeProp struct {
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

// RectProp represents properties from a rectangle (Image, QrCode or Barcode) inside a cell
type RectProp struct {
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

// TextProp represents properties from a Text inside a cell
type TextProp struct {
	// Top is space between the upper cell limit to the barcode, if align is not center
	Top float64
	// Family of the text, ex: Arial, helvetica and etc
	Family Family
	// Style of the text, ex: Normal, bold and etc
	Style Style
	// Size of the text
	Size float64
	// Align of the text
	Align Align
	// Extrapolate define if the text will automatically add a new line when
	// text reach the right cell boundary
	Extrapolate bool
}

// FontProp represents properties from a text
type FontProp struct {
	// Family of the text, ex: Arial, helvetica and etc
	Family Family
	// Style of the text, ex: Normal, bold and etc
	Style Style
	// Size of the text
	Size float64
}

// TableListProp represents properties from a TableList
type TableListProp struct {
	// HeaderHeight is the height of the cell with headers
	HeaderHeight float64
	// HeaderProp is the custom properties of the text inside
	// the headers
	HeaderProp FontProp
	// ContentHeight is the height of the cells with contents
	ContentHeight float64
	// ContentProp is the custom properties of the text inside
	// the contents
	ContentProp FontProp
	// Align is the align of the text (header and content) inside the columns
	Align Align
	// HeaderContentSpace is the space between the header and the contents
	HeaderContentSpace float64
}

// MakeValid from RectProp means will make the properties from a rectangle reliable to fit inside a cell
// and define default values for a rectangle
func (r *RectProp) MakeValid() {
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

// MakeValid from BarcodeProp means will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode
func (r *BarcodeProp) MakeValid() {
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

// MakeValid from TextProp define default values for a Text
func (f *TextProp) MakeValid() {
	if f.Family == "" {
		f.Family = Arial
	}

	if f.Style == "" {
		f.Style = Normal
	}

	if f.Align == "" {
		f.Align = Left
	}

	if f.Size == 0.0 {
		f.Size = 10.0
	}

	if f.Top < 0.0 {
		f.Top = 0.0
	}
}

// MakeValid from FontProp define default values for a Signature
func (f *FontProp) MakeValid() {
	if f.Family == "" {
		f.Family = Arial
	}

	if f.Style == "" {
		f.Style = Bold
	}

	if f.Size == 0.0 {
		f.Size = 8.0
	}
}

// ToTextProp from FontProp return a TextProp based on FontProp
func (f *FontProp) ToTextProp(align Align, top float64) TextProp {
	textProp := TextProp{
		Family: f.Family,
		Style:  f.Style,
		Size:   f.Size,
		Align:  align,
		Top:    top,
	}

	textProp.MakeValid()

	return textProp
}

// MakeValid from TableListProp define default values for a TableList
func (t *TableListProp) MakeValid() {
	if t.HeaderProp.Size == 0.0 {
		t.HeaderProp.Size = 10.0
	}

	if t.HeaderProp.Family == "" {
		t.HeaderProp.Family = Arial
	}

	if t.HeaderProp.Style == "" {
		t.HeaderProp.Style = Bold
	}

	if t.HeaderHeight == 0.0 {
		t.HeaderHeight = 7.0
	}

	if t.Align == "" {
		t.Align = Left
	}

	if t.ContentProp.Size == 0.0 {
		t.ContentProp.Size = 10.0
	}

	if t.ContentProp.Family == "" {
		t.ContentProp.Family = Arial
	}

	if t.ContentProp.Style == "" {
		t.ContentProp.Style = Normal
	}

	if t.ContentHeight == 0.0 {
		t.ContentHeight = 5.0
	}

	if t.HeaderContentSpace == 0.0 {
		t.HeaderContentSpace = 4.0
	}
}
