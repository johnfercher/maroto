# Maroto 

[![GoDoc](https://godoc.org/github.com/johnfercher/maroto?status.svg)](https://godoc.org/github.com/johnfercher/maroto)
[![Travis](https://travis-ci.com/johnfercher/maroto.svg?branch=master)][travis] 
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto) 
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)

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

* With `dep`:

```bash
dep ensure -add github.com/johnfercher/maroto
```

## Features

![result](internal/assets/images/diagram.png)

#### Grid System
* [Row](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Row)
* [Col](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Col)
* [ColSpace](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.ColSpace)
* [ColSpaces](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.ColSpaces)

#### Components To Use Inside a Col
* [Text w/ automatic new lines](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Text)
* [Signature](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Signature)
* Image ([From file](https://godoc.org/github.com/johnfercher/maroto#example-PdfMaroto-FileImage) or [Base64](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Base64Image))
* [QrCode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.QrCode)
* [Barcode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Barcode)   
    
#### Components To Use Outside a Row
* [TableList](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.TableList)
* [Line](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Line)
    
#### Components To Wrap Row, TableList and Line
* [RegisterHeader](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.RegisterHeader)

#### Others   
* Properties: most of the components has properties which you can use to customize appearance and behavior.
* [DebugMode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.SetDebugMode): Used to draw rectangles in every row and column
* Automatic New Page: New pages are generated automatically when needed.
* 100% Unicode
* Save: You can [save on disk](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.OutputFileAndClose) or export to a [base64 string](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Output)

#### Roadmap
* Create a RegisterFooter
* Increase Code Coverage
* Create a custom mock with better assertions

## Examples
In the [PDFs](internal/examples/pdfs) folder there are the PDFs generated
using Maroto, and in the [examples](internal/examples) folder there are subfolders
with the code to generate the PDFs.

![result](internal/assets/images/result.png)

#### Code
```go
package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	m.Row(40, func() {
		m.Col(func() {
			m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
				Left:    20,
				Center:  false,
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
				Percent: 75,
			})
		})
	})

	m.Line(10)

	m.Row(100, func() {
		m.Col(func() {
			m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
				Percent: 70,
			})
			m.Text("https://github.com/johnfercher/maroto", props.Text{
				Size:  20,
				Align: consts.Center,
				Top:   80,
			})
		})
	})

	m.SetDebugMode(true)

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

	m.SetDebugMode(false)

	m.OutputFileAndClose("internal/examples/internal/pdfs/zpl.pdf")
}
```

## Others

* [Medium Article: Creating PDFs using Golang](https://medium.com/@johnathanfercher/creating-pdfs-using-golang-98b722e99d6d)

[travis]: https://travis-ci.com/johnfercher/maroto
[test]: test.sh