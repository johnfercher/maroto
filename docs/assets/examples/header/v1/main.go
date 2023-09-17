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

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Header", props.Text{
					Size:  10,
					Style: consts.Bold,
					Align: consts.Center,
				})
			})
		})
	})

	for i := 0; i < 50; i++ {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("Dummy text", props.Text{
					Size: 8,
				})
			})
		})
	}

	err := m.OutputFileAndClose("docs/assets/pdf/header.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
