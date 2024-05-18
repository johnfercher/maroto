package code_test

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
)

// ExampleNewBar demonstrates how to generate a barcode and add it to maroto.
func ExampleNewBar() {
	m := maroto.New()

	barCode := code.NewBar("123456789", props.Barcode{Percent: 70.5})
	col := col.New(6).Add(barCode)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewBarCol demonstrates how to generate a column with a barcode and add it to maroto.
func ExampleNewBarCol() {
	m := maroto.New()

	barCodeCol := code.NewBarCol(6, "123456", props.Barcode{Percent: 70.5})
	m.AddRow(10, barCodeCol)

	// generate document
}

// ExampleNewBarRow demonstrates how to generate a row with a barcode and add it to maroto.
func ExampleNewBarRow() {
	m := maroto.New()

	barCodeRow := code.NewBarRow(10, "123456789", props.Barcode{Percent: 70.5})
	m.AddRows(barCodeRow)

	// generate document
}

// ExampleNewQr demonstrates how to generate a qrcode and add it to maroto.
func ExampleNewQr() {
	m := maroto.New()

	qrCode := code.NewQr("123456789", props.Rect{Percent: 70.5})
	col := col.New(6).Add(qrCode)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewQrCol demonstrates how to generate a column with a qrcode and add it to maroto.
func ExampleNewQrCol() {
	m := maroto.New()

	qrCodeCol := code.NewQrCol(12, "123456789", props.Rect{Percent: 70.5})
	m.AddRow(10, qrCodeCol)

	// generate document
}

// ExampleNewQrRow demonstrates how to generate a row with a qrcode and add it to maroto.
func ExampleNewQrRow() {
	m := maroto.New()

	qrCodeRow := code.NewQrRow(10, "123456789", props.Rect{Percent: 70.5})
	m.AddRows(qrCodeRow)

	// generate document
}

// ExampleNewMatrix demonstrates how to generate a matrixcode and add it to maroto.
func ExampleNewMatrix() {
	m := maroto.New()

	matrixCode := code.NewMatrix("123456789", props.Rect{Percent: 70.5})
	col := col.New(6).Add(matrixCode)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewMatrixCol demonstrates how to generate a column with a matrixcode and add it to maroto.
func ExampleNewMatrixCol() {
	m := maroto.New()

	matrixCodeCol := code.NewMatrixCol(12, "123456789", props.Rect{Percent: 70.5})
	m.AddRow(10, matrixCodeCol)

	// generate document
}

// ExampleNewMatrixRow demonstrates how to generate a row with a matrixcode and add it to maroto.
func ExampleNewMatrixRow() {
	m := maroto.New()

	matrixCodeRow := code.NewMatrixRow(10, "123456789", props.Rect{Percent: 70.5})
	m.AddRows(matrixCodeRow)

	// generate document
}
