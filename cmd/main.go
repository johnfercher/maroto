package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code/barcode"
	"github.com/johnfercher/maroto/pkg/v2/code/matrixcode"
	"github.com/johnfercher/maroto/pkg/v2/code/qrcode"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/page"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	pdf := v2.NewDocument()

	header := buildRow()
	content := buildContent()
	p := page.New().Add(
		content, content, content, content, content, content, content, content,
		content, content, content, content, content, content, content, content,
	)
	pdf.Add(header, p /*content  footer*/)

	err := pdf.Generate("v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildRow() v2.Component {
	r := row.New(20)

	//image := image.New("image1")

	col1 := col.New(4)
	col1.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	col2 := col.New(4)
	col2.Add(image.NewFromFile("internal/assets/images/biplane.jpg"))

	col3 := col.New(4)
	col3.Add(text.New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem. "))

	r.Add(col1, col2, col3)
	return r
}

func buildContent() v2.Component {
	row := row.New(20)

	col1 := col.New(4)
	col1.Add(barcode.New("barcode"))

	col2 := col.New(4)
	col2.Add(qrcode.New("qrcode"))

	col3 := col.New(4)
	col3.Add(matrixcode.New("qrcode"))

	row.Add(col1, col2, col3)
	return row
}
