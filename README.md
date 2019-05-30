# Maroto
A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

## Example

#### Result
Here is the [pdf](assets/pdf/maroto.pdf) generated.

![result](assets/images/result.png)

#### Code
```go
func main() {
	m := maroto.NewMaroto(maroto.Vertical, maroto.A4)
	//m.SetDebugMode(true)
	header, contents := getContents()

	m.Row("Barcode", 20, func() {
		m.Col("Logo", func() {
			m.Image("assets/images/marvel.jpg", 0)
		})

		m.ColSpaces(2)

		m.Col("Barcode", func() {
			id := "123456789"
			_ = m.Barcode(id, 30, 9, 5)
			m.Text(id, maroto.Arial, maroto.Bold, 8, 17, maroto.CenterH)
		})
	})

	m.Line()

	m.Row("Logo", 12, func() {
		m.Col("Logo", func() {
			m.Image("assets/images/shape.jpg", 0)
		})

		m.ColSpace()

		m.Col("Definition", func() {
			m.Text("Definition: Random", maroto.Arial, maroto.Normal, 9, 5, maroto.Left)
			m.Text("Type: Shocks", maroto.Arial, maroto.Normal, 9, 9, maroto.Left)
		})

		m.ColSpace()

		m.Col("Speed", func() {
			m.Text("FAST", maroto.Arial, maroto.Bold, 15, 7.5, maroto.CenterH)
		})
	})

	m.Line()

	m.Row("SubTitle", 22, func() {
		m.ColSpaces(2)

		m.Col("Packages", func() {
			m.Text("SUPCODE", maroto.Arial, maroto.Bold, 20, 10.5, maroto.CenterH)
			m.Text("1q2w3e4r", maroto.Arial, maroto.Normal, 12, 16, maroto.CenterH)
		})

		m.ColSpaces(2)
	})

	m.Line()

	m.RowTableList("List", header, contents)

	m.Row("Signature", 50, func() {
		m.Col("Nick", func() {
			m.Sign("Nick Fury", maroto.Arial, maroto.Bold, 8)
		})

		m.ColSpace()

		m.Col("Thanos", func() {
			m.Sign("Thanos", maroto.Arial, maroto.Bold, 8)
		})

		m.ColSpace()

		m.Col("Collector", func() {
			m.Sign("Collector", maroto.Arial, maroto.Bold, 8)
		})
	})

	m.OutputFileAndClose("maroto.pdf")
}
```