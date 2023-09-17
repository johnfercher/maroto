package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	gridSum := 14
	cfg := config.NewBuilder().
		WithDebug(true).
		WithMaxGridSize(gridSum).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	m.AddRows(text.NewRow(10, fmt.Sprintf("Table with %d Columns", gridSum), props.Text{Style: consts.Bold}))

	var headers []domain.Col
	var contents []domain.Col
	for i := 0; i < gridSum; i++ {
		headers = append(headers, text.NewCol(1, fmt.Sprintf("H %d", i), props.Text{Style: consts.Bold, Top: 1.5, Left: 1.5}))
		contents = append(contents, text.NewCol(1, fmt.Sprintf("C %d", i), props.Text{Top: 1, Left: 1.5, Size: 9}))
	}

	m.AddRow(8, headers...)
	m.AddRow(8, contents...)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/maxgridsumv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
