package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"log"
)

func main() {
	m := v2.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes := document.GetBytes()
}
