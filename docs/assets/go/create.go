package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"log"
)

func main() {
	m := v2.NewMaroto("v2.pdf")

	// Add things

	err := m.GenerateConcurrently()
	if err != nil {
		log.Fatal(err.Error())
	}
}
