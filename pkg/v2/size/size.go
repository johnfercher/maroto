package size

type PageSize string

const (
	A4 PageSize = "a4"

	MinTopMargin    = 10.0
	MinLeftMargin   = 10.0
	MinRightMargin  = 10.0
	MinBottomMargin = 20.0
)

func GetDimensions(pageSize PageSize) (width float64, height float64) {
	switch pageSize {
	default:
		return 210, 297
	}
}
