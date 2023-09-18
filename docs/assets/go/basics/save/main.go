package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg"
)

func main() {
	m := pkg.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("file.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}
