package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"os"
)

func main() {
	maroto := v2.NewMaroto("internal/examples/pdfs/textgridv2.pdf")
	m := v2.NewMarotoMetrified(maroto)

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	c1 := col.New(2).Add(text.New("Left-aligned text"))
	c2 := col.New(4).Add(text.New("Centered text", props.Text{Align: consts.Center}))
	c3 := col.New(6).Add(text.New("Right-aligned text", props.Text{Align: consts.Right}))
	r1 := row.New(40).Add(c1, c2, c3)

	c4 := col.New(12).Add(text.New("Aligned unindented text"))
	r2 := row.New(10).Add(c4)

	c5 := col.New(2).Add(text.New("Left-aligned text", props.Text{Top: 3, Left: 3, Align: consts.Left}))
	c6 := col.New(4).Add(text.New("Centered text", props.Text{Top: 3, Align: consts.Center}))
	c7 := col.New(6).Add(text.New("Right-aligned text", props.Text{Top: 3, Right: 3, Align: consts.Right}))
	r3 := row.New(40).Add(c5, c6, c7)

	c8 := col.New(12).Add(text.New("Aligned text with indentation"))
	r4 := row.New(10).Add(c8)

	c9 := col.New(2).Add(text.New(longText, props.Text{Align: consts.Left}))
	c10 := col.New(4).Add(text.New(longText, props.Text{Align: consts.Center}))
	c11 := col.New(6).Add(text.New(longText, props.Text{Align: consts.Right}))
	r5 := row.New(40).Add(c9, c10, c11)

	c12 := col.New(12).Add(text.New("Multiline text"))
	r6 := row.New(10).Add(c12)

	c13 := col.New(2).Add(text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Left}))
	c14 := col.New(4).Add(text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Center}))
	c15 := col.New(6).Add(text.New(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Right}))
	r7 := row.New(40).Add(c13, c14, c15)

	c16 := col.New(12).Add(text.New("Multiline text with indentation"))
	r8 := row.New(10).Add(c16)

	m.Add(r1, r2, r3, r4, r5, r6, r7, r8)

	report, err := m.GenerateWithReport()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	report.Print()
}
