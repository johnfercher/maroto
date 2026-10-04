package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const paragraph = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse et ex quam. " +
	"Proin aliquam nibh id arcu finibus, sit amet ultrices velit vulputate. Curabitur hendrerit et " +
	"metus a accumsan. Mauris quis purus rhoncus diam tempor. Suspendisse potenti."

var (
	headerBackground = &props.Color{Red: 55, Green: 55, Blue: 55}
	stripeBackground = &props.Color{Red: 235, Green: 235, Blue: 235}
	cellBorder       = &props.Cell{BorderType: border.Full, BorderColor: &props.Color{Red: 200, Green: 200, Blue: 200}}

	headerText = props.Text{Style: fontstyle.Bold, Size: 9, Color: &props.WhiteColor, Top: 1.5, Left: 1.5, Right: 1.5}
	cellText   = props.Text{Size: 8, Top: 1.5, Left: 1.5, Right: 1.5}
	bodyText   = props.Text{Size: 9, Bottom: 3, Align: align.Justify}
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/tablev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/tablev2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(text.NewRow(12, "Invoice", props.Text{Size: 16, Style: fontstyle.Bold}))
	for range 6 {
		m.AddRows(text.NewAutoRow(paragraph, bodyText))
	}

	m.AddTable([]core.Row{getHeader()}, getItems(45)...)

	// The table header is not repeated for the rows after the table.
	m.AddRows(row.New(8), text.NewRow(10, "Terms and conditions", props.Text{Size: 12, Style: fontstyle.Bold}))
	for range 24 {
		m.AddRows(text.NewAutoRow(paragraph, bodyText))
	}

	return m
}

func getHeader() core.Row {
	return row.New(7).Add(
		text.NewCol(2, "Code", headerText),
		text.NewCol(6, "Description", headerText),
		text.NewCol(2, "Qty", withAlign(headerText, align.Right)),
		text.NewCol(2, "Total", withAlign(headerText, align.Right)),
	).WithStyle(&props.Cell{BackgroundColor: headerBackground})
}

func getItems(n int) []core.Row {
	rows := make([]core.Row, 0, n)
	for i := range n {
		r := row.New(6).Add(
			text.NewCol(2, fmt.Sprintf("SKU-%03d", i+1), cellText).WithStyle(cellBorder),
			text.NewCol(6, fmt.Sprintf("Item %d", i+1), cellText).WithStyle(cellBorder),
			text.NewCol(2, fmt.Sprintf("%d", i%5+1), withAlign(cellText, align.Right)).WithStyle(cellBorder),
			text.NewCol(2, fmt.Sprintf("%.2f", float64(i%5+1)*9.5), withAlign(cellText, align.Right)).WithStyle(cellBorder),
		)
		if i%2 == 1 {
			r.WithStyle(&props.Cell{BackgroundColor: stripeBackground})
		}
		rows = append(rows, r)
	}

	return rows
}

func withAlign(base props.Text, a align.Type) props.Text {
	base.Align = a
	return base
}
