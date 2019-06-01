package maroto

// Represents position properties
type PositionProp struct {
	Left   float64
	Top    float64
	Center float64
}

// Represents rectangle properties
type RectProp struct {
	Left    float64
	Top     float64
	Percent float64
	Center  bool
}

// Represents font properties
type FontProp struct {
	Family Family
	Style  Style
	Size   float64
	Align  Align
}

// Represents signature properties
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

// Make font properties valid
func (f *FontProp) MakeValid() {
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
}

// Make signature properties valid
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
