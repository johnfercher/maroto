package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	backgroundImage := "docs/assets/images/certificate.png"
	m := GetMaroto(backgroundImage)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/backgroundv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/backgroundv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(image string) core.Maroto {
	bytes, err := os.ReadFile(image)
	if err != nil {
		log.Fatal(err)
	}
	b := config.NewBuilder().
		WithMargins(0, 0, 0).
		WithOrientation(orientation.Horizontal).
		WithMaxGridSize(20).
		WithBackgroundImage(bytes, extension.Png)

	mrt := maroto.New(b.Build())
	m := maroto.NewMetricsDecorator(mrt)

	m.AddPages(AddPage(), AddPage(), AddPage(), AddPage(), AddPage())
	return m
}

func AddPage() core.Page {
	return page.New().Add(
		row.New(70),
		row.New(20).Add(
			col.New(4),
			text.NewCol(12, "O GDG-Petrópolis certifica que Fulano de Tal 123 participou do Evento Exemplo 123 no dia 2019-03-30.", props.Text{
				Size: 18,
			}),
			col.New(4),
		),
		row.New(15),
		row.New(30).Add(
			image.NewFromFileCol(20, "docs/assets/images/signature.png", props.Rect{
				Center: true,
			}),
		),
	)
}
