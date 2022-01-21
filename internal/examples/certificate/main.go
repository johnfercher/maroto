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
	m := pdf.NewMaroto(consts.Landscape, consts.A4)
	// m.SetBorder(true)

	m.Row(20, func() {
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Percent: 88,
				Center:  true,
			})
		})
		m.Col(4, func() {
			m.Text("Golang Certificate", props.Text{
				Top:   6,
				Align: consts.Center,
				Size:  20,
				Style: consts.BoldItalic,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Percent: 90,
				Center:  true,
			})
		})
	})

	m.Row(130, func() {
		m.Col(12, func() {
			text := "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
			m.Text(text, props.Text{
				Size:            13,
				Align:           consts.Center,
				Top:             50,
				VerticalPadding: 2.0,
			})
		})
	})

	m.Row(25, func() {
		m.Col(4, func() {
			m.Signature("Gopher Senior")
		})
		m.Col(4, func() {
			m.Signature("Gopheroid")
		})
		m.Col(4, func() {
			m.Signature("Sign Here")
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/certificate.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
