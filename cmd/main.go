package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	pdf := v2.NewDocument()

	header := buildHeader()
	content := buildContent()
	//footer := buildHeader()
	pdf.Add(header, content /* footer*/)

	err := pdf.Generate("v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildHeader() v2.Row {
	row := row.New(20)

	tx1 := text.New("Lorem ipsum dolor sit amet, consectetur ad")
	col1 := col.New(3)
	col1.Add(tx1)

	tx2 := text.New("tx2")
	col2 := col.New(3)
	col2.Add(tx2)

	tx3 := text.New("tx3")
	col3 := col.New(6)
	col3.Add(tx3)

	row.Add(col1, col2, col3)
	return row
}

func buildContent() v2.Row {
	row := row.New(30)

	tx1 := text.New("txA")
	col1 := col.New(6)
	col1.Add(tx1)

	tx2 := text.New("txB")
	col2 := col.New(6)
	col2.Add(tx2)

	row.Add(col1, col2)
	return row
}
