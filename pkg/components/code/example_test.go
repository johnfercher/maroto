package code_test

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
)

// ExampleNewBar demonstrates how to generate a barcode and add it to maroto
func ExampleNewBar() {
	mrt := maroto.New()
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(10, col.New(6).Add(code.NewBar("123456789", props.Barcode{Percent: 70.5})))

	// generate document
}

// ExampleNewBar demonstrates how to generate a column with a barcode and add it to maroto
func ExampleNewBarCol() {
	mrt := maroto.New()
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(10, code.NewBarCol(6, "123456", props.Barcode{Percent: 70.5}))

	// generate document
}

// ExampleNewBarRow demonstrates how to generate a row with a barcode and add it to maroto
func ExampleNewBarRow() {
	mrt := maroto.New()
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(code.NewBarRow(10, "123456789", props.Barcode{Percent: 70.5}))

	// generate document
}
