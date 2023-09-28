package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/documenttype"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(101.6, 152.4).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	bytes, err := os.ReadFile("docs/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	m.AddRow(20,
		image.NewFromBytesCol(4, bytes, extension.Png, props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(8).Add(
			text.New("Maroto V2", props.Text{
				Size:  15,
				Style: fontstyle.Bold,
			}),
			text.New("ZPL Generator", props.Text{
				Top:   8,
				Size:  10,
				Style: fontstyle.Bold,
			}),
			text.New("Based On PDF", props.Text{
				Top:   12,
				Size:  10,
				Style: fontstyle.Bold,
			}),
		),
	)

	m.AddRows(line.NewRow(5, props.Line{OffsetPercent: 50}))

	m.AddRow(20,
		col.New(8).Add(
			text.New("Maroto V2", props.Text{
				Left:  10,
				Top:   3,
				Align: align.Left,
			}),
			text.New("ZPL Generator", props.Text{
				Left:  10,
				Top:   8,
				Align: align.Left,
			}),
			text.New("Based On PDF", props.Text{
				Left:  10,
				Top:   13,
				Align: align.Left,
			}),
		),
		code.NewQrCol(4, "qrcode", props.Rect{
			Center: true,
		}),
	)

	m.AddRows(line.NewRow(5, props.Line{OffsetPercent: 50}))

	m.AddRows(code.NewBarRow(20, "code", props.Barcode{
		Percent: 70,
		Center:  true,
	}))

	m.AddRows(text.NewRow(20, "code123"))

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/zplv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document, err = document.To(documenttype.ZPL)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/zpl/zplv2.zpl")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/zplv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
