package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	pdf := v2.NewDocument()

	header := buildRow()
	content := buildContent()
	//footer := buildRow()
	pdf.Add(header, content /* footer*/)

	err := pdf.Generate("v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildRow() v2.Component {
	row := row.New(20)

	image := image.New("image1")

	col1 := col.New(4)
	col1.Add(image)

	col2 := col.New(4)
	col3 := col.New(4)

	row.Add(col1, col2, col3)
	return row
}

func buildContent() v2.Component {
	row := row.New(20)

	tx := text.New("Hello World")

	col1 := col.New(4)
	col1.Add(tx)

	col2 := col.New(4)
	col3 := col.New(4)

	row.Add(col1, col2, col3)
	return row
}
