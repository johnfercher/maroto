package row_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a Row instance.
func ExampleNew() {
	// height defines the size of the useful area
	// which can be used in the set of columns and components
	// inside this row.
	height := 10.0
	row := row.New(height)

	m := maroto.New()

	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleRow_Add demonstrates how to add cols inside a Row.
func ExampleRow_Add() {
	textCol := text.NewCol(12, "text content")
	qrCodeCol := code.NewQrCol(12, "qrcode")
	signatureCol := signature.NewCol(12, "signature label")

	row := row.New(10).Add(textCol, qrCodeCol, signatureCol)

	m := maroto.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleRow_WithStyle demonstrates how to add style to Row.
func ExampleRow_WithStyle() {
	row := row.New(10)

	row.WithStyle(&props.Cell{
		BackgroundColor: &props.Color{
			Red:   10,
			Green: 100,
			Blue:  150,
		},
		BorderColor: &props.Color{
			Red:   55,
			Green: 10,
			Blue:  60,
		},
		BorderType:      border.Full,
		BorderThickness: 0.1,
		LineStyle:       linestyle.Dashed,
	})

	m := maroto.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}
