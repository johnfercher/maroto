package main

import "github.com/johnfercher/maroto"

func main() {
	m := maroto.NewMaroto(maroto.Landscape, maroto.A4)
	//m.SetDebugMode(true)

	m.Row(20, func() {
		m.Col(func() {
			m.FileImage("examples/internal/assets/images/frontpage.png")
		})
		m.Col(func() {
			m.Text("Golang Certificate", maroto.TextProp{
				Top:   12,
				Align: maroto.Center,
				Size:  20,
				Style: maroto.BoldItalic,
			})
		})
		m.Col(func() {
			m.FileImage("examples/internal/assets/images/frontpage.png")
		})
	})

	m.Row(130, func() {
		m.Col(func() {
			text := "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
			m.Text(text, maroto.TextProp{
				Size:  13,
				Align: maroto.Center,
				Top:   60,
			})
		})
	})

	m.Row(25, func() {
		m.Col(func() {
			m.Signature("Gopher Senior")
		})
		m.Col(func() {
			m.Signature("Gopheroid")
		})
		m.Col(func() {
			m.Signature("Sign Here")
		})
	})

	m.OutputFileAndClose("examples/internal/pdfs/certificate.pdf")
}
