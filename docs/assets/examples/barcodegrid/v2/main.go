package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/code"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

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

	err = document.GetReport().Save("docs/assets/text/barcodegridv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
