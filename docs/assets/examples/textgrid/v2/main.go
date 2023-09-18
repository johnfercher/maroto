package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/consts"
	"github.com/johnfercher/v2/maroto/props"
	"github.com/johnfercher/v2/maroto/text"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	m.AddRow(40,
		text.NewCol(2, "Left-aligned text"),
		text.NewCol(4, "Centered text", props.Text{Align: consts.Center}),
		text.NewCol(6, "Right-aligned text", props.Text{Align: consts.Right}),
	)

	m.AddRows(text.NewRow(10, "Aligned unindented text"))

	m.AddRow(40,
		text.NewCol(2, "Left-aligned text", props.Text{Top: 3, Left: 3, Align: consts.Left}),
		text.NewCol(4, "Centered text", props.Text{Top: 3, Align: consts.Center}),
		text.NewCol(6, "Right-aligned text", props.Text{Top: 3, Right: 3, Align: consts.Right}),
	)

	m.AddRows(text.NewRow(10, "Aligned text with indentation"))

	m.AddRow(40,
		text.NewCol(2, longText, props.Text{Align: consts.Left}),
		text.NewCol(4, longText, props.Text{Align: consts.Center}),
		text.NewCol(6, longText, props.Text{Align: consts.Right}),
	)

	m.AddRows(text.NewRow(10, "Multiline text"))

	m.AddRow(40,
		text.NewCol(2, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Left}),
		text.NewCol(4, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Center}),
		text.NewCol(6, longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Right}),
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

	document.GetReport().Print()
}
