package props

type HeatMap struct {
	TransparentValues []int
	HalfColor         bool
	InvertScale       bool
	Chart             *Chart
}

func (h *HeatMap) MakeValid() {

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
