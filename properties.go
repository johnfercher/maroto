package maroto

// Proportion represents a proportion from a rectangle, example: 16x9, 4x3...
type Proportion struct {
	Width  float64
	Height float64
}

// BarcodeProp represents properties from a barcode inside a cell
type BarcodeProp struct {
	Left       float64
	Top        float64
	Percent    float64
	Proportion Proportion
	Center     bool
}

// RectProp represents properties from a rectangle (Image, QrCode or Barcode) inside a cell
type RectProp struct {
	Left    float64
	Top     float64
	Percent float64
	Center  bool
}

// TextProp represents properties from a Text inside a cell
type TextProp struct {
	Top         float64
	Family      Family
	Style       Style
	Size        float64
	Align       Align
	Extrapolate bool
}

// FontProp represents properties from a text
type FontProp struct {
	Family Family
	Style  Style
	Size   float64
}

// TableListProp represents properties from a TableList
type TableListProp struct {
	HeaderHeight       float64
	HeaderProp         FontProp
	ContentProp        FontProp
	ContentHeight      float64
	Align              Align
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
