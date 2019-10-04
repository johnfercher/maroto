package pdf_test

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// ExamplePdfMaroto_Line demonstrates how to draw a line
// separator.
func ExamplePdfMaroto_Line() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.Line(1.0)

	// Do more things and save...
}

// ExamplePdfMaroto_Row demonstrates how to define a row.
func ExamplePdfMaroto_Row() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		// ... Add some columns
	})

	// Do more things and save...
}

// ExamplePdfMaroto_ColSpace demonstrates how to add
// an empty column inside a row.
func ExamplePdfMaroto_ColSpace() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.ColSpace()
	})

	// Do more things and save...
}

// ExamplePdfMaroto_ColSpaces demonstrates how to add
// some empty columns inside a row.
func ExamplePdfMaroto_ColSpaces() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.ColSpaces(2)
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Col demonstrates how to add
// an useful column
func ExamplePdfMaroto_Col() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			// Add Image, Text, Signature, QrCode or Barcode...
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_SetBorder demonstrates how to
// enable the line drawing in every cell
func ExamplePdfMaroto_SetBorder() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(true)

	// Add some Rows, Cols, Lines and etc...
	// Here will be drawn borders in every cell

	m.SetBorder(false)

	// Add some Rows, Cols, Lines and etc...
	// Here will not be drawn borders

	// Do more things and save...
}

// ExamplePdfMaroto_GetBorder demonstrates how to
// obtain the actual borders status
func ExamplePdfMaroto_GetBorder() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// false
	m.GetBorder()

	m.SetBorder(true)

	// true
	m.GetBorder()

	// Do more things and save...
}

// ExamplePdfMaroto_Text demonstrates how to add
// a Text inside a col. Passing nil on fontProp makes the method
// use: arial Font, normal style, size 10.0 and align left.
// Not passing family, makes the method use arial.
// Not passing style, makes the method use normal.
// Not passing size, makes the method use 10.0.
// Not passing align, makes the method use left.
func ExamplePdfMaroto_Text() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			m.Text("TextContent", props.Text{
				Size:   12.0,
				Style:  consts.BoldItalic,
				Family: consts.Courier,
				Align:  consts.Center,
				Top:    1.0,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Signature demonstrates how to add
// a Signature space inside a col. Passing nil on signatureProp make the method
// use: arial Font, normal style and size 10.0.
// Not passing family, make method use arial.
// Not passing style, make method use normal.
// Not passing size, make method use 10.0.
func ExamplePdfMaroto_Signature() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			m.Signature("LabelForSignature", props.Font{
				Size:   12.0,
				Style:  consts.BoldItalic,
				Family: consts.Courier,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_TableList demonstrates how to add a table
// with multiple rows and columns
func ExamplePdfMaroto_TableList() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	headers := []string{"Header1", "Header2"}
	contents := [][]string{
		{"Content1", "Content2"},
		{"Content3", "Content3"},
	}

	// 1 Row of header
	// 2 Rows of contents
	// Each row have 2 columns
	m.TableList(headers, contents)

	// Do more things and save...
}

// ExamplePdfMaroto_FileImage demonstrates how add an Image
// reading from disk.
// When barcodeProp is nil, method make Image fulfill the context
// cell, based on width and cell from Image and cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the Image inside the cell:
// Ex: 85, means that Image will have width of 85% of column width.
// When center is false, is possible to manually positioning the Image
// with left and top.
func ExamplePdfMaroto_FileImage() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			m.FileImage("path/Image.jpg", props.Rect{
				Left:    5,
				Top:     5,
				Center:  true,
				Percent: 85,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Base64Image demonstrates how add an Image
// from a base64 string.
// When barcodeProp is nil, the method makes the Image fulfill the context
// cell, based on width and height from Image and cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the Image inside the cell:
// Ex: 85, means that Image will have width of 85% of column width.
// When center is false, is possible to manually positioning the Image
// with left and top.
func ExamplePdfMaroto_Base64Image() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0
	base64String := "y7seWGHE923Sdgs..."

	m.Row(rowHeight, func() {
		m.Col(func() {
			m.Base64Image(base64String, consts.Png, props.Rect{
				Left:    5,
				Top:     5,
				Center:  true,
				Percent: 85,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_OutputFileAndClose demonstrates how to
// save a PDF into disk.
func ExamplePdfMaroto_OutputFileAndClose() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do a lot of things on rows and columns...

	err := m.OutputFileAndClose("path/file.pdf")
	if err != nil {
		return
	}
}

// ExamplePdfMaroto_Output demonstrates how to get a
// base64 string from PDF
func ExamplePdfMaroto_Output() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do a lot of things on rows and columns...

	_, err := m.Output()
	if err != nil {
		return
	}
}

// ExamplePdfMaroto_QrCode demonstrates how to add
// a QR Code inside a Col. Passing nil on rectProps makes
// the QR Code fills the context cell depending on width
// and height of the QR Code and the cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the QR Code inside the cell.
// i.e. 80 means that the QR Code will take up 80% of Col's width
// When center is false, positioning of the QR Code can be done through
// left and top.
func ExamplePdfMaroto_QrCode() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			m.QrCode("https://godoc.org/github.com/johnfercher/maroto", props.Rect{
				Left:    5,
				Top:     5,
				Center:  false,
				Percent: 80,
			})
		})
	})
}

// ExamplePdfMaroto_Barcode demonstrates how to place a barcode inside
// a Col.
// Passing nil on props parameter implies the Barcode fills it's
// context cell depending on it's size.
// It's possible to define the barcode positioning through
// the top and left parameters unless center parameter is true.
// In brief, when center parameter equals true, left and top parameters has no effect.
// Percent parameter represents the Barcode's width/height inside the cell.
// i.e. Percent: 75 means that the Barcode will take up 75% of Col's width
func ExamplePdfMaroto_Barcode() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do a lot of things on rows and columns...

	m.Col(func() {
		_ = m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
			Percent:    75,
			Proportion: props.Proportion{Width: 50, Height: 10},
			Center:     true,
		})
	})

	// do more things...
}
