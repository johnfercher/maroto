package maroto

// RectProp represents properties from a rectangle (Image, QrCode or Barcode) inside a cell
type RectProp struct {
	Left    float64
	Top     float64
	Percent float64
	Center  bool
}

// TextProp represents properties from a Text inside a cell
type TextProp struct {
	Top    float64
	Family Family
	Style  Style
	Size   float64
	Align  Align
}

// SignatureProp represents properties from a Signature inside a cell
type SignatureProp struct {
	Family Family
	Style  Style
	Size   float64
}

// TableListProp represents properties from a TableList
type TableListProp struct {
	HFontSize   float64
	HFontFamily Family
	HFontStyle  Style
	Align       Align
	HHeight     float64
	Space       float64
	CFontSize   float64
	CFontFamily Family
	CFontStyle  Style
	CHeight     float64
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

// MakeValid from SignatureProp define default values for a Signature
func (f *SignatureProp) MakeValid() {
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

// MakeValid from TableListProp define default values for a TableList
func (t *TableListProp) MakeValid() {
	if t.HFontSize == 0.0 {
		t.HFontSize = 10.0
	}

	if t.HFontFamily == "" {
		t.HFontFamily = Arial
	}

	if t.HFontStyle == "" {
		t.HFontStyle = Bold
	}

	if t.HHeight == 0.0 {
		t.HHeight = 7.0
	}

	if t.Align == "" {
		t.Align = Left
	}

	if t.CFontSize == 0.0 {
		t.CFontSize = 10.0
	}

	if t.CFontFamily == "" {
		t.CFontFamily = Arial
	}

	if t.CFontStyle == "" {
		t.CFontStyle = Normal
	}

	if t.CHeight == 0.0 {
		t.CHeight = 5.0
	}

	if t.Space == 0.0 {
		t.Space = 4.0
	}
}
