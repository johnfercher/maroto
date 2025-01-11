package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/pagenumberv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/pagenumberv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	pageNumber := props.PageNumber{
		Pattern:    "Page {current} of {total}",
		Place:      props.Bottom,
		Family:     fontfamily.Courier,
		Style:      fontstyle.Bold,
		MarginTop:  26,
		MarginLeft: 188,
		Size:       9,
		Color: &props.Color{
			Red: 255,
		},
	}

	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageNumber(pageNumber).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	for i := 0; i < 15; i++ {
		m.AddRows(text.NewRow(20, "dummy text"))
	}

	return m
}
