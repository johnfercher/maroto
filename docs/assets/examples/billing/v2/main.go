package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/maroto"
	"github.com/johnfercher/maroto/v2/maroto/code"
	"github.com/johnfercher/maroto/v2/maroto/color"
	"github.com/johnfercher/maroto/v2/maroto/config"
	"github.com/johnfercher/maroto/v2/maroto/consts"
	"github.com/johnfercher/maroto/v2/maroto/domain"
	"github.com/johnfercher/maroto/v2/maroto/grid/col"
	"github.com/johnfercher/maroto/v2/maroto/grid/row"
	"github.com/johnfercher/maroto/v2/maroto/image"
	"github.com/johnfercher/maroto/v2/maroto/props"
	"github.com/johnfercher/maroto/v2/maroto/text"
)

func main() {
	cfg := config.NewBuilder().
		WithMargins(&config.Margins{Left: 10, Top: 15, Right: 10}).
		Build()

	darkGrayColor := getDarkGrayColor()
	mrt := maroto.NewMaroto(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(getPageFooter())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(text.NewRow(10, "Invoice ABC123456789", props.Text{
		Top:   3,
		Style: consts.Bold,
		Align: consts.Center,
	}))

	m.AddRow(7,
		text.NewCol(3, "Transactions", props.Text{
			Top:   1.5,
			Size:  9,
			Style: consts.Bold,
			Align: consts.Center,
			Color: color.NewWhite(),
		}),
	).WithStyle(&props.Style{BackgroundColor: &darkGrayColor})

	m.AddRows(getTransactions()...)

	m.AddRow(15,
		col.New(6).Add(
			code.NewBar("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			}),
			text.New("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			}),
		),
		col.New(6),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/billingv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	document.GetReport().Print()
}

func getTransactions() []domain.Row {
	rows := []domain.Row{
		row.New(5).Add(
			col.New(3),
			text.NewCol(4, "Product", props.Text{Size: 9, Align: consts.Center, Style: consts.Bold}),
			text.NewCol(2, "Quantity", props.Text{Size: 9, Align: consts.Center, Style: consts.Bold}),
			text.NewCol(3, "Price", props.Text{Size: 9, Align: consts.Center, Style: consts.Bold}),
		),
	}

	var contentsRow []domain.Row
	contents := getContents()
	for i, content := range contents {
		r := row.New(4).Add(
			col.New(3),
			text.NewCol(4, content[1], props.Text{Size: 8, Align: consts.Center}),
			text.NewCol(2, content[2], props.Text{Size: 8, Align: consts.Center}),
			text.NewCol(3, content[3], props.Text{Size: 8, Align: consts.Center}),
		)
		if i%2 == 0 {
			gray := getGrayColor()
			r.WithStyle(&props.Style{BackgroundColor: &gray})
		}

		contentsRow = append(contentsRow, r)
	}

	rows = append(rows, contentsRow...)

	rows = append(rows, row.New(20).Add(
		col.New(7),
		text.NewCol(2, "Total:", props.Text{
			Top:   5,
			Style: consts.Bold,
			Size:  8,
			Align: consts.Right,
		}),
		text.NewCol(3, "R$ 2.567,00", props.Text{
			Top:   5,
			Style: consts.Bold,
			Size:  8,
			Align: consts.Center,
		}),
	))

	return rows
}

func getPageHeader() domain.Row {
	return row.New(20).Add(
		image.NewFromFileCol(3, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(6),
		col.New(3).Add(
			text.New("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
				Size:        8,
				Align:       consts.Right,
				Extrapolate: false,
				Color:       getRedColor(),
			}),
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   12,
				Style: consts.BoldItalic,
				Size:  8,
				Align: consts.Right,
				Color: getBlueColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   15,
				Style: consts.BoldItalic,
				Size:  8,
				Align: consts.Right,
				Color: getBlueColor(),
			}),
		),
	)
}

func getPageFooter() domain.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   13,
				Style: consts.BoldItalic,
				Size:  8,
				Align: consts.Left,
				Color: getBlueColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   16,
				Style: consts.BoldItalic,
				Size:  8,
				Align: consts.Left,
				Color: getBlueColor(),
			}),
		),
	)
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func getContents() [][]string {
	return [][]string{
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
}
