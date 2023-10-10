package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save("docs/assets/pdf/simplestv2.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaroto() core.Maroto {
	m := maroto.New()

	m.AddRow(20,
		code.NewBarCol(4, "barcode"),
		code.NewMatrixCol(4, "matrixcode"),
		code.NewQrCol(4, "qrcode"),
	)

	m.AddRow(10, col.New(12))

	m.AddRow(20,
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg"),
		signature.NewCol(4, "signature"),
		text.NewCol(4, "text"),
	)

	m.AddRow(10, col.New(12))

	m.AddRow(20, line.NewCol(12))

	return m
}
