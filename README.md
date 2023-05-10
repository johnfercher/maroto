# Maroto 

[![GoDoc](https://godoc.org/github.com/johnfercher/maroto?status.svg)](https://godoc.org/github.com/johnfercher/maroto)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  
[![CI](https://github.com/johnfercher/maroto/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto)

A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

You can write your PDFs like you are creating a site using Bootstrap. A Row may have many Cols, and a Col may have many components. 
Besides that, pages will be added when content may extrapolate the useful area. You can define a header which will be added
always when a new page appear, in this case, a header may have many rows, lines or tablelist. 

## Installation

* With `go get`:

```bash
go get -u github.com/johnfercher/maroto/internal
```

## Contributing

| Command         | Description                                       | Dependencies                                                 |
|-----------------|---------------------------------------------------|--------------------------------------------------------------|
| `make build`    | Build project                                     | `go`                                                         |
| `make test`     | Run unit tests                                    | `go`                                                         |
| `make fmt`      | Format files                                      | `gofmt`, `gofumpt` and `goimports`                           |
| `make lint`     | Check files                                       | `golangci-lint` and `goreportcard-cli`                       |
| `make dod`      | (Definition of Done) Format files and check files | Same as`make build`, `make test`, `make fmt` and `make lint` | 
| `make install`  | Install all dependencies                          | `go`, `curl` and `git`                                       |
| `make examples` | Run all examples                                  | `go`                                                         |

## Features
![result](internal/assets/images/diagram.png)

#### Constructors
* [New](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#NewMaroto)
* [NewCustomSize](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#NewMarotoCustomSize)

#### Grid System
* [Row](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Row)
* [Col](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Col)
* [ColSpace](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.ColSpace)

#### Components To Use Inside a Col
* [Text w/ automatic new lines](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Text)
* [Signature](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Signature)
* Image ([From file](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.FileImage) or [Base64](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Base64Image))
* [QrCode](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.QrCode)
* [Barcode](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Barcode) 
* [DataMatrixCode](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.DataMatrixCode)  
    
#### Components To Use Outside a Row
* [TableList](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.TableList)
* [Line](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.Line)
    
#### Components To Wrap Row, TableList and Line
* [RegisterHeader](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.RegisterHeader)
* [RegisterFooter](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.RegisterFooter)

#### Saving PDF
* [Output](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.Output): Generate PDF and return a Base64 string
* [OutputFileAndClose](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.OutputFileAndClose): Generate PDF and save in disk

#### Metadata
* [SetCompression](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetCompression) Set/ unset compression for a page. Default value is on
* [SetProtection](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetProtection): Define a password to open the PDF
* [SetAuthor](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetAuthor): Define the author
* [SetCreator](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetCreator): Define the creator
* [SetSubject](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetSubject): Define the subject
* [SetTitle](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetTitle): Define the title
* [SetCreationDate](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetCreationDate): Define the creation date

#### Fonts   
* [AddUTF8Font](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.AddUTF8Font): Add a custom utf8 font and allow any character of any language.
* [SetDefaultFontFamily](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetProtection): Define a default font family, useful to use with a custom font.
* [SetFontLocation](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetFontLocation): Define the default path to search the custom font
* 100% Unicode

#### Page
* [AddPage](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.AddPage): Skip the current page and start the build a new one
* [SetPageMargins](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.SetPageMargins): Customize the page margins
* [SetAliasNbPages](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetAliasNbPages): Set placeholder to use in texts for total count of pages
* [SetFirstPageNb](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf#PdfMaroto.SetFirstPageNb): Set first number for page numbering
* Automatic New Page: New pages are generated automatically when needed.

#### Customizations
* [Properties](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/props?tab=doc): most of the components has properties which you can use to customize appearance and behavior
* [SetBorder](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.SetBorder): Used to draw rectangles in every row and column
* [SetBackgroundColor](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.SetBackgroundColor): Used to change the background color of one cell
* [SetMaxGridSum](https://pkg.go.dev/github.com/johnfercher/maroto/pkg/pdf?tab=doc#PdfMaroto.SetMaxGridSum): Sets max amount of cols per row, default is 12 cols per row

#### Roadmap
* Updated in [Issues](https://github.com/johnfercher/maroto/issues)

## Examples
In the [PDFs](internal/examples/pdfs) folder there are the PDFs generated
using Maroto, and in the [examples](internal/examples) folder there are subfolders
with the code to generate the PDFs. There are examples of: [barcode](internal/examples/barcode),
[billing](internal/examples/billing), [certificate](internal/examples/certificate), 
[custom pdf size](internal/examples/customsize), [image inside grid](internal/examples/imagegrid),
[qrcode inside grid](internal/examples/qrgrid), [sample with almost all features together](internal/examples/sample1), 
[text inside grid](internal/examples/textgrid), [custom utf8 fonts (any language)](internal/examples/utfsample) and a
[label zpl](internal/examples/zpl).

![result](internal/assets/images/result.png)

## Other Examples

| Title | Media |
|---|---|
| [How to create PDFs with Go, Maroto & GoFakeIt](https://www.youtube.com/watch?v=jwOy4JgleTU) | Video |
| [Creating a PDF with Go, Maroto & Gofakeit](https://divrhino.com/articles/create-pdf-document-with-go-maroto-gofakeit) | Article |
| [divrhino/fruitful-pdf](https://github.com/divrhino/fruitful-pdf) | Repository |
| [Creating PDFs using Golang](https://medium.com/@johnathanfercher/creating-pdfs-using-golang-98b722e99d6d) | Article |

#### Code
> This is part of the example [billing](internal/examples/billing).
```go
// Billing example
package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

func main() {
	begin := time.Now()

	darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	whiteColor := color.NewWhite()
	header := getHeader()
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
					Center:  true,
					Percent: 80,
				})
			})

			m.ColSpace(6)

			m.Col(3, func() {
				m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
				})
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   12,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
			})
		})
	})

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice ABC123456789", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Transactions", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(15, func() {
		m.Col(6, func() {
			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			})
			m.Text("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			})
		})
		m.ColSpace(6)
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/johnfercher/maroto.svg)](https://starchart.cc/johnfercher/maroto)
