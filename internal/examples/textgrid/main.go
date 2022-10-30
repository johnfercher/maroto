package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	longText := "This is a longer sentence that will be broken into multiple lines " +
		"as it does not fit into the column otherwise."

	m.Row(40, func() {
		m.Col(2, func() {
			// Note: Left-alignment is the default for text
			m.Text("Left-aligned text")
		})
		m.Col(4, func() {
			m.Text("Centered text", props.Text{Align: consts.Center})
		})
		m.Col(6, func() {
			m.Text("Right-aligned text", props.Text{Align: consts.Right})
		})
	})

	m.Row(10, func() { m.Text("Aligned unindented text") })

	m.Row(40, func() {
		m.Col(2, func() {
			m.Text("Left-aligned text", props.Text{Top: 3, Left: 3, Align: consts.Left})
		})
		m.Col(4, func() {
			m.Text("Centered text", props.Text{Top: 3, Align: consts.Center})
		})
		m.Col(6, func() {
			m.Text("Right-aligned text", props.Text{Top: 3, Right: 3, Align: consts.Right})
		})
	})

	m.Row(10, func() { m.Text("Aligned text with indentation") })

	m.Row(40, func() {
		m.Col(2, func() {
			m.Text(longText, props.Text{Align: consts.Left})
		})
		m.Col(4, func() {
			m.Text(longText, props.Text{Align: consts.Center})
		})
		m.Col(6, func() {
			m.Text(longText, props.Text{Align: consts.Right})
		})
	})

	m.Row(10, func() { m.Text("Multiline text") })

	m.Row(40, func() {
		m.Col(2, func() {
			m.Text(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Left})
		})
		m.Col(4, func() {
			m.Text(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Center})
		})
		m.Col(6, func() {
			m.Text(longText, props.Text{Top: 3, Left: 3, Right: 3, Align: consts.Right})
		})
	})

	m.Row(10, func() { m.Text("Multiline text with indentation") })

	err := m.OutputFileAndClose("internal/examples/pdfs/textgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
