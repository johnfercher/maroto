package main

import (
	"github.com/johnfercher/maroto/v2/maroto"
	"log"
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
