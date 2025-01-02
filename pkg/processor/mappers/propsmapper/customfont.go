package propsmapper

type CustomFont struct {
	Family string
	Style  string
	File   string
}

// newCustomFont is responsible for creating the CustomFont, if the CustomFont fields cannot be
// converted, an invalid value is set.
func newCustomFont(font interface{}) *CustomFont {
	fontMap, ok := font.(map[string]interface{})
	if !ok {
		return nil
	}

	return &CustomFont{
		Family: *convertFields(fontMap["family"], ""),
		Style:  NewFontStyle(*convertFields(fontMap["style"], "")),
		File:   *convertFields(fontMap["file_path"], ""),
	}
}

// NewCustomFont is responsible for creating a CustomFont list
func NewCustomFonts(interfaceFonts interface{}) []*CustomFont {
	fonts, ok := interfaceFonts.([]interface{})
	if !ok {
		return nil
	}
	customFonts := make([]*CustomFont, 0, len(fonts))

	for _, font := range fonts {
		if newFont := newCustomFont(font); newFont != nil {
			customFonts = append(customFonts, newCustomFont(font))
		}
	}

	return customFonts
}
