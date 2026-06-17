package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/richtextgridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/richtextgridv2.txt")
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

	m.AddRows(text.NewRow(10, "Bold word in a sentence", props.Text{
		Top:    3,
		Left:   2,
		Bottom: 2,
		Style:  fontstyle.Bold,
		Size:   9,
	}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This sentence has a ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("bold", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk(" word in the middle."),
		),
	)

	m.AddRows(text.NewRow(10, "Mixed styles and colors", props.Text{
		Top:    3,
		Left:   2,
		Bottom: 2,
		Style:  fontstyle.Bold,
		Size:   9,
	}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("Normal, ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("bold, ", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk("italic, ", props.Text{Style: fontstyle.Italic}),
			richtext.NewChunk("red", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
			richtext.NewChunk(", "),
			richtext.NewChunk("green", props.Text{Style: fontstyle.Bold, Color: &props.GreenColor}),
			richtext.NewChunk(", and "),
			richtext.NewChunk("blue", props.Text{Style: fontstyle.Bold, Color: &props.BlueColor}),
			richtext.NewChunk(" in one flowing paragraph."),
		),
	)

	m.AddRows(text.NewRow(10, "Mixed font sizes", props.Text{
		Top:    3,
		Left:   2,
		Bottom: 2,
		Style:  fontstyle.Bold,
		Size:   9,
	}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("Small ", props.Text{Top: 2, Left: 2, Bottom: 2, Size: 8}),
			richtext.NewChunk("Medium ", props.Text{Size: 12}),
			richtext.NewChunk("Large ", props.Text{Size: 16, Style: fontstyle.Bold}),
			richtext.NewChunk("back to small.", props.Text{Size: 8}),
		),
	)

	m.AddRows(text.NewRow(10, "Word wrapping across styles", props.Text{
		Top:    3,
		Left:   2,
		Bottom: 2,
		Style:  fontstyle.Bold,
		Size:   9,
	}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This is a longer paragraph that demonstrates how ", props.Text{
				Top:   2,
				Left:  2,
				Right: 2,
			}),
			richtext.NewChunk("rich text", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk(" handles "),
			richtext.NewChunk("word wrapping", props.Text{Style: fontstyle.Italic}),
			richtext.NewChunk(" across multiple lines while "),
			richtext.NewChunk("preserving each chunk style", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
			richtext.NewChunk(" inside the same paragraph."),
		),
	)

	m.AddRows(text.NewRow(10, "Alignment and line breaks", props.Text{
		Top:    3,
		Left:   2,
		Bottom: 2,
		Style:  fontstyle.Bold,
		Size:   9,
	}))
	m.AddAutoRow(
		richtext.NewCol(6,
			richtext.NewChunk("Centered rich text\nwith explicit line breaks", props.Text{
				Top:                2,
				Left:               2,
				Right:              2,
				Align:              align.Center,
				PreserveLineBreaks: true,
			}),
			richtext.NewChunk(" and "),
			richtext.NewChunk("bold emphasis", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk("."),
		),
		richtext.NewCol(6,
			richtext.NewChunk("Justified text keeps the paragraph flowing while distributing the available width between words.", props.Text{
				Top:   2,
				Left:  2,
				Right: 2,
				Align: align.Justify,
			}),
		),
	)

	return m
}
