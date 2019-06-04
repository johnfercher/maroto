package maroto

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

type Code interface {
	AddQr(code string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64)
	AddBar(code string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64) (err error)
}

type code struct {
	pdf  gofpdf.Pdf
	math Math
}

func NewCode(pdf gofpdf.Pdf, math Math) Code {
	return &code{
		pdf,
		math,
	}
}

func (c *code) AddQr(code string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64) {
	key := barcode.RegisterQR(c.pdf, code, qr.H, qr.Unicode)

	actualWidthPerCol := c.math.GetWidthPerCol(qtdCols)

	x, y, w, h := c.math.GetRectCenterColProperties(actualWidthPerCol, actualWidthPerCol, qtdCols, colHeight, indexCol, percent)

	barcode.Barcode(c.pdf, key, x, y+marginTop, w, h, false)
}

func (c *code) AddBar(code string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64) (err error) {
	bcode, err := code128.Encode(code)

	if err != nil {
		return
	}

	actualWidthPerCol := c.math.GetWidthPerCol(qtdCols)

	x, y, w, h := c.math.GetRectCenterColProperties(actualWidthPerCol, actualWidthPerCol*0.33, qtdCols, colHeight, indexCol, percent)

	barcode.Barcode(c.pdf, barcode.Register(bcode), x, y+marginTop, w, h, false)
	return
}
