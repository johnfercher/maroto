package main

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageTopMargin(0)
	m.SetPageLeftMargin(0)
	m.SetPageRightMargin(0)

	m.Row(12, func() {
		m.SetBackgroundColor(color.Color{})
		m.Col(12, func() {
			m.Text("Hello World", props.Text{Color: color.NewWhite()})
		})
	})

	m.Row(12, func() {
		m.SetBackgroundColor(color.Color{})
		m.SetPageRightMargin(50)
		m.Col(12, func() {
			m.Text("Hello World", props.Text{Color: color.NewWhite()})
		})
		m.SetPageRightMargin(0)
	})

	m.Row(12, func() {
		m.SetBackgroundColor(color.Color{})
		m.SetPageLeftMargin(50)
		m.Col(12, func() {
			m.Text("Hello World", props.Text{Color: color.NewWhite()})
		})
		m.SetPageLeftMargin(0)
	})

	m.Row(12, func() {
		m.SetBackgroundColor(color.Color{})
		m.SetPageMargins(0, 0, 0)
		m.Col(12, func() {
			m.Text("Hello World", props.Text{Color: color.NewWhite()})
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/usingmargins.pdf")
	assert.Nil(nil, err)
}
