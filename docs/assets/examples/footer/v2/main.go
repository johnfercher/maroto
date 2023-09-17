package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	err := m.RegisterFooter(text.NewRow(20, "Footer", props.Text{
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

	err = document.Save("docs/assets/pdf/footerv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
