package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithMargins(&config.Margins{
			Left:  20,
			Right: 20,
			Top:   20,
		}).
		WithDebug(true).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	m.AddRow(40,
		image.NewFromFileCol(4, "docs/assets/images/gopherbw.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		text.NewCol(4, "Margins Test", props.Text{
			Top:  12,
			Size: 12,
		}),
		col.New(4),
	)

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
