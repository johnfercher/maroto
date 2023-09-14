package main

import (
	"fmt"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/providers"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/size"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	begin := time.Now()

	provider := providers.NewGofpdf(size.A4)
	m := v2.NewMaroto(provider, "internal/examples/pdfs/imagegridv2.pdf")

	biplane1 := image.NewFromFile("internal/assets/images/biplane.jpg", props.Rect{
		Center:  true,
		Percent: 80,
	})

	c1 := col.New(2).Add(biplane1)
	c2 := col.New(4).Add(biplane1)
	c3 := col.New(6).Add(biplane1)
	r1 := row.New(40).Add(c1, c2, c3)

	m.Add(r)

	m.Row(40, func() {
		m.Col(2, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(6, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
	})

	m.Row(40, func() {
		m.Col(2, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  false,
				Percent: 50,
				Left:    10,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  false,
				Percent: 50,
				Top:     10,
			})
		})
		m.Col(6, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  false,
				Percent: 50,
				Left:    15,
				Top:     15,
			})
		})
	})

	m.Row(40, func() {
		m.Col(8, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  false,
				Percent: 80,
				Top:     5,
				Left:    10,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  false,
				Percent: 80,
				Top:     5,
			})
		})
		m.Col(2, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  false,
				Percent: 80,
				Left:    5,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(2, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
	})

	m.Row(40, func() {
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(8, func() {
			_ = m.FileImage("internal/assets/images/frontpage.png", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/imagegrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
