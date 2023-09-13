package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	pdf := v2.NewDocument("v2.pdf")

	header := buildRow()
	content := buildContent()

	pdf.Add(header /*content  footer*/)
	pdf.Add(
		content, content, content, content, content, content, content, content,
		content, content, content, content, content, content, content, content,
		content, content, content, content, content, content, content, content,
		content, content, content, content, content, content, content, content,
	)

	//pdf.Add()
	//pdf.ForceAddPage(p)

	err := pdf.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildRow() domain.Row {
	r := row.New(70)

	//image := image.New("image1")

	col1 := col.New(4)
	col1.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	col2 := col.New(4)
	col2.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	col3 := col.New(4)
	col3.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	r.Add(col1, col2, col3)
	return r
}

func buildContent() domain.Row {
	row := row.New(70)

	col1 := col.New(4)
	col1.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	col2 := col.New(4)
	col2.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	col3 := col.New(4)
	col3.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	row.Add(col1, col2, col3)
	return row
}
