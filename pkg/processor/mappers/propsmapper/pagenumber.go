package propsmapper

// PageNumber have attributes of page number
type PageNumber struct {
	// Pattern is the string pattern which will be used to apply the page count component.
	Pattern *string
	// Place defines where the page count component will be placed.
	Place *string
	// Family defines which font family will be applied to page count.
	Family *string
	// Style defines which font style will be applied to page count.
	Style *string
	// Size defines which font size will be applied to page count.
	Size *float64
	// Color defines which will be applied to page count.
	Color *Color
}

// NewPageNumber is responsible for creating the pageNumber, if the pageNumber fields cannot be
// converted, an invalid value is set.
func NewPageNumber(pageNumber interface{}) *PageNumber {
	pageNumberMap, ok := pageNumber.(map[string]interface{})
	if !ok {
		return nil
	}

	return &PageNumber{
		Pattern: convertFields(pageNumberMap["pattern"], ""),
		Place:   convertFields(pageNumberMap["place"], ""),
		Family:  convertFields(pageNumberMap["family"], ""),
		Style:   convertFields(pageNumberMap["style"], ""),
		Size:    convertFields(pageNumberMap["size"], 0.0),
		Color:   NewColor(pageNumberMap["color"]),
	}
}
