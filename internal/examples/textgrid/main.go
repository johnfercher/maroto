package main

import (
	"fmt"
	"github.com/Vale-sail/maroto/pkg/consts"
	"github.com/Vale-sail/maroto/pkg/pdf"
	"os"
	"time"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(2, func() {
			m.Text("Any Text1")
		})
		m.Col(4, func() {
			m.Text("Any Text2")
		})
		m.Col(6, func() {
			m.Text("Any Text3")
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/textgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
