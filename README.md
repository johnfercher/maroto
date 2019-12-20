# Maroto 

[![GoDoc](https://godoc.org/github.com/johnfercher/maroto?status.svg)](https://godoc.org/github.com/johnfercher/maroto)
[![Travis](https://travis-ci.com/johnfercher/maroto.svg?branch=master)][travis] 
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto) 
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  

A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

You can write your PDFs like you are creating a site using Bootstrap. A Row may have many Cols, and a Col may have many components. 
Besides that, pages will be added when content may extrapolate the useful area. You can define a header which will be added
always when a new page appear, in this case, a header may have many rows, lines or tablelist. 

## Installation

* With `go get`:

```bash
go get -u github.com/johnfercher/maroto
```

## Features

![result](internal/assets/images/diagram.png)

#### Grid System
* [Row](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Row)
* [Col](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Col)
* [ColSpace](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-ColSpace)
* [ColSpaces](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-ColSpaces)

#### Components To Use Inside a Col
* [Text w/ automatic new lines](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Text)
* [Signature](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Signature)
* Image ([From file](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-FileImage) or [Base64](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Base64Image))
* [QrCode](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.QrCode)
* [Barcode](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.Barcode)   
    
#### Components To Use Outside a Row
* [TableList](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-TableList)
* [Line](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Line)
    
#### Components To Wrap Row, TableList and Line
* [RegisterHeader](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.RegisterHeader)
* [RegisterFooter](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.RegisterFooter)

#### Others   
* [Properties](https://godoc.org/github.com/johnfercher/maroto/pkg/props): most of the components has properties which you can use to customize appearance and behavior.
* [SetBorder](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.SetDebugMode): Used to draw rectangles in every row and column
* Automatic New Page: New pages are generated automatically when needed.
* 100% Unicode
* Save: You can [save on disk](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-OutputFileAndClose) or export to a [base64 string](https://godoc.org/github.com/johnfercher/maroto/pkg/pdf#example-PdfMaroto-Output)

#### Roadmap
* Updated in [Issues](https://github.com/johnfercher/maroto/issues)

## Examples
In the [PDFs](internal/examples/pdfs) folder there are the PDFs generated
using Maroto, and in the [examples](internal/examples) folder there are subfolders
with the code to generate the PDFs.

![result](internal/assets/images/result.png)

#### Code
```go
package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	m.Row(40, func() {
		m.Col(func() {
			_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.Col(func() {
			m.Text("Gopher International Shipping, Inc.", props.Text{
				Top:         15,
				Size:        20,
				Extrapolate: true,
			})
			m.Text("1000 Shipping Gopher Golang TN 3691234 GopherLand (GL)", props.Text{
				Size: 12,
				Top:  21,
			})
		})
		m.ColSpace()
	})

	m.Line(10)

	m.Row(40, func() {
		m.Col(func() {
			m.Text("João Sant'Ana 100 Main Street Stringfield TN 39021 United Stats (USA)", props.Text{
				Size: 15,
				Top:  14,
			})
		})
		m.ColSpace()
		m.Col(func() {
			m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
	})

	m.Line(10)

	m.Row(100, func() {
		m.Col(func() {
			_ = m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
				Center:  true,
				Percent: 70,
			})
			m.Text("https://github.com/johnfercher/maroto", props.Text{
				Size:  20,
				Align: consts.Center,
				Top:   80,
			})
		})
	})

	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(func() {
			m.Text("CODE: 123412351645231245564 DATE: 20-07-1994 20:20:33", props.Text{
				Size: 15,
				Top:  19,
			})
		})
		m.Col(func() {
			m.Text("CA", props.Text{
				Top:   30,
				Size:  85,
				Align: consts.Center,
			})
		})
	})

	m.SetBorder(false)

	err := m.OutputFileAndClose("internal/examples/pdfs/zpl.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
}
```

## Others

* [Medium Article: Creating PDFs using Golang](https://medium.com/@johnathanfercher/creating-pdfs-using-golang-98b722e99d6d)

[travis]: https://travis-ci.com/johnfercher/maroto
[test]: test.sh

## Stargazers over time

[![Stargazers over time](https://starchart.cc/johnfercher/maroto.svg)](https://starchart.cc/johnfercher/maroto)