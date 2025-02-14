package props

type HeatMap struct {
	TransparentValues []int
	WidthPercent      float64
	HeightPercent     float64
	Chart             *Chart
	Center            bool
}

func (h *HeatMap) MakeValid() {
	if h.WidthPercent == 0 {
		h.WidthPercent = 1
	}

	if h.HeightPercent == 0 {
		h.HeightPercent = 1
	}
}

func (h *HeatMap) ToMap() map[string]any {
	if h == nil {
		return nil
	}

	m := make(map[string]interface{})

	if len(h.TransparentValues) > 0 {
		m["prop_transparent_values"] = h.TransparentValues
	}

	return m
}
