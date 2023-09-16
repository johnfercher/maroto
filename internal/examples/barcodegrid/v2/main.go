package main

import (
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 50,
				}),
			),
			col.New(4).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 75,
				}),
			),
			col.New(6).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 100,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 50,
				}),
			),
			col.New(4).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 75,
				}),
			),
			col.New(6).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 100,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(6).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 50,
				}),
			),
			col.New(4).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 75,
				}),
			),
			col.New(2).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Percent: 100,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(6).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 50,
				}),
			),
			col.New(4).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 75,
				}),
			),
			col.New(2).Add(
				code.NewBar("https://github.com/johnfercher/maroto", props.Barcode{
					Center:  true,
					Percent: 100,
				}),
			),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("internal/examples/pdfs/barcodegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
