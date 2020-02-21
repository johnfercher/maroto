package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

func main() {
	begin := time.Now()

	darkGrayColor := color.Color{
		Red:   144,
		Green: 144,
		Blue:  144,
	}

	grayColor := color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}

	whiteColor := color.NewWhite()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(func() {
				_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				})
			})
			m.ColSpaces(2)
			m.Col(func() {
				m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
					Top:         4,
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
				})
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   19,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
			})
		})
		m.Row(5, func() {
			m.ColSpace()
		})
	})

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(func() {
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Col(func() {
			m.Text("Invoice ABC123456789", props.Text{
				Top:   6,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(darkGrayColor)
	m.Row(7, func() {
		m.Col(func() {
			m.Text("Transactions", props.Text{
				Top:   4.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.ColSpaces(3)
	})
	m.SetBackgroundColor(whiteColor)

	header := []string{"", "Product", "Quantity", "Price"}
	contents := [][]string{
		{"", "Swamp", "12", "R$ 4,00"},
		{"", "Sorin, A Planeswalker", "4", "R$ 90,00"},
		{"", "Tassa", "4", "R$ 30,00"},
		{"", "Skinrender", "4", "R$ 9,00"},
		{"", "Island", "12", "R$ 4,00"},
		{"", "Mountain", "12", "R$ 4,00"},
		{"", "Plain", "12", "R$ 4,00"},
		{"", "Black Lotus", "1", "R$ 1.000,00"},
		{"", "Time Walk", "1", "R$ 1.000,00"},
		{"", "Emberclave", "4", "R$ 44,00"},
		{"", "Anax", "4", "R$ 32,00"},
		{"", "Murderous Rider", "4", "R$ 22,00"},
		{"", "Gray Merchant of Asphodel", "4", "R$ 2,00"},
		{"", "Ajani's Pridemate", "4", "R$ 2,00"},
		{"", "Renan, Chatuba", "4", "R$ 19,00"},
		{"", "Tymarett", "4", "R$ 13,00"},
		{"", "Doom Blade", "4", "R$ 5,00"},
		{"", "Dark Lord", "3", "R$ 7,00"},
		{"", "Memory of Thanatos", "3", "R$ 32,00"},
		{"", "Poring", "4", "R$ 1,00"},
		{"", "Deviling", "4", "R$ 99,00"},
		{"", "Seiya", "4", "R$ 45,00"},
		{"", "Harry Potter", "4", "R$ 62,00"},
		{"", "Goku", "4", "R$ 77,00"},
		{"", "Phreoni", "4", "R$ 22,00"},
		{"", "Katheryn High Wizard", "4", "R$ 25,00"},
		{"", "Lord Seyren", "4", "R$ 55,00"},
	}

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.Font{
			Size: 9,
		},
		ContentProp: props.Font{
			Size: 8,
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpaces(2)
		m.Col(func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(15, func() {
		m.Col(func() {
			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			})
			m.Text("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			})
		})
		m.ColSpace()
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
