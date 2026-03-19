package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
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

	// Bold word in a sentence
	m.AddRows(text.NewRow(10, "Bold word in a sentence", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This sentence has a ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("bold", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk(" word in the middle."),
		),
	)

	// Mixed styles
	m.AddRows(text.NewRow(10, "Mixed styles", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("Normal, ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("bold, ", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk("italic, ", props.Text{Style: fontstyle.Italic}),
			richtext.NewChunk("and bold-italic ", props.Text{Style: fontstyle.BoldItalic}),
			richtext.NewChunk("text in one paragraph."),
		),
	)

	// Colored words
	m.AddRows(text.NewRow(10, "Colored words", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This has ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("red", props.Text{Color: &props.RedColor, Style: fontstyle.Bold}),
			richtext.NewChunk(", "),
			richtext.NewChunk("green", props.Text{Color: &props.GreenColor, Style: fontstyle.Bold}),
			richtext.NewChunk(", and "),
			richtext.NewChunk("blue", props.Text{Color: &props.BlueColor, Style: fontstyle.Bold}),
			richtext.NewChunk(" colored words."),
		),
	)

	// Mixed font sizes with baseline alignment
	m.AddRows(text.NewRow(10, "Mixed font sizes (baseline aligned)", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("Small ", props.Text{Size: 8, Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("Medium ", props.Text{Size: 12}),
			richtext.NewChunk("Large ", props.Text{Size: 16, Style: fontstyle.Bold}),
			richtext.NewChunk("back to small.", props.Text{Size: 8}),
		),
	)

	// Word wrapping with mixed styles
	m.AddRows(text.NewRow(10, "Word wrapping with mixed styles", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This is a longer paragraph that demonstrates how ", props.Text{Top: 2, Left: 2, Right: 2, Bottom: 2}),
			richtext.NewChunk("rich text", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk(" handles "),
			richtext.NewChunk("word wrapping", props.Text{Style: fontstyle.Italic}),
			richtext.NewChunk(" across multiple lines. When the text exceeds the available column width, it automatically breaks to the next line while "),
			richtext.NewChunk("preserving the style", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
			richtext.NewChunk(" of each individual chunk."),
		),
	)

	// Multiple columns
	m.AddRows(text.NewRow(10, "Multiple columns", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(4,
			richtext.NewChunk("Price: ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("$99.99", props.Text{Style: fontstyle.Bold, Color: &props.RedColor}),
		),
		richtext.NewCol(4,
			richtext.NewChunk("Status: ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("Active", props.Text{Style: fontstyle.Bold, Color: &props.GreenColor}),
		),
		richtext.NewCol(4,
			richtext.NewChunk("Category: ", props.Text{Top: 2, Left: 2, Bottom: 2}),
			richtext.NewChunk("Electronics", props.Text{Style: fontstyle.Italic}),
		),
	)

	// With margins and vertical padding
	m.AddRows(text.NewRow(10, "With margins and vertical padding", props.Text{Top: 3, Left: 2, Bottom: 2, Style: fontstyle.Bold, Size: 9}))
	m.AddAutoRow(
		richtext.NewCol(12,
			richtext.NewChunk("This paragraph has top margin, left/right margins, and vertical padding between lines. ", props.Text{
				Top:             4,
				Bottom:          4,
				Left:            5,
				Right:           5,
				VerticalPadding: 3,
			}),
			richtext.NewChunk("The bold text here", props.Text{Style: fontstyle.Bold}),
			richtext.NewChunk(" continues flowing with the same layout settings applied from the first chunk."),
		),
	)

	return m
}
