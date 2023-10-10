package entity

type Dimensions struct {
	Width  float64
	Height float64
}

func (d *Dimensions) AppendMap(m map[string]interface{}) map[string]interface{} {
	if d.Width != 0 {
		m["dimension_width"] = d.Width
	}

	if d.Height != 0 {
		m["dimension_height"] = d.Height
	}

	return m
}
