package maroto

type PositionProp struct {
	Left   float64
	Top    float64
	Center float64
}

type RectProp struct {
	Left    float64
	Top     float64
	Percent float64
	Center  bool
}

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
