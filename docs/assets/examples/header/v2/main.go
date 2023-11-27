package main

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/headerv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/headerv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	// start header from page 1
	headerConfig := entity.HeaderConfig{StartPage: 1}

	cfg := config.NewBuilder().
		WithDebug(true).
		WithHeaderConfig(&headerConfig).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(text.NewRow(20, "Header", props.Text{
		Size:  10,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 70; i++ {
		m.AddRows(
			text.NewRow(10, "Dummy text", props.Text{
				Size: 8,
			}),
		)
	}

	return m
}
