package config

type PageSize string

const (
	A4 PageSize = "a4"

	MinTopMargin    = 10.0
	MinLeftMargin   = 10.0
	MinRightMargin  = 10.0
	MinBottomMargin = 20.0
)

func GetDimensions(pageSize PageSize) *Dimensions {
	switch pageSize {
	default:
		return &Dimensions{215.9, 279.4}
	}
}
