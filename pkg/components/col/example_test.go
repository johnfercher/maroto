package col_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a Col instance.
func ExampleNew() {
	// size is an optional parameters, if not provided, maroto
	// will apply the maximum size, even if custom size is applied.
	size := 12
	col := col.New(size)

	row := row.New(10).Add(col)

	m := maroto.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleCol_Add demonstrates how to add components to Col.
func ExampleCol_Add() {
	col := col.New()

	text := text.New("text content")
	qrCode := code.NewQr("qrcode")
	signature := signature.New("signature label")

	col.Add(text, qrCode, signature)

	row := row.New(10).Add(col)

	m := maroto.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleCol_WithStyle demonstrates how to add style to Col.
func ExampleCol_WithStyle() {
	col := col.New()

	col.WithStyle(&props.Cell{
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

	row := row.New(10).Add(col)

	m := maroto.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}
