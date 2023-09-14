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
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

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
