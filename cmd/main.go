package main

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code/barcode"
	"github.com/johnfercher/maroto/pkg/v2/code/matrixcode"
	"github.com/johnfercher/maroto/pkg/v2/code/qrcode"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/signature"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
	"os"
)

func main() {
	pdf := buildMarotoPDF()
	html := buildMarotoHTML()

	gen(pdf)
	gen(html)
}

func buildMarotoPDF() domain.MarotoMetrified {
	m := v2.NewMaroto("v2.pdf")
	return v2.NewMarotoMetrified(m)
}

func buildMarotoHTML() domain.MarotoMetrified {
	m := v2.NewMaroto("v2.html")
	return v2.NewMarotoMetrified(m)
}

func gen(m domain.MarotoMetrified) {
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())

	report, err := m.GenerateWithReport()
	if err != nil {
		log.Fatal(err.Error())
	}

	report.Print()
}

func buildCodesRow() domain.Row {
	r := row.New(70)

	col1 := col.New(4)
	col1.Add(barcode.New("barcode"))

	col2 := col.New(4)
	col2.Add(qrcode.New("qrcode"))

	col3 := col.New(4)
	col3.Add(matrixcode.New("matrixcode"))

	r.Add(col1, col2, col3)
	return r
}

func buildImagesRow() domain.Row {
	row := row.New(70)

	col1 := col.New(6)
	col1.Add(image.NewFromFile("internal/assets/images/biplane.jpg"))

	byteSlices, err := os.ReadFile("internal/assets/images/gopherbw.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)
	col2 := col.New(6)
	col2.Add(image.NewFromBase64(stringBase64, consts.Png))

	row.Add(col1, col2)

	return row
}

func buildTextsRow() domain.Row {
	row := row.New(70)

	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	col1 := col.New(6)
	col1.Add(text.New(colText, props.Text{
		Align: consts.Center,
	}))

	col2 := col.New(6)
	col2.Add(signature.New("Fulano de Tal", props.Font{
		Style:  consts.Italic,
		Size:   20,
		Family: consts.Courier,
	}))

	row.Add(col1, col2)

	return row
}
