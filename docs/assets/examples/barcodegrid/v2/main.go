package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/code"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/props"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(40,
		code.NewBarCol(2, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 50,
		}),
		code.NewBarCol(4, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 75,
		}),
		code.NewBarCol(6, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewBarCol(2, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 50,
		}),
		code.NewBarCol(4, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 75,
		}),
		code.NewBarCol(6, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewBarCol(6, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 50,
		}),
		code.NewBarCol(4, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 75,
		}),
		code.NewBarCol(2, "https://github.com/johnfercher/maroto", props.Barcode{
			Percent: 100,
		}),
	)

	m.AddRow(40,
		code.NewBarCol(6, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 50,
		}),
		code.NewBarCol(4, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 75,
		}),
		code.NewBarCol(2, "https://github.com/johnfercher/maroto", props.Barcode{
			Center:  true,
			Percent: 100,
		}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/barcodegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
