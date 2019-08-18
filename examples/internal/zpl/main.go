package main

import "github.com/johnfercher/maroto"

func main() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.Letter)
	//m.SetDebugMode(true)

	m.Row(40, func() {
		m.Col(func() {
			m.FileImage("examples/internal/assets/images/biplane.jpg", maroto.RectProp{
				Left:    20,
				Center:  false,
				Percent: 80,
			})
		})
		m.Col(func() {
			m.Text("Gopher International Shipping, Inc.", maroto.TextProp{
				Top:         15,
				Size:        20,
				Extrapolate: true,
			})
			m.Text("1000 Shipping Gopher Golang TN 3691234 GopherLand (GL)", maroto.TextProp{
				Size: 12,
				Top:  21,
			})
		})
		m.ColSpace()
	})

	m.Line(10)

	m.Row(40, func() {
		m.Col(func() {
			m.Text("João Sant'Ana 100 Main Street Stringfield TN 39021 United Stats (USA)", maroto.TextProp{
				Size: 15,
				Top:  14,
			})
		})
		m.ColSpace()
		m.Col(func() {
			m.QrCode("https://github.com/johnfercher/maroto", maroto.RectProp{
				Percent: 75,
			})
		})
	})

	m.Line(10)

	m.Row(100, func() {
		m.Col(func() {
			m.Barcode("https://github.com/johnfercher/maroto", maroto.BarcodeProp{
				Percent: 70,
			})
			m.Text("https://github.com/johnfercher/maroto", maroto.TextProp{
				Size:  20,
				Align: maroto.Center,
				Top:   80,
			})
		})
	})

	m.SetDebugMode(true)

	m.Row(40, func() {
		m.Col(func() {
			m.Text("CODE: 123412351645231245564 DATE: 20-07-1994 20:20:33", maroto.TextProp{
				Size: 15,
				Top:  19,
			})
		})
		m.Col(func() {
			m.Text("CA", maroto.TextProp{
				Top:   30,
				Size:  85,
				Align: maroto.Center,
			})
		})
	})

	m.SetDebugMode(false)

	m.OutputFileAndClose("examples/internal/pdfs/zpl.pdf")
}
