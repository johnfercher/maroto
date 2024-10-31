package processorprovider_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateBarCode(t *testing.T) {
	t.Run("when CreateBarCode is called, should generate a barcode", func(t *testing.T) {
		m := processorprovider.NewMaroto()
		barcode := m.CreateBarCode("code",
			&propsmapper.Barcode{
				Left: 10.0, Top: 10.0, Percent: 100.0,
				Proportion: propsmapper.Proportion{Width: 10, Height: 2},
				Center:     false, Type: propsmapper.NewCodeType("code128"),
			},
		)

		assert.IsType(t, barcode, &code.Barcode{})
		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/barcode.json")
	})
}

func TestCreateMatrixCode(t *testing.T) {
	t.Run("when CreateMatrixCode is called, should generate a matrixcode", func(t *testing.T) {
		m := processorprovider.NewMaroto()
		barcode := m.CreateMatrixCode("code",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/matrixcode.json")
	})
}

func TestCreateQRCode(t *testing.T) {
	t.Run("when CreateQrCode is called, should generate a qrcode", func(t *testing.T) {
		m := processorprovider.NewMaroto()
		barcode := m.CreateQrCode("code",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/qrcode.json")
	})
}

func TestCreateImage(t *testing.T) {
	t.Run("when CreateImage is called, should generate a image", func(t *testing.T) {
		m := processorprovider.NewMaroto()
		image := m.CreateImage(make([]byte, 0), "png",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(image.GetStructure()).Equals("processor/provider/image.json")
	})
}

func TestCreateLine(t *testing.T) {
	t.Run("when CreateLine is called, should generate a line", func(t *testing.T) {
		m := processorprovider.NewMaroto()
		barcode := m.CreateLine(
			&propsmapper.Line{
				Color: &propsmapper.Color{Red: 10, Green: 10, Blue: 10}, Style: "solid", Thickness: 10.0,
				Orientation: "vertical", OffsetPercent: 50, SizePercent: 50,
			},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/line.json")
	})
}
