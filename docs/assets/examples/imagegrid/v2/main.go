package main

import (
	"log"

	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/image"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Top:     10,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    15,
			Top:     15,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left:    5,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(8, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/imagegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
