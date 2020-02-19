package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	//m.SetBorder(true)

	m.Row(40, func() {
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(4, func() {
			m.Text("Gopher International Shipping, Inc.", props.Text{
				Top:         15,
				Size:        20,
				Extrapolate: true,
			})
			m.Text("1000 Shipping Gopher Golang TN 3691234 GopherLand (GL)", props.Text{
				Size: 12,
				Top:  21,
			})
		})
		m.ColSpace(4)
	})

	m.Line(10)

	m.Row(40, func() {
		m.Col(4, func() {
			m.Text("João Sant'Ana 100 Main Street Stringfield TN 39021 United Stats (USA)", props.Text{
				Size: 15,
				Top:  14,
			})
		})
		m.ColSpace(4)
		m.Col(4, func() {
			m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
	})

	m.Line(10)

	m.Row(100, func() {
		m.Col(12, func() {
			_ = m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
				Center:  true,
				Percent: 70,
			})
			m.Text("https://github.com/johnfercher/maroto", props.Text{
				Size:  20,
				Align: consts.Center,
				Top:   80,
			})
		})
	})

	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(6, func() {
			m.Text("CODE: 123412351645231245564 DATE: 20-07-1994 20:20:33", props.Text{
				Size: 15,
				Top:  19,
			})
		})
		m.Col(6, func() {
			m.Text("CA", props.Text{
				Top:   30,
				Size:  85,
				Align: consts.Center,
			})
		})
	})

	m.SetBorder(false)

	err := m.OutputFileAndClose("internal/examples/pdfs/zpl.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
