package propsmapper

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: constf.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: constf.Normal, bold and etc.
	Style string
	// Size of the text.
	Size float64
	// Color define the font color.
	Color *Color
}
