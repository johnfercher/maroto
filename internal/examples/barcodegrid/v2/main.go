package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"os"
)

func main() {
	maroto := v2.NewMaroto("internal/examples/pdfs/barcodegridv2.pdf")
	m := v2.NewMarotoMetrified(maroto)

	c1 := col.New(2).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 50,
	}))
	c2 := col.New(4).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 75,
	}))
	c3 := col.New(6).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 100,
	}))

	r1 := row.New(40).Add(c1, c2, c3)

	c4 := col.New(2).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 50,
	}))
	c5 := col.New(4).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 75,
	}))
	c6 := col.New(6).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 100,
	}))

	r2 := row.New(40).Add(c4, c5, c6)

	c7 := col.New(6).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 50,
	}))
	c8 := col.New(4).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 75,
	}))
	c9 := col.New(2).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Percent: 100,
	}))

	r3 := row.New(40).Add(c7, c8, c9)

	c10 := col.New(6).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 50,
	}))
	c11 := col.New(4).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 75,
	}))
	c12 := col.New(2).Add(code.NewBarcode("https://github.com/johnfercher/maroto", props.Barcode{
		Center:  true,
		Percent: 100,
	}))

	r4 := row.New(40).Add(c10, c11, c12)

	m.Add(r1, r2, r3, r4)

	report, err := m.GenerateWithReport()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	report.Print()
}
