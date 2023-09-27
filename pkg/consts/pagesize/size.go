package pagesize

type Type string

const (
	A1      Type = "a1"
	A2      Type = "a2"
	A3      Type = "a3"
	A4      Type = "a4"
	A5      Type = "a5"
	A6      Type = "a6"
	Letter  Type = "letter"
	Legal   Type = "legal"
	Tabloid Type = "tabloid"

	MinTopMargin      = 10.0
	MinLeftMargin     = 10.0
	MinRightMargin    = 10.0
	MinBottomMargin   = 20.0025
	DefaultFontSize   = 10.0
	DefaultMaxGridSum = 12.0
)

func GetDimensions(pageSize Type) (float64, float64) {
	switch pageSize {
	case A1:
		return 594.0, 841.0
	case A2:
		return 419.9, 594.0
	case A3:
		return 297.0, 419.9
	case A5:
		return 148.4, 210.0
	case A6:
		return 105.0, 148.5
	case Letter:
		return 215.9, 279.4
	case Legal:
		return 215.9, 355.6
	case Tabloid:
		return 279.4, 431.8
	default: // A4
		return 210.0, 297.0
	}
}
