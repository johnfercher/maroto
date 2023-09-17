package main

import (
	"log"

	v2 "github.com/johnfercher/maroto/pkg/v2"
)

func main() {
	m := v2.NewMaroto()
	mMetrified := v2.NewMetricsDecorator(m)

	// AddRows things

	document, err := mMetrified.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
