package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/maroto"
)

func main() {
	mrt := maroto.NewMaroto()
	m := maroto.NewMetricsDecorator(mrt)

	// AddRows things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
