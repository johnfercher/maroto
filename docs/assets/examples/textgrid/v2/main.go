package main

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"

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

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	m.AddRow(40,
		text.NewCol(2, "Left-aligned text"),
		text.NewCol(4, "Centered text", props.Text{Align: align.Center}),
		text.NewCol(6, "Right-aligned text", props.Text{Align: align.Right}),
	)

	m.AddRows(text.NewRow(10, "Aligned unindented text"))

	m.AddRow(40,
		text.NewCol(2, "Left-aligned text", props.Text{Top: 3, Left: 3, Align: align.Left}),
		text.NewCol(4, "Centered text", props.Text{Top: 3, Align: align.Center}),
		text.NewCol(6, "Right-aligned text", props.Text{Top: 3, Right: 3, Align: align.Right}),
	)

	m.AddRows(text.NewRow(10, "Aligned text with indentation"))

	m.AddRow(40,
		text.NewCol(2, longText, props.Text{Align: align.Left}),
		text.NewCol(4, longText, props.Text{Align: align.Center}),
		text.NewCol(6, longText, props.Text{Align: align.Right}),
	)

	m.AddRows(text.NewRow(10, "Multiline text"))

	m.AddRow(40,
		text.NewCol(2, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: align.Left, BreakLineStrategy: breakline.DashStrategy}),
		text.NewCol(4, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: align.Center}),
		text.NewCol(6, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: align.Right}),
	)

	m.AddRows(text.NewRow(10, "Multiline text with indentation"))

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/textgridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/textgridv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
