package main

import (
	"log"
	"time"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithAuthor("author", false).
		WithCreator("creator", false).
		WithSubject("subject", false).
		WithTitle("title", false).
		WithCreationDate(time.Now()).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(30, "metadatas"),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/metadatasv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/metadatasv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
