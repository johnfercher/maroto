package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"log"
)

func main() {
	m := v2.NewMaroto("v2.pdf")
	metrics := v2.NewMarotoMetrified(m)

	// Add things

	report, err := metrics.GenerateWithReport()
	if err != nil {
		log.Fatal(err.Error())
	}

	report.Print()
}
