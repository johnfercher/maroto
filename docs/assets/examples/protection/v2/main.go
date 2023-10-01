package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"

	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithProtection(protection.None, "user", "owner").
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(30, "supersecret content"),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/protectionv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/protectionv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
