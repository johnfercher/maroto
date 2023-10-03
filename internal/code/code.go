package code

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// codeInstance is the singleton of code, opted to use a singleton to ensure that
// this will not be instantiated more than once since there is no need to do this
// because code is stateless.
var codeInstance *code = nil

type code struct{}

// New create a Code (Singleton).
func New() *code {
	if codeInstance == nil {
		codeInstance = &code{}
	}
	return codeInstance
}

// GenDataMatrix is responsible to generate a data matrix byte array.
func (c *code) GenDataMatrix(code string) (*entity.Image, error) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		return nil, err
	}

	return c.getImage(dataMatrix)
}

// GenQr is responsible to generate a qr code byte array.
func (c *code) GenQr(code string) (*entity.Image, error) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return c.getImage(qrCode)
}

// GenBar is responsible to generate a barcode byte array.
func (c *code) GenBar(code string, cell *entity.Cell, prop *props.Barcode) (*entity.Image, error) {
	barCode, err := code128.Encode(code)
	if err != nil {
		return nil, err
	}

	width := float64(barCode.Bounds().Dx())
	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width
	height := int(width * heightPercentFromWidth)

	scaledBarCode, err := barcode.Scale(barCode, int(width), height)
	if err != nil {
		return nil, err
	}

	return c.getImage(scaledBarCode)
}

func (c *code) getImage(img image.Image) (*entity.Image, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	imgEntity := &entity.Image{
		Bytes:     buf.Bytes(),
		Extension: extension.Jpg,
		Dimensions: &entity.Dimensions{
			Width:  float64(img.Bounds().Dx()),
			Height: float64(img.Bounds().Dy()),
		},
	}

	return imgEntity, nil
}
