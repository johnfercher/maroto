package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/grid/col"
	"github.com/johnfercher/v2/maroto/image"
	"github.com/johnfercher/v2/maroto/props"
	"github.com/johnfercher/v2/maroto/text"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(&config.Dimensions{
			200, 200,
		}).
		WithDebug(true).
		Build()

	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(40,
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		text.NewCol(4, "Gopher International Shipping, Inc.", props.Text{
			Top:         12,
			Size:        12,
			Extrapolate: true,
		}),
		col.New(4),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/customsizev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
