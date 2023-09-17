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
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(true)
	m.SetPageMargins(20, 20, 20)

	m.Row(40, func() {
		m.Col(4, func() {
			_ = m.FileImage("docs/assets/images/gopherbw.png", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.Text("Margins Test", props.Text{
				Top:         12,
				Size:        12,
				Extrapolate: true,
			})
		})
		m.ColSpace(4)
	})

	err := m.OutputFileAndClose("docs/assets/pdf/margins.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
