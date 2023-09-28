package core

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetInnerCenterCell(inner *config.Dimensions, outer *config.Dimensions, percent float64) *Cell
	GetInnerNonCenterCell(inner *config.Dimensions, outer *config.Dimensions, prop *props.Rect) *Cell
	GetCenterCorrection(outerSize, innerSize float64) float64
}

// Code is the abstraction which deals of how to add QrCodes or Barcode in a PDF.
type Code interface {
	AddQr(code string, cell *Cell, margins *config.Margins, prop *props.Rect)
	AddBar(code string, cell *Cell, margins *config.Margins, prop *props.Barcode)
	AddDataMatrix(code string, cell *Cell, margins *config.Margins, prop *props.Rect)
}

// Font is the abstraction which deals of how to set fontstyle configurations.
type Font interface {
	SetFamily(family string)
	SetStyle(style fontstyle.Type)
	SetSize(size float64)
	SetFont(family string, style fontstyle.Type, size float64)
	GetFamily() string
	GetStyle() fontstyle.Type
	GetSize() float64
	GetFont() (string, fontstyle.Type, float64)
	GetScaleFactor() (scaleFactor float64)
	SetColor(color *props.Color)
	GetColor() *props.Color
}

// Image is the abstraction which deals of how to add images in a PDF.
type Image interface {
	AddFromBase64(stringBase64 string, cell *Cell, margins *config.Margins, prop *props.Rect, extension extension.Type) (err error)
	AddFromBytes(imgBytes []byte, cell *Cell, margins *config.Margins, prop *props.Rect, extension extension.Type) (err error)
}

type Line interface {
	Add(cell *Cell, prop *props.Line)
}

// Signature is the abstraction which deals of how to add a signature space inside PDF.
type Signature interface {
	AddSpaceFor(label string, cell *Cell, textProp *props.Text)
}

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell *Cell, textProp *props.Text)
	GetLinesQuantity(text string, fontFamily props.Text, colWidth float64) int
}
