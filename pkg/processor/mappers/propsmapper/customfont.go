package propsmapper

type CustomFont struct {
	Family *string
	Style  *string
	File   *string
	Bytes  []*byte
}

// NewCustomFont is responsible for creating the CustomFont, if the CustomFont fields cannot be
// converted, an invalid value is set.
func NewCustomFont(font interface{}) *CustomFont {
	fontMap, ok := font.(map[string]interface{})
	if !ok {
		return nil
	}

	return &CustomFont{
		Family: convertFields(fontMap["family"], ""),
		Style:  convertFields(fontMap["style"], ""),
		File:   convertFields(fontMap["file"], ""),
		// Bytes:  NewColor(fontMap["bytes"]),
	}
}
