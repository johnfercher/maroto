package maroto_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"log"
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

	m.AddRow(10, code.NewBarCol(12, "barcode"))

	// Do things and generate
}

// ExampleMaroto_AddRows demonstrates how to add new rows in maroto.
func ExampleMaroto_AddRows() {
	m := maroto.New()

	m.AddRows(
		code.NewBarRow(12, "barcode"),
		code.NewQrRow(12, "barcode"),
		code.NewMatrixRow(12, "barcode"),
	)

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
