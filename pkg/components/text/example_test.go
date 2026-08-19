package text_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/rotationpivot"
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

// ExampleNewRow demonstrates how to create a text component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	textRow := text.NewRow(10, "text")
	m.AddRows(textRow)

	// generate document
}

// ExampleNew_rotated demonstrates how to rotate text and select the pivot point.
// The cell auto-expands vertically to contain the rotated bounding box.
func ExampleNew_rotated() {
	m := maroto.New()

	rotated := text.New("Draft", props.Text{
		Rotation: 30,
		RotationPivot: rotationpivot.Pivot{
			Horizontal: rotationpivot.Center,
			Vertical:   rotationpivot.Middle,
		},
	})
	m.AddRow(10, col.New(12).Add(rotated))

	// generate document
}
