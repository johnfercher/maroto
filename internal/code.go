package internal

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

// Code is the abstraction which deals of how to add QrCodes or Barcode in a PDF.
type Code interface {
	AddQr(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect)
	AddBar(code string, cell *core.Cell, margins *config.Margins, prop *props.Barcode) (err error)
	AddDataMatrix(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect)
}

type code struct {
	pdf  fpdf.Fpdf
	math math.Math
}

// NewCode create a Code.
func NewCode(pdf fpdf.Fpdf, math math.Math) *code {
	return &code{
		pdf,
		math,
	}
}

// AddDataMatrix creates a DataMatrix code inside a cell.
func (s *code) AddDataMatrix(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	key := barcode.RegisterDataMatrix(s.pdf, code)
	dimensions := &config.Dimensions{Width: cell.Width, Height: cell.Width}

	var rectCell *core.Cell
	if prop.Center {
		rectCell = s.math.GetRectCenterColProperties(dimensions, cell, margins, prop.Percent)
	} else {
		rectCell = s.math.GetRectNonCenterColProperties(dimensions, cell, margins, prop)
	}

	barcode.Barcode(s.pdf, key, rectCell.X, rectCell.Y+cell.Y, rectCell.Width, rectCell.Height, false)
}

// AddQr create a QrCode inside a cell.
func (s *code) AddQr(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	key := barcode.RegisterQR(s.pdf, code, qr.H, qr.Unicode)
	dimensions := &config.Dimensions{Width: cell.Width, Height: cell.Width}

	var rectCell *core.Cell
	if prop.Center {
		rectCell = s.math.GetRectCenterColProperties(dimensions, cell, margins, prop.Percent)
	} else {
		rectCell = s.math.GetRectNonCenterColProperties(dimensions, cell, margins, prop)
	}

	barcode.Barcode(s.pdf, key, rectCell.X, rectCell.Y+cell.Y, rectCell.Width, rectCell.Height, false)
}

// AddBar create a Barcode inside a cell.
func (s *code) AddBar(code string, cell *core.Cell, margins *config.Margins, prop *props.Barcode) (err error) {
	bcode, err := code128.Encode(code)
	if err != nil {
		return
	}

	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width
	dimensions := &config.Dimensions{Width: cell.Width, Height: cell.Width * heightPercentFromWidth}

	var rectCell *core.Cell
	if prop.Center {
		rectCell = s.math.GetRectCenterColProperties(dimensions, cell, margins, prop.Percent)
	} else {
		rectProps := &props.Rect{Left: prop.Left, Top: prop.Top, Center: prop.Center, Percent: prop.Percent}
		rectCell = s.math.GetRectNonCenterColProperties(dimensions, cell, margins, rectProps)
	}

	barcode.Barcode(s.pdf, barcode.Register(bcode), rectCell.X, rectCell.Y+cell.Y, rectCell.Width, rectCell.Height, false)
	return
}
