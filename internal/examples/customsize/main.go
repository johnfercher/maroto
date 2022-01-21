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
	m := pdf.NewMarotoCustomSize(consts.Landscape, "C6", "mm", 114.0, 162.0)
	m.SetPageMargins(5, 5, 5)
	// m.SetBorder(true)

	m.Row(40, func() {
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.Text("Gopher International Shipping, Inc.", props.Text{
				Top:         12,
				Size:        12,
				Extrapolate: true,
			})
		})
		m.ColSpace(4)
	})

	m.Line(10)

	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("João Sant'Ana 100 Main Street", props.Text{
				Size:  10,
				Align: consts.Right,
			})
			m.Text("Springfield TN 39021", props.Text{
				Size:  10,
				Align: consts.Right,
				Top:   10,
			})
			m.Text("United States (USA)", props.Text{
				Size:  10,
				Align: consts.Right,
				Top:   20,
			})
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/customsize.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
