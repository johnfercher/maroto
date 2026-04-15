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

	rtCol := richtext.NewCol(
		12,
		richtext.NewChunk("Price: "),
		richtext.NewChunk("$99.99", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
	)
	m.AddRow(10, rtCol)

	// generate document
}

// ExampleNewAutoRow demonstrates how to create a richtext component with automatic row height.
func ExampleNewAutoRow() {
	m := maroto.New()

	rtRow := richtext.NewAutoRow(
		richtext.NewChunk("This paragraph has "),
		richtext.NewChunk("mixed styles", props.Text{Style: fontstyle.Bold}),
		richtext.NewChunk(" and wraps as one flow."),
	)
	m.AddRows(rtRow)

	// generate document
}
