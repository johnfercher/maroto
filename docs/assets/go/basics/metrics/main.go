package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/pkg"
)

func main() {
	mrt := pkg.NewMaroto()
	m := pkg.NewMetricsDecorator(mrt)

	// AddRows things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(document.GetReport().String())
}
