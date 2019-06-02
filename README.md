# Maroto [![Travis](https://img.shields.io/badge/coverage-46.4%25-orange.svg)][travis]
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
	//m.SetDebugMode(true)

	byteSlices, _ := ioutil.ReadFile("assets/images/gopher2.png")

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	header, contents := getContents()

	m.Row("Codes", 20, func() {
		m.Col("Logo", func() {
			m.Base64Image(base64, maroto.Png, &maroto.RectProp{
				Percent: 45,
			})
		})

		m.ColSpace()

		m.Col("Link", func() {
			m.QrCode("https://github.com/johnfercher/maroto")
		})

		m.Col("Barcode", func() {
			id := "123456789"
			_ = m.Barcode(id, 30, 9, 5)
			m.Text(id, 17, &maroto.FontProp{
				Size: 8,
				Align: maroto.Center,
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
			m.Text("PDFGenerator: Maroto", 4, nil)
			m.Text("Type: Easy & Fast", 9.5, nil)
		})

		m.ColSpace()

		m.Col("Speed", func() {
			m.Text("GPL3", 7.5, &maroto.FontProp{
				Size: 16,
				Style: maroto.Bold,
			})
		})
	})

	m.Line()

	m.Row("SubTitle", 22, func() {
		m.Col("Packages", func() {
			m.Text("Grid System", 9, &maroto.FontProp{
				Size: 18,
				Style: maroto.Bold,
				Align: maroto.Center,
			})
			m.Text("Bootstrap Like", 16, &maroto.FontProp{
				Size: 12,
				Align: maroto.Center,
			})
		})
	})

	m.Line()

	m.RowTableList("List", header, contents)

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

	m.OutputFileAndClose("maroto.pdf")
}
```

[travis]: https://travis-ci.com/johnfercher/maroto