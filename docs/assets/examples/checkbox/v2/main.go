package main

import (
	"log"

	maroto "github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/checkbox"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/checkboxv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/checkboxv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	// Default: unchecked and checked with label
	m.AddRow(20,
		checkbox.NewCol(4, "Option A"),
		checkbox.NewCol(4, "Option B", props.Checkbox{Checked: true}),
		checkbox.NewCol(4, "Option C"),
	)
	m.AddRows(text.NewRow(8, "Default: unchecked / checked / unchecked"))

	// Custom size
	m.AddRow(20,
		checkbox.NewCol(4, "Option A", props.Checkbox{Size: 8}),
		checkbox.NewCol(4, "Option B", props.Checkbox{Size: 8, Checked: true}),
		checkbox.NewCol(4, "Option C", props.Checkbox{Size: 8}),
	)
	m.AddRows(text.NewRow(8, "Custom size: 8mm"))

	// With top and left offset
	m.AddRow(20,
		checkbox.NewCol(4, "Option A", props.Checkbox{Top: 5, Left: 5}),
		checkbox.NewCol(4, "Option B", props.Checkbox{Top: 5, Left: 5, Checked: true}),
		checkbox.NewCol(4, "Option C", props.Checkbox{Top: 5, Left: 5}),
	)
	m.AddRows(text.NewRow(8, "With top=5 left=5 offset"))

	// Auto row
	m.AddAutoRow(
		checkbox.NewCol(3, "Item 1"),
		checkbox.NewCol(3, "Item 2", props.Checkbox{Checked: true}),
		checkbox.NewCol(3, "Item 3"),
		checkbox.NewCol(3, "Item 4", props.Checkbox{Checked: true}),
	)
	m.AddRows(text.NewRow(8, "Auto row"))

	return m
}
