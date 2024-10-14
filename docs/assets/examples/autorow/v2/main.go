package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/autorow.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/autorow.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddAutoRow(
		image.NewFromFileCol(5, "docs/assets/images/biplane.jpg"),
		text.NewCol(7, intro),
	)

	m.AddAutoRow(
		image.NewFromFileCol(5, "docs/assets/images/biplane.jpg"),
		text.NewCol(7, intro, props.Text{
			Size: 13,
		}),
	)

	m.AddAutoRow(
		image.NewFromFileCol(5, "docs/assets/images/biplane.jpg"),
		text.NewCol(7, intro, props.Text{
			Size:   13,
			Top:    8,
			Bottom: 9,
		}),
	)

	m.AddAutoRow(
		code.NewBarCol(4, "code"),
		text.NewCol(8, intro),
	)

	m.AddAutoRow(
		code.NewMatrixCol(3, "code"),
		text.NewCol(9, intro),
	)

	m.AddAutoRow(
		code.NewQrCol(2, "code"),
		text.NewCol(10, intro),
	)

	return m
}

var intro = `Numa toca no chão vivia um hobbit. Não uma toca nojenta, suja, úmida, 
cheia de pontas de minhocas e um cheiro de limo, nem tam pouco uma toca seca, vazia, arenosa, 
sem nenhum lugar onde se sentar ou onde comer: era uma toca de hobbit, e isso significa conforto.
Ela tinha uma porta perfeitamente redonda feito uma escotilha, pintada de verde, com uma maçaneta
amarela e brilhante de latão exatamente no meio. A porta se abria para um corredor em forma de tubo,
feito um túnel: um túnel muito confortável, sem fumaça, de paredes com painéis e assoalhos
azulejados e acarpetados, com cadeiras enceradas e montes e montes de cabieiros para chapéus e
casacos - o hobbit apreciava visitas.`
