package maroto_test

import "github.com/johnfercher/maroto"

// ExamplePdfMaroto_Line demonstrates how to draw a line
// separator.
func ExamplePdfMaroto_Line() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	m.Line()

	// Do more things and save...
}

// ExamplePdfMaroto_Row demonstrates how to define a row.
func ExamplePdfMaroto_Row() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		// ... Add some columns
	})

	// Do more things and save...
}

// ExamplePdfMaroto_ColSpace demonstrates how to add
// a empty column inside a row.
func ExamplePdfMaroto_ColSpace() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.ColSpace()
	})

	// Do more things and save...
}

// ExamplePdfMaroto_ColSpaces demonstrates how to add
// some empty columns inside a row.
func ExamplePdfMaroto_ColSpaces() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.ColSpaces(2)
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Col demonstrates how to add
// a useful column
func ExamplePdfMaroto_Col() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.Col("MyCol", func() {
			// Add Image, Text, Signature, QrCode or Barcode...
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_SetDebugMode demonstrates how to
// define debug mode
func ExamplePdfMaroto_SetDebugMode() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	m.SetDebugMode(true)

	// Add some Rows, Cols, Lines and etc...
	// Here will be drawn borders in every cell

	m.SetDebugMode(false)

	// Add some Rows, Cols, Lines and etc...
	// Here will not be drawn borders

	// Do more things and save...
}

// ExamplePdfMaroto_GetDebugMode demonstrates how to
// obtain the actual debug mode value
func ExamplePdfMaroto_GetDebugMode() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	// false
	m.GetDebugMode()

	m.SetDebugMode(true)

	// true
	m.GetDebugMode()

	// Do more things and save...
}

// ExamplePdfMaroto_Text demonstrates how to add
// a Text inside a col. Passing nil on fontProp make the method
// use: arial Font, normal style, size 10.0 and align left.
// Not passing family, make method use arial.
// Not passing style, make method use normal.
// Not passing size, make method use 10.0.
// Not passing align, make method use left.
func ExamplePdfMaroto_Text() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.Col("MyCol", func() {
			m.Text("TextContent", &maroto.TextProp{
				Size:   12.0,
				Style:  maroto.BoldItalic,
				Family: maroto.Courier,
				Align:  maroto.Center,
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
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.Col("MyCol", func() {
			m.Signature("LabelForSignature", &maroto.SignatureProp{
				Size:   12.0,
				Style:  maroto.BoldItalic,
				Family: maroto.Courier,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_RowTableList demonstrates how to add a table
// with multiple rows and columns
func ExamplePdfMaroto_RowTableList() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	headers := []string{"Header1", "Header2"}
	contents := [][]string{
		{"Content1", "Content2"},
		{"Content3", "Content3"},
	}

	// 1 Row of header
	// 2 Rows of contents
	// Each row have 2 columns
	m.RowTableList("RowTableList1", headers, contents, nil)

	// Do more things and save...
}

// ExamplePdfMaroto_FileImage demonstrates how add a Image
// reading from disk.
// When rectProp is nil, method make Image fullfill the context
// cell, based on width and cell from Image and cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the Image inside the cell:
// Ex: 85, means that Image will have width of 85% of column width.
// When center is false, is possible to manually positioning the Image
// with left and top.
func ExamplePdfMaroto_FileImage() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0

	m.Row("MyRow", rowHeight, func() {
		m.Col("MyCol", func() {
			m.FileImage("path/Image.jpg", &maroto.RectProp{
				Left:    5,
				Top:     5,
				Center:  true,
				Percent: 85,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Base64Image demonstrates how add a Image
// reading a base64 string.
// When rectProp is nil, method make Image fullfill the context
// cell, based on width and cell from Image and cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the Image inside the cell:
// Ex: 85, means that Image will have width of 85% of column width.
// When center is false, is possible to manually positioning the Image
// with left and top.
func ExamplePdfMaroto_Base64Image() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)
	rowHeight := 5.0
	base64String := "y7seWGHE923Sdgs..."

	m.Row("MyRow", rowHeight, func() {
		m.Col("MyCol", func() {
			m.Base64Image(base64String, maroto.Png, &maroto.RectProp{
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
// save a PDF in disk.
func ExamplePdfMaroto_OutputFileAndClose() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	// Do a lot of things on rows and columns...

	err := m.OutputFileAndClose("path/file.pdf")
	if err != nil {
		return
	}
}

// ExamplePdfMaroto_Output demonstrates how to get a
// base64 string from PDF
func ExamplePdfMaroto_Output() {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	// Do a lot of things on rows and columns...

	_, err := m.Output()
	if err != nil {
		return
	}
}
