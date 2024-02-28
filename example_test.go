package maroto_test

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

// ExampleNew demonstrates how to create a maroto instance.
func ExampleNew() {
	// optional
	b := config.NewBuilder()
	cfg := b.Build()

	m := maroto.New(cfg) // cfg is an optional

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleNewMetricsDecorator demonstrates how to create a maroto metrics decorator instance.
func ExampleNewMetricsDecorator() {
	// optional
	b := config.NewBuilder()
	cfg := b.Build()

	mrt := maroto.New(cfg)               // cfg is an optional
	m := maroto.NewMetricsDecorator(mrt) // decorator of maroto

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleMaroto_AddPages demonstrates how to add a new page in maroto.
func ExampleMaroto_AddPages() {
	m := maroto.New()

	p := page.New()
	p.Add(code.NewBarRow(10, "barcode"))

	m.AddPages(p)

	// Do things and generate
}

// ExampleMaroto_AddRow demonstrates how to add a new row in maroto.
func ExampleMaroto_AddRow() {
	m := maroto.New()

	m.AddRow(10, text.NewCol(12, "text"))

	// Do things and generate
}

// ExampleMaroto_AddRows demonstrates how to add new rows in maroto.
func ExampleMaroto_AddRows() {
	m := maroto.New()

	m.AddRows(
		code.NewBarRow(12, "barcode"),
		text.NewRow(12, "text"),
	)

	// Do things and generate
}

// ExampleMaroto_RegisterFooter demonstrates how to register a footer to me added in every new page.
// An error is returned if the area occupied by the footer is greater than the page area.
func ExampleMaroto_RegisterFooter() {
	m := maroto.New()

	err := m.RegisterFooter(
		code.NewBarRow(12, "barcode"),
		text.NewRow(12, "text"))
	if err != nil {
		panic(err)
	}

	// Do things and generate
}

// ExampleMaroto_Generate demonstrates how to generate a file.
func ExampleMaroto_Generate() {
	m := maroto.New()

	// Add rows, pages and etc.

	doc, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	// You can retrieve as Base64, Save file, Merge with another file or GetReport.
	_ = doc.GetBytes()
}

// ExampleMaroto_GetCurrentHeight demonstrates how to get current height of rows on page.
func ExampleMaroto_GetCurrentHeight() {
	m := maroto.New()

	// Get dimensions

	currentHeight := m.GetCurrentHeight()

	fmt.Printf("Current Heigh: %f", currentHeight)
}

// ExampleMaroto_GetDimensions demonstrates how to get configured width and height dimensions.
func ExampleMaroto_GetDimensions() {
	m := maroto.New()

	// Get dimensions

	dimensions := m.GetDimensions()

	// You can access Width and Height fields
	fmt.Printf("Width: %f, Height: %f", dimensions.Width, dimensions.Height)
}
