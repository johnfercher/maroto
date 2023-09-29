package code

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

type code struct{}

// New create a Code.
func New() *code {
	return &code{}
}

// GenDataMatrix is responsible to generate a data matrix byte array.
func (c *code) GenDataMatrix(code string) ([]byte, error) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		return nil, err
	}

	return c.getBytes(dataMatrix)
}

// GenQr is responsible to generate a qr code byte array.
func (c *code) GenQr(code string) ([]byte, error) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return c.getBytes(qrCode)
}

// GenBar is responsible to generate a barcode byte array.
func (c *code) GenBar(code string, cell *core.Cell, prop *props.Barcode) ([]byte, error) {
	barCode, err := code128.Encode(code)
	if err != nil {
		return nil, err
	}

	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width

	proportion := float64(len(code)) * 20.0 / cell.Width

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
