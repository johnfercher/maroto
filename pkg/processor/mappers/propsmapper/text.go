package propsmapper

import "fmt"

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Left is the minimal amount of space between the left cell boundary and the text.
	Left float64
	// Right is the minimal amount of space between the right cell boundary and the text.
	Right float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style string
	// Size of the text.
	Size float64
	// Align of the text.
	Align string
	// BreakLineStrategy define the break line strategy.
	BreakLineStrategy string
	// VerticalPadding define an additional space between linet.
	VerticalPadding float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink string
}

// NewText is responsible for creating the Text, if the font fields cannot be
// converted, an invalid value is set.
func NewText(signature interface{}) (*Text, error) {
	signatureMap, ok := signature.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure text props can be converted to map[string] interface{}")
	}

	return &Text{
		Top:               *convertFields(signatureMap["top"], -1.0),
		Left:              *convertFields(signatureMap["left"], -1.0),
		Right:             *convertFields(signatureMap["right"], -1.0),
		Family:            *convertFields(signatureMap["family"], ""),
		Style:             NewFontStyle(*convertFields(signatureMap["style"], "")),
		Size:              *convertFields(signatureMap["size"], 0.0),
		Align:             NewAlign(*convertFields(signatureMap["align"], "")),
		BreakLineStrategy: NewBreakLineStrategy(*convertFields(signatureMap["break_line_strategy"], "")),
		VerticalPadding:   *convertFields(signatureMap["vertical_padding"], -1.0),
		Color:             NewColor(signatureMap["color"]),
		Hyperlink:         *convertFields(signatureMap["hyperlink"], ""),
	}, nil
}
