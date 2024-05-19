package line_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
)

// ExampleNew demonstrates how create a line component.
func ExampleNew() {
	m := maroto.New()

	line := line.New()
	col := col.New(12).Add(line)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewCol demonstrates how to crete a line wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	lineCol := line.NewCol(12)
	m.AddRow(10, lineCol)

	// generate document
}

// ExampleNewRow demonstrates how to crete a line wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	lineRow := line.NewRow(10)
	m.AddRows(lineRow)

	// generate document
}
