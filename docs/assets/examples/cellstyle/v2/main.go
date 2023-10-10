package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/cellstylev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/cellstylev2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(false).
		Build()

	colStyle := &props.Cell{
		BackgroundColor: &props.Color{80, 80, 80},
		BorderType:      border.Full,
		BorderColor:     &props.Color{200, 0, 0},
		LineStyle:       linestyle.Dashed,
		BorderThickness: 0.5,
	}

	rowStyles := []*props.Cell{
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.None,
			BorderColor:     &props.Color{0, 0, 200},
		},
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.Full,
			BorderColor:     &props.Color{0, 0, 200},
		},
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.Left,
			BorderColor:     &props.Color{0, 0, 200},
		},
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.Right,
			BorderColor:     &props.Color{0, 0, 200},
		},
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.Top,
			BorderColor:     &props.Color{0, 0, 200},
		},
		{
			BackgroundColor: &props.Color{220, 220, 220},
			BorderType:      border.Bottom,
			BorderColor:     &props.Color{0, 0, 200},
		},
	}

	whiteText := props.Text{
		Color: &props.Color{255, 255, 255},
		Style: fontstyle.Bold,
		Size:  12,
		Align: align.Center,
		Top:   2,
	}

	blackText := props.Text{
		Style: fontstyle.Bold,
		Size:  12,
		Align: align.Center,
		Top:   2,
	}

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	count := 0
	for i := 0; i < 15; i++ {
		m.AddRows(
			row.New(10).Add(
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
				text.NewCol(4, "string", whiteText).WithStyle(colStyle),
			),
		)

		m.AddRows(row.New(10))

		m.AddRows(
			row.New(10).WithStyle(rowStyles[count]).Add(
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
			),
		)

		m.AddRows(row.New(10))
		count++
		if count >= len(rowStyles) {
			count = 0
		}
	}
	return m
}
