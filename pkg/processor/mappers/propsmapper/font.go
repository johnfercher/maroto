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

// NewFont is responsible for creating the fonts, if the font fields cannot be
// converted, an invalid value is set.
func NewFont(font interface{}) *Font {
	fontMap, ok := font.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Font{
		Family: *convertFields(fontMap["family"], ""),
		Style:  *convertFields(fontMap["style"], ""),
		Size:   *convertFields(fontMap["size"], 0.0),
		Color:  NewColor(fontMap["color"]),
	}
}

func NewListFont(fontsSource interface{}) []*Font {
	listOfFonts, ok := fontsSource.([]interface{})
	if !ok {
		return nil
	}
	fonts := make([]*Font, len(listOfFonts))

	for i, font := range listOfFonts {
		fonts[i] = NewFont(font)
	}

	return fonts
}
