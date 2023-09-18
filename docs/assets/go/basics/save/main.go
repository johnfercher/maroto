package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"log"
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
