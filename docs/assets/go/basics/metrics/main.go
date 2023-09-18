package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"log"
)

func main() {
	mrt := pkg.NewMaroto()
	m := pkg.NewMetricsDecorator(mrt)

	// AddRows things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
