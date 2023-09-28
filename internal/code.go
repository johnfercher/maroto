package internal

import (
	"bytes"
	image2 "image"
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/codabar"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/merror"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type code struct {
	math  core.Math
	image core.Image
	text  core.Text
}

// NewCode create a Code.
func NewCode(math core.Math, image core.Image, text core.Text) *code {
	return &code{
		math:  math,
		image: image,
		text:  text,
	}
}

// AddDataMatrix creates a DataMatrix code inside a cell.
func (c *code) AddDataMatrix(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		c.text.Add("could not generate data matrix", cell, merror.DefaultErrorText)
		return
	}

	c.addImage(dataMatrix, cell, margins, prop)
}

func (c *code) AddCoda(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	codaBar, err := codabar.Encode(code)
	if err != nil {
		c.text.Add("could not generate data matrix", cell, merror.DefaultErrorText)
		return
	}

	c.addImage(codaBar, cell, margins, prop)
}

// AddQr create a QrCode inside a cell.
func (c *code) AddQr(code string, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		c.text.Add("could not generate qr code", cell, merror.DefaultErrorText)
		return
	}

	c.addImage(qrCode, cell, margins, prop)
}

// AddBar create a Barcode inside a cell.
func (c *code) AddBar(code string, cell *core.Cell, margins *config.Margins, prop *props.Barcode) {
	barCode, err := code128.Encode(code)
	if err != nil {
		c.text.Add("could not generate barcode", cell, merror.DefaultErrorText)
		return
	}

	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width

	proportion := 442.0 / cell.Width

	width := int(proportion * cell.Width)
	height := int(cell.Width * heightPercentFromWidth * proportion)

	scaledBarCode, err := barcode.Scale(barCode, width, height)
	if err != nil {
		c.text.Add("could not scale barcode", cell, merror.DefaultErrorText)
		return
	}

	c.addImage(scaledBarCode, cell, margins, prop.ToRectProp())
}

func (c *code) addImage(img image2.Image, cell *core.Cell, margins *config.Margins, rect *props.Rect) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		c.text.Add("could not encode barcode", cell, merror.DefaultErrorText)
		return
	}

	err = c.image.AddFromBytes(buf.Bytes(), cell, margins, rect, extension.Jpg)
	if err != nil {
		c.text.Add("failed to add barcode to document", cell, merror.DefaultErrorText)
	}
}
