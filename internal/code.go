package internal

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

// Code is the abstraction which deals of how to add QrCodes or Barcode in a PDF
type Code interface {
	AddQr(code string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect)
	AddBar(code string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Barcode) (err error)
}

type code struct {
	pdf  gofpdf.Pdf
	math Math
}

// NewCode create a Code
func NewCode(pdf gofpdf.Pdf, math Math) *code {
	return &code{
		pdf,
		math,
	}
}

// AddQr create a QrCode inside a cell
func (s *code) AddQr(code string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect) {
	key := barcode.RegisterQR(s.pdf, code, qr.H, qr.Unicode)

	var x, y, w, h float64
	if prop.Center {
		x, y, w, h = s.math.GetRectCenterColProperties(colWidth, colWidth, colWidth, colHeight, xColOffset, prop.Percent)
	} else {
		x, y, w, h = s.math.GetRectNonCenterColProperties(colWidth, colWidth, colWidth, colHeight, xColOffset, prop)
	}

	barcode.Barcode(s.pdf, key, x, y+yColOffset, w, h, false)
}

// AddBar create a Barcode inside a cell
func (s *code) AddBar(code string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Barcode) (err error) {
	bcode, err := code128.Encode(code)

	if err != nil {
		return
	}

	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width
	var x, y, w, h float64
	if prop.Center {
		x, y, w, h = s.math.GetRectCenterColProperties(colWidth, colWidth*heightPercentFromWidth, colWidth, colHeight, xColOffset, prop.Percent)
	} else {
		rectProps := props.Rect{Left: prop.Left, Top: prop.Top, Center: prop.Center, Percent: prop.Percent}
		x, y, w, h = s.math.GetRectNonCenterColProperties(colWidth, colWidth*heightPercentFromWidth, colWidth, colHeight, xColOffset, rectProps)
	}

	barcode.Barcode(s.pdf, barcode.Register(bcode), x, y+yColOffset, w, h, false)
	return
}
