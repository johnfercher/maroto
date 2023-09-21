package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	gridSum := 14
	cfg := config.NewBuilder().
		WithDebug(true).
		WithMaxGridSize(gridSum).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	m.AddRows(text.NewRow(10, fmt.Sprintf("Table with %d Columns", gridSum), props.Text{Style: fontstyle.Bold}))

	var headers []core.Col
	var contents []core.Col
	for i := 0; i < gridSum; i++ {
		headers = append(headers, text.NewCol(1, fmt.Sprintf("H %d", i), props.Text{Style: fontstyle.Bold, Top: 1.5, Left: 1.5}))
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

	err = document.GetReport().Save("docs/assets/text/maxgridsumv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
