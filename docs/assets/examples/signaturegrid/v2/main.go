package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/core/color"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	m.AddRow(40,
		signature.NewCol(2, "Signature 1"),
		signature.NewCol(4, "Signature 2", props.Font{Family: fontfamily.Courier}),
		signature.NewCol(6, "Signature 3", props.Font{Style: fontstyle.BoldItalic}),
	)

	m.AddRow(40,
		signature.NewCol(6, "Signature 4", props.Font{Style: fontstyle.Italic}),
		signature.NewCol(4, "Signature 5", props.Font{Size: 12}),
		signature.NewCol(2, "Signature 6", props.Font{Color: &color.Color{255, 0, 0}}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/signaturegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
