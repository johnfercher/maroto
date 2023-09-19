package main

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"log"
	"math/rand"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/text"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(false).
		Build()

	colStyle := &props.Style{
		BackgroundColor: &color.Color{80, 80, 80},
		Border:          border.Full,
		BorderColor:     &color.Color{200, 0, 0},
	}

	rowStyles := []*props.Style{
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.None,
			BorderColor:     &color.Color{0, 0, 200},
		},
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.Full,
			BorderColor:     &color.Color{0, 0, 200},
		},
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.Left,
			BorderColor:     &color.Color{0, 0, 200},
		},
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.Right,
			BorderColor:     &color.Color{0, 0, 200},
		},
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.Top,
			BorderColor:     &color.Color{0, 0, 200},
		},
		{
			BackgroundColor: &color.Color{220, 220, 220},
			Border:          border.Bottom,
			BorderColor:     &color.Color{0, 0, 200},
		},
	}

	whiteText := props.Text{
		Color: &color.Color{255, 255, 255},
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

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

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
			row.New(10).WithStyle(rowStyles[rand.Intn(len(rowStyles))]).Add(
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
				text.NewCol(4, "string", blackText),
			),
		)

		m.AddRows(row.New(10))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/cellstylev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
