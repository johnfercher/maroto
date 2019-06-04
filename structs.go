package maroto

// Represents rectangle properties
type RectProp struct {
	Left    float64
	Top     float64
	Percent float64
	Center  bool
}

// Represents TextHelper properties
type TextProp struct {
	Top    float64
	Family Family
	Style  Style
	Size   float64
	Align  Align
}

// Represents Signature properties
type SignatureProp struct {
	Family Family
	Style  Style
	Size   float64
}

// Make rectangle properties valid
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

// Make Text properties valid
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

// Make Signature properties valid
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
