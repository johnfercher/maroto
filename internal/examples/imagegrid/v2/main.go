package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/providers"
	"github.com/johnfercher/maroto/pkg/v2/size"
	"os"
)

func main() {
	provider := providers.NewGofpdf(size.A4)
	maroto := v2.NewMaroto(provider, "internal/examples/pdfs/imagegridv2.pdf")
	m := v2.NewMarotoMetrified(maroto)

	c1 := col.New(2).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	c2 := col.New(4).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	c3 := col.New(6).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	r1 := row.New(40).Add(c1, c2, c3)
	m.Add(r1)

	c4 := col.New(2).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  false,
		Percent: 50,
		Left:    10,
	}))

	c5 := col.New(4).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  false,
		Percent: 50,
		Top:     10,
	}))

	c6 := col.New(6).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  false,
		Percent: 50,
		Left:    15,
		Top:     15,
	}))

	r2 := row.New(40).Add(c4, c5, c6)
	m.Add(r2)

	c7 := col.New(8).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	c8 := col.New(4).Add(image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	r3 := row.New(40).Add(c7, c8)
	m.Add(r3)

	c9 := col.New(6).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  false,
		Percent: 80,
		Top:     5,
		Left:    10,
	}))

	c10 := col.New(4).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  false,
		Percent: 80,
		Top:     5,
	}))

	c11 := col.New(2).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  false,
		Percent: 80,
		Left:    5,
	}))

	r4 := row.New(40).Add(c9, c10, c11)
	m.Add(r4)

	c12 := col.New(6).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  true,
		Percent: 50,
	}))

	c13 := col.New(4).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  true,
		Percent: 50,
	}))

	c14 := col.New(2).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  true,
		Percent: 50,
	}))

	r5 := row.New(40).Add(c12, c13, c14)
	m.Add(r5)

	c15 := col.New(6).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	c16 := col.New(6).Add(image.NewFromFile("internal/assets/images/frontpage.png", props.Rect{
		Center:  true,
		Percent: 80,
	}))

	r6 := row.New(40).Add(c15, c16)
	m.Add(r6)

	report, err := m.GenerateWithReport()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	report.Print()
}
