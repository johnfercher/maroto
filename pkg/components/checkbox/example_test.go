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

	checkedProps := props.Checkbox{Checked: true}
	uncheckedProps := props.Checkbox{Checked: false}

	maleCheckbox := checkbox.New("Male", checkedProps)
	femaleCheckbox := checkbox.New("Female", uncheckedProps)

	maleCol := col.New(12).Add(maleCheckbox)
	femaleCol := col.New(12).Add(femaleCheckbox)

	m.AddRow(7, maleCol)
	m.AddRow(7, femaleCol)

	// generate document
}

// ExampleNewCol demonstrates how to create a checkbox component wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	checkedProps := props.Checkbox{Checked: true}
	uncheckedProps := props.Checkbox{Checked: false}

	maleCheckboxCol := checkbox.NewCol(12, "Male", checkedProps)
	femaleCheckboxCol := checkbox.NewCol(12, "Female", uncheckedProps)

	m.AddRow(7, maleCheckboxCol)
	m.AddRow(7, femaleCheckboxCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a checkbox component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	checkedProps := props.Checkbox{Checked: true}
	uncheckedProps := props.Checkbox{Checked: false}

	maleCheckboxRow := checkbox.NewRow(7, "Male", checkedProps)
	femaleCheckboxRow := checkbox.NewRow(7, "Female", uncheckedProps)

	m.AddRows(maleCheckboxRow, femaleCheckboxRow)

	// generate document
}
