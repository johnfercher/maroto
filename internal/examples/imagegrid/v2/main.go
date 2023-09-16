package main

import (
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/image"
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
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
			col.New(6).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  false,
					Percent: 50,
					Left:    10,
				}),
			),
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  false,
					Percent: 50,
					Top:     10,
				}),
			),
			col.New(6).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  false,
					Percent: 50,
					Left:    15,
					Top:     15,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(8).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(6).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  false,
					Percent: 80,
					Top:     5,
					Left:    10,
				}),
			),
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  false,
					Percent: 80,
					Top:     5,
				}),
			),
			col.New(2).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  false,
					Percent: 80,
					Left:    5,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(6).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  true,
					Percent: 50,
				}),
			),
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  true,
					Percent: 50,
				}),
			),
			col.New(2).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  true,
					Percent: 50,
				}),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(4).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
			col.New(8).Add(
				image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
					Center:  true,
					Percent: 80,
				}),
			),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("internal/examples/pdfs/imagegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
