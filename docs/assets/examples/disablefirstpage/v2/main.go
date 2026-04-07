package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	img := "docs/assets/images/certificate.png"
	m := GetMaroto(img)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/disablefirstpagev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/disablefirstpagev2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(source string) core.Maroto {
	bytes, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	b := config.NewBuilder().
		WithDisableFirstPage(true).
		Build()

	mrt := maroto.New(b)
	m := maroto.NewMetricsDecorator(mrt)

	img := image.NewAutoFromBytesRow(bytes, extension.Png)
	m.AddAutoRow(img.GetColumns()...)

	return m
}
