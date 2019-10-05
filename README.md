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
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"io/ioutil"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	byteSlices, _ := ioutil.ReadFile("internal/assets/images/biplane.jpg")

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	headerSmall, smallContent := getSmallContent()
	headerMedium, mediumContent := getMediumContent()

	m.RegisterHeader(func() {

		m.Row(20, func() {
			m.Col(func() {
				m.Base64Image(base64, consts.Jpg, props.Rect{
					Percent: 70,
				})
			})

			m.ColSpaces(2)

			m.Col(func() {
				m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
					Percent: 75,
				})
			})

			m.Col(func() {
				id := "https://github.com/johnfercher/maroto"
				_ = m.Barcode(id, props.Barcode{
					Proportion: props.Proportion{50, 10},
					Percent:    75,
				})
				m.Text(id, props.Text{
					Size:  7,
					Align: consts.Center,
					Top:   16,
				})
			})
		})

		m.Line(1.0)

		m.Row(12, func() {
			m.Col(func() {
				m.FileImage("internal/assets/images/gopherbw.png")
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("Packages Report: Daily", props.Text{
					Top: 4,
				})
				m.Text("Type: Small, Medium", props.Text{
					Top: 10,
				})
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("20/07/1994", props.Text{
					Size:   10,
					Style:  consts.BoldItalic,
					Top:    7.5,
					Family: consts.Helvetica,
				})
			})
		})

		m.Line(1.0)

		m.Row(22, func() {
			m.Col(func() {
				m.Text(fmt.Sprintf("Small: %d, Medium %d", len(smallContent), len(mediumContent)), props.Text{
					Size:  15,
					Style: consts.Bold,
					Align: consts.Center,
					Top:   9,
				})
				m.Text("Brasil / São Paulo", props.Text{
					Size:  12,
					Align: consts.Center,
					Top:   17,
				})
			})
		})

		m.Line(1.0)

	})

	m.RegisterFooter(func() {
		m.Row(40, func() {
			m.Col(func() {
				m.Signature("Signature 1", props.Font{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9,
				})
			})

			m.Col(func() {
				m.Signature("Signature 2")
			})

			m.Col(func() {
				m.Signature("Signature 3")
			})
		})
	})

	m.Row(15, func() {
		m.Col(func() {
			m.Text("Small Packages / 39u.", props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerSmall, smallContent)

	m.Row(15, func() {
		m.Col(func() {
			m.Text("Medium Packages / 22u.", props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerMedium, mediumContent, props.TableList{
		Align: consts.Center,
		HeaderProp: props.Font{
			Family: consts.Courier,
			Style:  consts.BoldItalic,
		},
		ContentProp: props.Font{
			Family: consts.Courier,
			Style:  consts.Italic,
		},
	})

	_ = m.OutputFileAndClose("internal/examples/pdfs/sample1.pdf")
}
```

## Others

* [Medium Article: Creating PDFs using Golang](https://medium.com/@johnathanfercher/creating-pdfs-using-golang-98b722e99d6d)

[travis]: https://travis-ci.com/johnfercher/maroto
[test]: test.sh

## Stargazers over time

[![Stargazers over time](https://starchart.cc/johnfercher/maroto.svg)](https://starchart.cc/johnfercher/maroto)