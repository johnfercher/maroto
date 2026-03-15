package checkbox_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/checkbox"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a checkbox component.
func ExampleNew() {
	m := maroto.New()

	checkbox := checkbox.New("checkbox label", props.Checkbox{
		Checked: true,
	})
	col := col.New(12).Add(checkbox)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewCol demonstrates how to create a checkbox component wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	checkboxCol := checkbox.NewCol(12, "checkbox label")
	m.AddRow(10, checkboxCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a checkbox component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	checkboxRow := checkbox.NewRow(10, "checkbox label")
	m.AddRows(checkboxRow)

	// generate document
}

// ExampleNewAutoRow demonstrates how to create a checkbox component wrapped into a row with automatic height.
func ExampleNewAutoRow() {
	m := maroto.New()

	checkboxRow := checkbox.NewAutoRow("checkbox label", props.Checkbox{
		Checked: true,
		Size:    5,
		Top:     2,
		Left:    2,
	})
	m.AddRows(checkboxRow)

	// generate document
}
