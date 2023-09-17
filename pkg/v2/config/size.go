package config

type PageSize string

const (
	A1      PageSize = "a1"
	A2      PageSize = "a2"
	A3      PageSize = "a3"
	A4      PageSize = "a4"
	A5      PageSize = "a5"
	A6      PageSize = "a6"
	Letter  PageSize = "letter"
	Legal   PageSize = "legal"
	Tabloid PageSize = "tabloid"

	MinTopMargin    = 10.0
	MinLeftMargin   = 10.0
	MinRightMargin  = 10.0
	MinBottomMargin = 20.0
)

func GetDimensions(pageSize PageSize) *Dimensions {
	switch pageSize {
	case A1:
		return &Dimensions{594.0, 841.0}
	case A2:
		return &Dimensions{419.9, 594.0}
	case A3:
		return &Dimensions{297.0, 419.9}
	case A5:
		return &Dimensions{148.4, 210.0}
	case A6:
		return &Dimensions{105.0, 148.5}
	case Letter:
		return &Dimensions{215.9, 279.4}
	case Legal:
		return &Dimensions{215.9, 355.6}
	case Tabloid:
		return &Dimensions{279.4, 431.8}
	default: // A4
		return &Dimensions{210.0, 297.0}
	}
}
