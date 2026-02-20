package richtext_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a richtext component.
func ExampleNew() {
	m := maroto.New()

	rt := richtext.New(
		richtext.NewChunk("This is "),
		richtext.NewChunk("bold", props.Text{Style: fontstyle.Bold}),
		richtext.NewChunk(" text."),
	)
	c := col.New(12).Add(rt)
	m.AddRow(10, c)

	// generate document
}

// ExampleNewCol demonstrates how to create a richtext component wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	rtCol := richtext.NewCol(12,
		richtext.NewChunk("Price: "),
		richtext.NewChunk("$99.99", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
	)
	m.AddRow(10, rtCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a richtext component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	rtRow := richtext.NewRow(10,
		richtext.NewChunk("Hello "),
		richtext.NewChunk("World", props.Text{Style: fontstyle.BoldItalic}),
	)
	m.AddRows(rtRow)

	// generate document
}

// ExampleNewAutoRow demonstrates how to create a richtext component wrapped into an auto-height row.
func ExampleNewAutoRow() {
	m := maroto.New()

	rtRow := richtext.NewAutoRow(
		richtext.NewChunk("This is a paragraph with "),
		richtext.NewChunk("mixed styles", props.Text{Style: fontstyle.Bold}),
		richtext.NewChunk(" that auto-sizes its row height."),
	)
	m.AddRows(rtRow)

	// generate document
}
