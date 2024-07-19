package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/signaturegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/signaturegridv2.txt")
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

	m.AddRow(40,
		signature.NewCol(2, "Signature 1"),
		signature.NewCol(4, "Signature 2", props.Signature{FontFamily: fontfamily.Courier}),
		signature.NewCol(6, "Signature 3", props.Signature{FontStyle: fontstyle.BoldItalic}),
	)

	m.AddRow(40,
		signature.NewCol(6, "Signature 4", props.Signature{FontStyle: fontstyle.Italic}),
		signature.NewCol(4, "Signature 5", props.Signature{FontSize: 12}),
		signature.NewCol(2, "Signature 6", props.Signature{FontColor: &props.RedColor}),
	)

	m.AddRow(40,
		signature.NewCol(4, "Signature 7", props.Signature{LineColor: &props.RedColor}),
		signature.NewCol(4, "Signature 8", props.Signature{LineStyle: linestyle.Dashed}),
		signature.NewCol(4, "Signature 9", props.Signature{LineThickness: 0.5}),
	)

	m.AddAutoRow(
		signature.NewCol(4, "Signature 7", props.Signature{LineColor: &props.RedColor}),
		signature.NewCol(4, "Signature 8", props.Signature{LineStyle: linestyle.Dashed}),
		signature.NewCol(4, "Signature 9", props.Signature{LineThickness: 0.5}),
	)

	return m
}
