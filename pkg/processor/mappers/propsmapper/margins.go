package propsmapper

// Margins is the representation of a margin.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// NewMargins is responsible for creating the margins, if the font fields cannot be
// converted, an invalid value is set.
func NewMargins(margins interface{}) *Margins {
	marginsMap, ok := margins.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Margins{
		Left:   *convertFields(marginsMap["left"], -1.0),
		Right:  *convertFields(marginsMap["right"], -1.0),
		Top:    *convertFields(marginsMap["top"], -1.0),
		Bottom: *convertFields(marginsMap["bottom"], -1.0),
	}
}
