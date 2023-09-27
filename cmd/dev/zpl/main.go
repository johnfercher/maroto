package main

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/karmdip-mi/go-fitz"
	"log"
	"os"
	"simonwaldherr.de/go/zplgfa"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func main() {
	cfg := config.NewBuilder().
		WithDimensions(101.6, 152.4).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	byteSlices, err := os.ReadFile("docs/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)

	m.AddRow(20,
		image.NewFromBase64Col(4, stringBase64, extension.Png, props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(8).Add(
			text.New("Maroto V2", props.Text{
				Size:  15,
				Style: fontstyle.Bold,
			}),
			text.New("ZPL Generator", props.Text{
				Top:   8,
				Size:  10,
				Style: fontstyle.Bold,
			}),
			text.New("Based On PDF", props.Text{
				Top:   12,
				Size:  10,
				Style: fontstyle.Bold,
			}),
		),
	)

	m.AddRows(line.NewRow(5, props.Line{OffsetPercent: 50}))

	m.AddRow(20,
		col.New(8).Add(
			text.New("Maroto V2", props.Text{
				Left:  10,
				Top:   3,
				Align: align.Left,
			}),
			text.New("ZPL Generator", props.Text{
				Left:  10,
				Top:   8,
				Align: align.Left,
			}),
			text.New("Based On PDF", props.Text{
				Left:  10,
				Top:   13,
				Align: align.Left,
			}),
		),
		code.NewQrCol(4, "qrcode", props.Rect{
			Center: true,
		}),
	)

	m.AddRows(line.NewRow(5, props.Line{OffsetPercent: 50}))

	m.AddRows(code.NewBarRow(20, "code", props.Barcode{
		Percent: 70,
		Center:  true,
	}))

	m.AddRows(text.NewRow(20, "code123"))

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	docBytes := document.GetBytes()

	doc, err := fitz.NewFromMemory(docBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		flat := zplgfa.FlattenImage(img)
		gfimg := zplgfa.ConvertToZPL(flat, zplgfa.CompressedASCII)
		err = os.WriteFile("test.zpl", []byte(gfimg), os.ModePerm)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func buildCodesRow() []core.Row {
	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Barcode:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewBarCol(8, "barcode", props.Barcode{Center: true, Percent: 70}),
		),
		row.New(20).Add(
			text.NewCol(4, "QrCode:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewQrCol(8, "qrcode", props.Rect{Center: true, Percent: 70}),
		),
		row.New(20).Add(
			text.NewCol(4, "MatrixCode:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewMatrixCol(8, "matrixcode", props.Rect{Center: true, Percent: 70}),
		),
	}
}

func buildImagesRow() []core.Row {
	byteSlices, err := os.ReadFile("docs/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)

	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Image From File:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", props.Rect{Center: true, Percent: 90}),
		),
		row.New(20).Add(
			text.NewCol(4, "Image From Base64::", props.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromBase64Col(8, stringBase64, extension.Png, props.Rect{Center: true, Percent: 90}),
		),
	}
}

func buildTextsRow() []core.Row {
	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Text:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			text.NewCol(8, colText, props.Text{Size: 12, Top: 5, Align: align.Center}),
		),
		row.New(40).Add(
			text.NewCol(4, "Signature:", props.Text{Size: 15, Top: 17, Align: align.Center}),
			signature.NewCol(8, "Name", props.Font{Size: 10}),
		),
	}
}
