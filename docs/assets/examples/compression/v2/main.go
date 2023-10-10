package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func main() {
	m := GetMaroto("docs/assets/images/frontpage.png")
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/compressionv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/compressionv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(imagePath string) core.Maroto {
	cfg := config.NewBuilder().
		WithCompression(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(20, "Main features", props.Text{Size: 15, Top: 6.5}),
	)

	m.AddRows(
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
	)

	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	m.AddRows(
		row.New(20).Add(
			text.NewCol(4, "Image From File:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", props.Rect{Center: true, Percent: 90}),
		),
		row.New(20).Add(
			text.NewCol(4, "Image From Base64::", props.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromBytesCol(8, bytes, extension.Png, props.Rect{Center: true, Percent: 90}),
		),
	)

	m.AddRows(
		row.New(20).Add(
			text.NewCol(4, "Text:", props.Text{Size: 15, Top: 6, Align: align.Center}),
			text.NewCol(8, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem.", props.Text{Size: 12, Top: 5, Align: align.Center}),
		),
		row.New(40).Add(
			text.NewCol(4, "Signature:", props.Text{Size: 15, Top: 17, Align: align.Center}),
			signature.NewCol(8, "Name", props.Signature{FontSize: 10}),
		),
	)

	return m
}

func buildImagesRow() []core.Row {
	return []core.Row{}
}

func buildTextsRow() []core.Row {
	return []core.Row{}
}
