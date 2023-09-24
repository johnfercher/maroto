package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/metrics"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

var background = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

func main() {
	var content string
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		content += fmt.Sprintf("%f", run().Value) + "\n"
	}

	err := os.WriteFile("docs/assets/text/benchmark.txt", []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func run() *metrics.Time {
	cfg := config.NewBuilder().
		WithPageNumber("Page {current} of {total}", props.South).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(buildHeader()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(buildFooter()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(
		text.NewRow(20, "Main features", props.Text{Size: 15, Top: 6.5}),
	)

	objects := getObjects(1158)
	rows, err := list.Build[Object](objects)
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(rows...)

	for i := 0; i < 1158; i++ {
		m.AddRows(buildCodesRow()...)
		m.AddRows(buildImagesRow()...)
		m.AddRows(buildTextsRow()...)
	}

	m.AddRows(
		text.NewRow(15, "Dummy Data", props.Text{Size: 12, Top: 5, Align: align.Center}),
	)

	for i := 0; i < 1158; i++ {
		m.AddRows(text.NewRow(20, dummyText+dummyText+dummyText+dummyText+dummyText))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	return document.GetReport().TimeMetrics[0].Avg
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

func buildHeader() []core.Row {
	r1 := row.New(30).Add(
		col.New(12).Add(
			text.New("Config V2", props.Text{
				Top:   5,
				Size:  15,
				Align: align.Center,
			}),
			text.New("Grid system, fast generation, embedded metrics and testable.", props.Text{
				Top:   13,
				Size:  13,
				Align: align.Center,
			}),
		),
	)

	return []core.Row{r1}
}

func buildFooter() []core.Row {
	return []core.Row{
		row.New(10).Add(
			text.NewCol(2, "Site: https://maroto.io/"),
			text.NewCol(5, "Discussions: https://github.com/johnfercher/maroto/issues/257"),
			text.NewCol(5, "Branch: https://github.com/johnfercher/maroto/tree/v2"),
		),
	}
}

type Object struct {
	Key   string
	Value string
}

func (o Object) GetHeader() core.Row {
	return row.New(10).Add(
		text.NewCol(4, "Key", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "Value", props.Text{Style: fontstyle.Bold}),
	)
}

func (o Object) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(4, o.Key),
		text.NewCol(8, o.Value),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: background,
		})
	}

	return r
}

func getObjects(max int) []Object {
	var objects []Object
	for i := 0; i < max; i++ {
		objects = append(objects, Object{
			Key:   fmt.Sprintf("Key: %d", i),
			Value: fmt.Sprintf("Value: %d", i),
		})
	}
	return objects
}
