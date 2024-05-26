package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	savedPdf, err := os.ReadFile("docs/assets/pdf/v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Merge(savedPdf)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/mergepdfv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/mergepdfv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	for i := 0; i < 50; i++ {
		m.AddRows(text.NewRow(20, "content"))
	}

	return m
}
