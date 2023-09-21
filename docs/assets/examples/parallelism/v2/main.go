package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithWorkerPoolSize(7).
		WithDebug(true).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

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

	err = document.Save("docs/assets/pdf/parallelismv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/parallelismv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
