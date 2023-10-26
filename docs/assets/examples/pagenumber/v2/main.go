package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/pagenumberv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/pagenumberv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageNumber("Page {current} of {total}", props.Bottom).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	for i := 0; i < 15; i++ {
		m.AddRows(text.NewRow(20, "dummy text"))
	}

	return m
}
