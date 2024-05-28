package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/row"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/marginsv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/marginsv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithTopMargin(20).
		WithLeftMargin(20).
		WithRightMargin(20).
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(
		row.New(40).Add(
			image.NewFromFileCol(4, "docs/assets/images/gopherbw.png", props.Rect{
				Center:  true,
				Percent: 50,
			}),
			text.NewCol(4, "Margins Test", props.Text{
				Top:  12,
				Size: 12,
			}),
			col.New(4),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		m.AddRows(text.NewRow(30, "any text"))
	}

	return m
}
