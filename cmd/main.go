package main

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/config"
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
	cfg := config.NewBuilder().WithWorkerPoolSize(7).Build()

	maroto := v2.NewMaroto(cfg)

	m := v2.NewMetricsDecorator(maroto)

	for _ = range [100]int{} {
		m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}

func buildCodesRow() domain.Row {
	r := row.New(70)

	col1 := col.New(4)
	col1.Add(code.NewBar("barcode"))

	col2 := col.New(4)
	col2.Add(code.NewQr("qrcode"))

	col3 := col.New(4)
	col3.Add(code.NewMatrix("matrixcode"))

	r.Add(col1, col2, col3)
	return r
}

func buildImagesRow() domain.Row {
	r := row.New(70)

	col1 := col.New(6)
	col1.Add(image.NewFromFile("internal/assets/images/frontpage.png"))

	byteSlices, err := os.ReadFile("internal/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)
	col2 := col.New(6)
	col2.Add(image.NewFromBase64(stringBase64, consts.Png))

	r.Add(col1, col2)
	return r
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
