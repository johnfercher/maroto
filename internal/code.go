package internal

import (
	"bytes"
	image2 "image"
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type code struct {
}

// NewCode create a Code.
func NewCode() *code {
	return &code{}
}

func (c *code) GenDataMatrix(code string) ([]byte, error) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		return nil, err
	}

	return c.getBytes(dataMatrix)
}

func (c *code) GenQr(code string) ([]byte, error) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return c.getBytes(qrCode)
}

func (c *code) GenBar(code string, cell *core.Cell, prop *props.Barcode) ([]byte, error) {
	barCode, err := code128.Encode(code)
	if err != nil {
		return nil, err
	}

	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width

	proportion := 442.0 / cell.Width

	width := int(proportion * cell.Width)
	height := int(cell.Width * heightPercentFromWidth * proportion)

	scaledBarCode, err := barcode.Scale(barCode, width, height)
	if err != nil {
		return nil, err
	}

	return c.getBytes(scaledBarCode)
}

func (c *code) getBytes(img image2.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
