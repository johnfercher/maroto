package main

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/signature"
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
				signature.New("Signature 1"),
			),
			col.New(4).Add(
				signature.New("Signature 2", props.Font{Family: consts.Courier}),
			),
			col.New(6).Add(
				signature.New("Signature 3", props.Font{Style: consts.BoldItalic}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(6).Add(
				signature.New("Signature 4", props.Font{Style: consts.Italic}),
			),
			col.New(4).Add(
				signature.New("Signature 5", props.Font{Size: 12}),
			),
			col.New(2).Add(
				signature.New("Signature 6", props.Font{Color: color.Color{255, 0, 0}}),
			),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("internal/examples/pdfs/signaturegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
