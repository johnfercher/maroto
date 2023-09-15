package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/signature"
	"os"
)

func main() {
	maroto := v2.NewMaroto("internal/examples/pdfs/signaturegridv2.pdf")
	m := v2.NewMarotoMetrified(maroto)

	c1 := col.New(2).Add(signature.New("Signature 1"))
	c2 := col.New(4).Add(signature.New("Signature 2", props.Font{Family: consts.Courier}))
	c3 := col.New(6).Add(signature.New("Signature 3", props.Font{Style: consts.BoldItalic}))
	r1 := row.New(40).Add(c1, c2, c3)

	c5 := col.New(6).Add(signature.New("Signature 4", props.Font{Style: consts.Italic}))
	c6 := col.New(4).Add(signature.New("Signature 5", props.Font{Size: 12}))
	c7 := col.New(2).Add(signature.New("Signature 6", props.Font{Color: color.Color{255, 0, 0}}))
	r2 := row.New(40).Add(c5, c6, c7)

	m.Add(r1, r2)

	report, err := m.GenerateWithReport()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	report.Print()
}
