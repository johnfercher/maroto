# Maroto [![Travis](https://img.shields.io/badge/coverage-95.5%25-brightgreen.svg)][travis]
A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

## Features
* Grid System
    * [Row](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Row)
    * [Col](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Col)
    * [ColSpace](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.ColSpace)
    * [ColSpaces](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.ColSpaces)

* Components To Use Inside a Col
    * [Text](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Text)
    * [Signature](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Signature)
    * Image ([From file](https://godoc.org/github.com/johnfercher/maroto#example-PdfMaroto-FileImage) or [Base64](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Base64Image))
    * [QrCode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.QrCode)
    * [Barcode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Barcode)   
    
* Components To Use Outside a Row
    * [TableList](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.TableList)
    * [Line](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Line)
    
* Components To Wrap Row, TableList and Line
    * [RegisterHeader](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.RegisterHeader)
    
* Properties: most of the components has properties which you can use to customize appearance and behavior.
* [DebugMode](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.SetDebugMode): Used to draw rectangles in every row and column
* Automatic New Page: New pages are generated automatically when needed.
* 100% Unicode
* Save: You can [save on disk](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.OutputFileAndClose) or export to a [base64 string](https://godoc.org/github.com/johnfercher/maroto#PdfMaroto.Output)

**TODO**
* RegisterFooter
* Increase Code Coverage
* Create a custom mock with better assertions
## Example

#### Result
Here is the [pdf](assets/pdf/maroto.pdf) generated.

![result](assets/images/result.png)

#### Code
```go
func main() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	//m.SetDebugMode(true)

	byteSlices, _ := ioutil.ReadFile("assets/images/gopher2.png")

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	header, contents := getContents()

	m.RegisterHeader(func() {

		// Image, Barcode and QrCode
		m.Row(20, func() {
			m.Col(func() {
				m.Base64Image(base64, maroto.Png, &maroto.RectProp{
					Percent: 85,
				})
			})

			m.ColSpaces(2)

			m.Col(func() {
				m.QrCode("https://github.com/johnfercher/maroto", &maroto.RectProp{
					Percent: 75,
				})
			})

			m.Col(func() {
				id := "123456789"
				_ = m.Barcode(id, &maroto.RectProp{
					Percent: 70,
				})
				m.Text(id, &maroto.TextProp{
					Size:  8,
					Align: maroto.Center,
					Top:   17,
				})
			})
		})

		m.Line(1.0)

		// Image and Old License
		m.Row(12, func() {
			m.Col(func() {
				m.FileImage("assets/images/gopher1.jpg", nil)
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("PDFGenerator: Maroto", &maroto.TextProp{
					Top: 4,
				})
				m.Text("Type: Easy & Fast", &maroto.TextProp{
					Top: 10,
				})
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("GPL3", &maroto.TextProp{
					Size:  16,
					Style: maroto.Bold,
					Top:   8,
				})
			})
		})

		m.Line(1.0)

		// Features
		m.Row(22, func() {
			m.Col(func() {
				m.Text("Grid System", &maroto.TextProp{
					Size:  18,
					Style: maroto.Bold,
					Align: maroto.Center,
					Top:   9,
				})
				m.Text("Bootstrap Like + Úñîçòdë", &maroto.TextProp{
					Size:  12,
					Align: maroto.Center,
					Top:   17,
				})
			})
		})

		m.Line(1.0)

	})

	m.TableList(header, contents, nil)

	// Signatures
	m.Row(30, func() {
		m.Col(func() {
			m.Signature("Signature 1", nil)
		})

		m.Col(func() {
			m.Signature("Signature 2", nil)
		})

		m.Col(func() {
			m.Signature("Signature 3", nil)
		})
	})

	_ = m.OutputFileAndClose("maroto.pdf")
}
```

[travis]: https://travis-ci.com/johnfercher/maroto