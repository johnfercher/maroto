package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/config"
)

var redColor = props.Color{
	Red:   255,
	Green: 0,
	Blue:  0,
}

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(40,
		line.NewCol(2),
		line.NewCol(4),
		line.NewCol(6),
	)

	m.AddRow(40,
		line.NewCol(6),
		line.NewCol(4),
		line.NewCol(2),
	)

	m.AddRow(40,
		line.NewCol(2, props.Line{Thickness: 0.5}),
		line.NewCol(4, props.Line{Color: redColor}),
		line.NewCol(6, props.Line{Orientation: orientation.Vertical}),
	)

	m.AddRow(40,
		line.NewCol(6, props.Line{OffsetPercent: 50}),
		line.NewCol(4, props.Line{OffsetPercent: 50, Orientation: orientation.Vertical}),
		line.NewCol(2, props.Line{SizePercent: 50}),
	)

	m.AddRow(40,
		line.NewCol(2, props.Line{Style: linestyle.Dashed}),
		line.NewCol(4, props.Line{Color: redColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Vertical, OffsetPercent: 70, SizePercent: 70}),
		line.NewCol(6, props.Line{Color: redColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Horizontal, OffsetPercent: 40, SizePercent: 40}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/linegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/linegridv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
