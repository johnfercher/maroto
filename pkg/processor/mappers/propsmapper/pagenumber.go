package propsmapper

// PageNumber have attributes of page number
type PageNumber struct {
	// Pattern is the string pattern which will be used to apply the page count component.
	Pattern string
	// Place defines where the page count component will be placed.
	Place string
	// Family defines which font family will be applied to page count.
	Family string
	// Style defines which font style will be applied to page count.
	Style string
	// Size defines which font size will be applied to page count.
	Size float64
	// Color defines which will be applied to page count.
	Color *Color
}
