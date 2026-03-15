package code

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/png"

	libBarcode "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"

	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var (
	ErrCannotEncodePNG        = errors.New("cannot encode png")
	ErrCannotScaleBarcode     = errors.New("cannot scale barcode")
	ErrCannotEncodeQRcode     = errors.New("cannot encode qr code")
	ErrCannotEncodeDataMatrix = errors.New("cannot encode data matrix")
)

type Code struct{}

// New create a Code (Singleton).
func New() *Code {
	return &Code{}
}

// GenDataMatrix is responsible to generate a data matrix byte array.
func (c *Code) GenDataMatrix(code string) (*entity.Image, error) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotEncodeDataMatrix, err)
	}

	return c.getImage(dataMatrix)
}

// GenQr is responsible to generate a qr code byte array.
func (c *Code) GenQr(code string) (*entity.Image, error) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotEncodeQRcode, err)
	}

	return c.getImage(qrCode)
}

// GenBar is responsible to generate a barcode byte array.
func (c *Code) GenBar(code string, _ *entity.Cell, prop *props.Barcode) (*entity.Image, error) {
	barcodeGen := getBarcodeClosure(prop.Type)

	barCode, err := barcodeGen(code)
	if err != nil {
		return nil, err
	}

	width := float64(barCode.Bounds().Dx())
	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width
	height := int(width * heightPercentFromWidth)

	scaledBarCode, err := libBarcode.Scale(barCode, int(width), height)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotScaleBarcode, err)
	}

	return c.getImage(scaledBarCode)
}

func getBarcodeClosure(
	barcodeType barcode.Type,
) func(code string) (libBarcode.BarcodeIntCS, error) {
	switch barcodeType {
	case barcode.EAN:
		return ean.Encode
	case barcode.Code128:
		return code128.Encode
	default:

		return code128.Encode
	}
}

func (c *Code) getImage(img image.Image) (*entity.Image, error) {
	var buf bytes.Buffer

	dst := image.NewPaletted(img.Bounds(), palette.Plan9)
	drawer := draw.Drawer(draw.Src)
	drawer.Draw(dst, dst.Bounds(), img, img.Bounds().Min)

	err := png.Encode(&buf, dst)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotEncodePNG, err)
	}

	imgEntity := &entity.Image{
		Bytes:     buf.Bytes(),
		Extension: extension.Png,
		Dimensions: &entity.Dimensions{
			Width:  float64(dst.Bounds().Dx()),
			Height: float64(dst.Bounds().Dy()),
		},
	}

	return imgEntity, nil
}
