package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func main() {
	cfg := config.NewBuilder().
		WithPageNumber("{current} / {total}", props.SouthEast).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	pdf, err := os.ReadFile("docs/assets/pdf/v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddPDFs(pdf)

	for i := 0; i < 50; i++ {
		m.AddRows(text.NewRow(20, "content"))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/addpdfv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/addpdfv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
