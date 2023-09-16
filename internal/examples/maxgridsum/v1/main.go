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
	gridSum := 14.0
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetMaxGridSum(gridSum)

	m.Row(10, func() {
		m.Col(uint(gridSum), func() {
			m.Text(fmt.Sprintf("Table with %v Columns", gridSum), props.Text{Style: consts.Bold})
		})
	})

	m.SetBorder(true)

	m.Row(8, func() {
		for i := 1; i <= int(gridSum); i++ {
			m.Col(1, func() {
				m.Text(fmt.Sprintf("H %d", i), props.Text{Style: consts.Bold, Top: 1.5, Left: 1.5})
			})
		}
	})
	m.Row(6, func() {
		for i := 1; i <= int(gridSum); i++ {
			m.Col(1, func() {
				m.Text(fmt.Sprintf("C %d", i), props.Text{Top: 1, Left: 1.5, Size: 9})
			})
		}
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/maxgridsum.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
