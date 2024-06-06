package text_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how to create a text components with different styles.
func ExampleNew() {
	m := maroto.New()

	text := text.New("text")
	col := col.New(12).Add(text)
	m.AddRow(10, col)

	// generate document
}

// ExampleNew demonstrates how to create a text component with
func ExampleNewCustomText() {
	m := maroto.New()

	subText1 := entity.NewSubText("This is a text", props.SubText{Color: &props.BlueColor})
	subText2 := entity.NewSubText(" with multiple", props.SubText{Size: 7})
	subText3 := entity.NewSubText(" styles", props.SubText{Color: &props.RedColor})

	customText := col.New(12).Add(text.NewCustomText([]*entity.SubText{subText1, subText2, subText3}))
	m.AddRows(row.New(10).Add(customText))

	// generate document
}

// ExampleNewCol demonstrates how to create a text component wrapped into a column.
func ExampleNewCol() {
	m := maroto.New()

	textCol := text.NewCol(12, "text")
	m.AddRow(10, textCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a text component wrapped into a row.
func ExampleNewRow() {
	m := maroto.New()

	textRow := text.NewRow(10, "text")
	m.AddRows(textRow)

	// generate document
}
