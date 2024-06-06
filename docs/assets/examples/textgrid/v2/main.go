package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
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

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	m.AddRow(40,
		text.NewCol(2, "Red text", props.Text{Color: &props.RedColor}),
		text.NewCol(6, "Green text", props.Text{Color: &props.GreenColor}),
		text.NewCol(4, "Blue text", props.Text{Color: &props.BlueColor}),
	)

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

	google := "https://google.com"

	m.AddRows(text.NewRow(10, "text with hyperlink", props.Text{Hyperlink: &google}))

	subText1 := entity.NewSubText("This is a text", props.SubText{Color: &props.BlueColor})
	subText2 := entity.NewSubText(" with multiple", props.SubText{Size: 7})
	subText3 := entity.NewSubText(" styles", props.SubText{Color: &props.RedColor})

	customText := col.New(12).Add(text.NewCustomText([]*entity.SubText{subText1, subText2, subText3}))
	m.AddRows(row.New(10).Add(customText))
	return m
}
