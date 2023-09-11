package main

import (
	"github.com/johnfercher/maroto/pkg/v22"
	"github.com/johnfercher/maroto/pkg/v22/col"
	"github.com/johnfercher/maroto/pkg/v22/image"
	"github.com/johnfercher/maroto/pkg/v22/row"
)

func main() {
	pdf := v22.NewDocument("pdfzin")

	header := buildRow()
	content := buildRow()
	footer := buildRow()

	pdf.Add(header, content, footer)

	pdf.Render()
}

func buildRow() v22.Component {
	row := row.New(20)

	image := image.New("image1")

	col1 := col.New(4)
	col1.Add(image)

	col2 := col.New(4)
	col3 := col.New(4)

	row.Add(col1, col2, col3)
	return row
}
