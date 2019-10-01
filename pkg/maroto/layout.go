package maroto

//Margins define a margins of a page layout
type Margins struct {
	Left  float64
	Right float64
	Top   float64
}

//Font define a style of a layout Font
type Font struct {
	Family string
	Style  string
	Size   int
}

type PageSize string

const (
	// A3 represents DIN/ISO A3 page size
	A3 PageSize = "A3"
	// A4 represents DIN/ISO A4 page size
	A4 PageSize = "A4"
	// A5 represents DIN/ISO A5 page size
	A5 PageSize = "A5"
	// Letter represents US Letter page size
	Letter PageSize = "Letter"
	// Legal represents US Legal page size
	Legal PageSize = "Legal"
)

type Layout struct {
	Margins     Margins
	Font        Font
	PageSize    PageSize
	Orientation string
}
