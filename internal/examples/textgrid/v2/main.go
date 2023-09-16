package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	maroto := v2.NewMaroto(cfg)
	m := v2.NewMetricsDecorator(maroto)

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				text.New("Left-aligned text"),
			),
			col.New(4).Add(
				text.New("Centered text", props.Text{Align: consts.Center}),
			),
			col.New(6).Add(
				text.New("Right-aligned text", props.Text{Align: consts.Right}),
			),
		),
	)

	m.Add(
		row.New(10).Add(
			col.New(12).Add(
				text.New("Aligned unindented text"),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				text.New("Left-aligned text", props.Text{Top: 3, Left: 3, Align: consts.Left}),
			),
			col.New(4).Add(
				text.New("Centered text", props.Text{Top: 3, Align: consts.Center}),
			),
			col.New(6).Add(
				text.New("Right-aligned text", props.Text{Top: 3, Right: 3, Align: consts.Right}),
			),
		),
	)

	m.Add(
		row.New(10).Add(
			col.New(12).Add(
				text.New("Aligned text with indentation"),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				text.New(longText, props.Text{Align: consts.Left}),
			),
			col.New(4).Add(
				text.New(longText, props.Text{Align: consts.Center}),
			),
			col.New(6).Add(
				text.New(longText, props.Text{Align: consts.Right}),
			),
		),
	)

	m.Add(
		row.New(10).Add(
			col.New(12).Add(
				text.New("Multiline text"),
			),
		),
	)

	m.Add(
		row.New(40).Add(
			col.New(2).Add(
				text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Left}),
			),
			col.New(4).Add(
				text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Center}),
			),
			col.New(6).Add(
				text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Right}),
			),
		),
	)

	m.Add(
		row.New(10).Add(
			col.New(12).Add(
				text.New("Multiline text with indentation"),
			),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("internal/examples/pdfs/textgridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}
