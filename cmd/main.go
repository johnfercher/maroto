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
	cfg := config.NewBuilder().
		WithThreadPool(10).
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)

	m := v2.NewMetricsDecorator(maroto)

	err := m.RegisterHeader(buildHeader()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _ = range [10]int{} {
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
	return row.New(40).Add(
		col.New(4).Add(
			code.NewBar("barcode"),
		),
		col.New(4).Add(
			code.NewQr("qrcode"),
		),
		col.New(4).Add(
			code.NewMatrix("matrixcode"),
		),
	)
}

func buildImagesRow() domain.Row {
	byteSlices, err := os.ReadFile("internal/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)

	return row.New(40).Add(
		col.New(6).Add(
			image.NewFromBase64(stringBase64, consts.Png),
		),
		col.New(6).Add(
			image.NewFromFile("internal/assets/images/frontpage.png"),
		),
	)
}

func buildTextsRow() domain.Row {
	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	return row.New(40).Add(
		col.New(6).Add(
			text.New(colText, props.Text{
				Align: consts.Center,
			}),
		),
		col.New(6).Add(
			signature.New("signature", props.Font{
				Style:  consts.Italic,
				Size:   15,
				Family: consts.Courier,
			}),
		),
	)
}

func buildHeader() []domain.Row {
	r1 := row.New(15).Add(
		col.New(12).Add(
			text.New("Maroto V2", props.Text{
				Size:  15,
				Align: consts.Center,
			}),
			text.New("The New Standard", props.Text{
				Top:   8,
				Size:  13,
				Align: consts.Center,
			}),
		),
	)

	r2 := row.New(10).Add(
		col.New(2).Add(
			text.New("Site: https://maroto.io/"),
		),
		col.New(5).Add(
			text.New("Discussions: https://github.com/johnfercher/maroto/issues/257"),
		),
		col.New(5).Add(
			text.New("Branch: https://github.com/johnfercher/maroto/tree/v2"),
		),
	)

	return []domain.Row{r1, r2}
}
