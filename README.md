# Maroto [![Travis](https://img.shields.io/badge/coverage-79.4%25-brightgreen.svg)][travis]
A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

## Example

#### Result
Here is the [pdf](assets/pdf/maroto.pdf) generated.

![result](assets/images/result.png)

#### Code
```go
func main() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	byteSlices, _ := ioutil.ReadFile("assets/images/gopher2.png")

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	header, contents := getContents()

	m.Row("Codes", 20, func() {
		m.Col("Logo", func() {
			m.Base64Image(base64, maroto.Png, &maroto.RectProp{
				Percent: 85,
			})
		})

		m.ColSpaces(2)

		m.Col("Link", func() {
			m.QrCode("https://github.com/johnfercher/maroto", &maroto.RectProp{
				Percent: 75,
			})
		})

		m.Col("Barcode", func() {
			id := "123456789"
			_ = m.Barcode(id, &maroto.RectProp{
				Percent: 70,
			})
			m.Text(id, &maroto.TextProp{
				Size: 8,
				Align: maroto.Center,
				Top: 17,
			})
		})
	})

	m.Line()

	m.Row("Logo", 12, func() {
		m.Col("Logo", func() {
			m.FileImage("assets/images/gopher1.jpg", nil)
		})

		m.ColSpace()

		m.Col("Definition", func() {
			m.Text("PDFGenerator: Maroto", &maroto.TextProp{
				Top: 4,
			})
			m.Text("Type: Easy & Fast", &maroto.TextProp{
				Top: 9.5,
			})
		})

		m.ColSpace()

		m.Col("Speed", func() {
			m.Text("GPL3", &maroto.TextProp{
				Size: 16,
				Style: maroto.Bold,
				Top: 7.5,
			})
		})
	})

	m.Line()

	m.Row("SubTitle", 22, func() {
		m.Col("Packages", func() {
			m.Text("Grid System", &maroto.TextProp{
				Size: 18,
				Style: maroto.Bold,
				Align: maroto.Center,
				Top: 9,
			})
			m.Text("Bootstrap Like", &maroto.TextProp{
				Size: 12,
				Align: maroto.Center,
				Top: 16,
			})
		})
	})

	m.Line()

	m.RowTableList("List", header, contents, nil)

	m.Row("Signature", 30, func() {
		m.Col("Nick", func() {
			m.Signature("Nick Fury", nil)
		})

		m.Col("Thanos", func() {
			m.Signature("Thanos", nil)
		})

		m.Col("Collector", func() {
			m.Signature("Collector", nil)
		})
	})

	_ = m.OutputFileAndClose("maroto.pdf")
}
```

[travis]: https://travis-ci.com/johnfercher/maroto