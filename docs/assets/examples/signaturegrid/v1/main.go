package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(2, func() {
			m.Signature("Signature 1")
		})
		m.Col(4, func() {
			m.Signature("Signature 2", props.Font{Family: consts.Courier})
		})
		m.Col(6, func() {
			m.Signature("Signature 3", props.Font{Style: consts.BoldItalic})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			m.Signature("Signature 4", props.Font{Style: consts.Italic})
		})
		m.Col(4, func() {
			m.Signature("Signature 5", props.Font{Size: 12})
		})
		m.Col(2, func() {
			m.Signature("Signature 6", props.Font{Color: color.Color{255, 0, 0}})
		})
	})

	err := m.OutputFileAndClose("docs/assets/pdf/signaturegrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
