package image_test

import (
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
)

// ExampleNewFromBytes demonstrates how to create an image component reading bytes.
func ExampleNewFromBytes() {
	m := maroto.New()

	bytes, _ := os.ReadFile("image.png")

	image := image.NewFromBytes(bytes, extension.Png)
	col := col.New(12).Add(image)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewFromBytesCol demonstrates how to create an image component wrapped into a column reading bytes.
func ExampleNewFromBytesCol() {
	m := maroto.New()

	bytes, _ := os.ReadFile("image.png")

	imageCol := image.NewFromBytesCol(12, bytes, extension.Png)
	m.AddRow(10, imageCol)

	// generate document
}

// ExampleNewFromBytesRow demonstrates how to create an image component wrapped into a row reading bytes.
func ExampleNewFromBytesRow() {
	m := maroto.New()

	bytes, _ := os.ReadFile("image.png")

	imageRow := image.NewFromBytesRow(10, bytes, extension.Png)
	m.AddRows(imageRow)

	// generate document
}

// ExampleNewFromFile demonstrates how to create an image component reading file.
func ExampleNewFromFile() {
	m := maroto.New()

	image := image.NewFromFile("image.png")
	col := col.New(12).Add(image)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewFromFileCol demonstrates how to create an image component wrapped into a column reading file.
func ExampleNewFromFileCol() {
	m := maroto.New()

	imageCol := image.NewFromFileCol(12, "image.png")
	m.AddRow(10, imageCol)

	// generate document
}

// ExampleNewFromFileRow demonstrates how to create an image component wrapped into a row reading file.
func ExampleNewFromFileRow() {
	m := maroto.New()

	imageRow := image.NewFromFileRow(10, "image.png")
	m.AddRows(imageRow)
	// generate document
}
