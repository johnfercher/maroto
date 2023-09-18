package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/consts"
	"github.com/johnfercher/v2/maroto/props"
	"github.com/johnfercher/v2/maroto/text"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(text.NewRow(20, "Header", props.Text{
		Size:  10,
		Style: consts.Bold,
		Align: consts.Center,
	}))

	for i := 0; i < 50; i++ {
		m.AddRows(
			text.NewRow(10, "Dummy text", props.Text{
				Size: 8,
			}),
		)
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/headerv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
