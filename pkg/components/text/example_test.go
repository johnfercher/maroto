package text_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a text component.
func ExampleNew() {
	m := maroto.New()

	text := text.New("text")
	col := col.New(12).Add(text)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewCol demonstrates how to create a text component wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	textCol := text.NewCol(12, "text")
	m.AddRow(10, textCol)

	// generate document
}

// ExampleNewCol_rtl demonstrates how to create a text component that renders
// right to left text, such as Arabic, in the correct visual order.
func ExampleNewCol_rtl() {
	m := maroto.New()

	textCol := text.NewCol(12, "مرحبا بالعالم", props.Text{RTL: true})
	m.AddRow(10, textCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a text component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	textRow := text.NewRow(10, "text")
	m.AddRows(textRow)

	// generate document
}
