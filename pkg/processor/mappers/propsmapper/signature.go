package propsmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

// Signature represents properties from a signature.
type Signature struct {
	// FontFamily of the text, ex: consts.Arial, helvetica and etc.
	FontFamily string
	// FontStyle of the text, ex: consts.Normal, bold and etc.
	FontStyle string
	// FontSize of the text.
	FontSize float64
	// FontColor define the font color.
	FontColor *Color
	// LineColor define the line color.
	LineColor *Color
	// LineStyle define the line style (solid or dashed).
	LineStyle linestyle.Type
	// LineThickness define the line thickness.
	LineThickness float64

	SafePadding float64
}

// NewSignature is responsible for creating the Signature, if the font fields cannot be
// converted, an invalid value is set.
func NewSignature(signature interface{}) (*Signature, error) {
	signatureMap, ok := signature.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure barcode props can be converted to map[string] interface{}")
	}

	return &Signature{
		FontFamily:    *convertFields(signatureMap["font_family"], ""),
		FontStyle:     NewFontStyle(*convertFields(signatureMap["font_style"], "")),
		FontSize:      *convertFields(signatureMap["font_size"], 0.0),
		FontColor:     NewColor(signatureMap["font_color"]),
		LineColor:     NewColor(signatureMap["line_color"]),
		LineStyle:     linestyle.Type(NewLineStyle(*convertFields(signatureMap["line_style"], ""))),
		LineThickness: *convertFields(signatureMap["line_thickness"], 0.0),
		SafePadding:   *convertFields(signatureMap["safe_padding"], -1.0),
	}, nil
}
