package main

import (
	"log"

	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	m.AddRow(40,
		code.NewMatrixCol(2, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 50,
		}),
		code.NewMatrixCol(4, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 75,
		}),
		code.NewMatrixCol(6, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewMatrixCol(2, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		code.NewMatrixCol(4, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 75,
		}),
		code.NewMatrixCol(6, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewMatrixCol(6, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 50,
		}),
		code.NewMatrixCol(4, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 75,
		}),
		code.NewMatrixCol(2, "https://github.com/johnfercher/maroto", props.Rect{
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewMatrixCol(6, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		code.NewMatrixCol(4, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 75,
		}),
		code.NewMatrixCol(2, "https://github.com/johnfercher/maroto", props.Rect{
			Center:  true,
			Percent: 100,
		}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/datamatrixgridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
