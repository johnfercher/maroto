package pdf_test

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"time"
)

// ExampleNewMaroto demonstratos how to create maroto
func ExampleNewMaroto() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do things
	m.GetPageMargins()

	// Do more things and save...
}

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

// ExamplePdfMaroto_SetBackgroundColor demonstrates how
// to use the SetBackgroundColor method.
func ExamplePdfMaroto_SetBackgroundColor() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetBackgroundColor(color.Color{
		Red:   100,
		Green: 20,
		Blue:  30,
	})

	// This Row will be filled with the color
	m.Row(20, func() {
		m.Col(func() {
			// Add components
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	// This Row will not be filled with the color
	m.Row(20, func() {
		m.Col(func() {
			// Add components
		})
	})

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
				Size:            12.0,
				Style:           consts.BoldItalic,
				Family:          consts.Courier,
				Align:           consts.Center,
				Top:             1.0,
				Extrapolate:     false,
				VerticalPadding: 1.0,
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
	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.Font{
			Family: consts.Arial,
			Style:  consts.Bold,
			Size:   11.0,
		},
		ContentProp: props.Font{
			Family: consts.Courier,
			Style:  consts.Normal,
			Size:   10.0,
		},
		Align: consts.Center,
		AlternatedBackground: &color.Color{
			Red:   100,
			Green: 20,
			Blue:  255,
		},
		HeaderContentSpace: 10.0,
		Line:               false,
	})

	// Do more things and save...
}

// ExamplePdfMaroto_FileImage demonstrates how add an Image
// reading from disk.
// When props.Rect is nil, method make Image fulfill the context
// cell, based on width and cell from Image and cell.
// When center is true, left and top has no effect.
// Percent represents the width/height of the Image inside the cell:
// Ex: 85, means that Image will have width of 85% of column width.
// When center is false, is possible to manually positioning the Image
// with left and top.AddFromBase64(string, float64, float64, float64, float64, float64, consts.Extension)
func ExamplePdfMaroto_FileImage() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(func() {
			_ = m.FileImage("path/Image.jpg", props.Rect{
				Left:    5,
				Top:     5,
				Center:  true,
				Percent: 85,
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Base64Image demonstrates how to add an Image
// from a base64 string.
// When rect properties is nil, the method makes the Image fulfill the context
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
			_ = m.Base64Image(base64String, consts.Png, props.Rect{
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
// save a PDF object into disk.
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
// Passing nil on barcode props parameter implies the Barcode fills it's
// context cell depending on it's size.
// It's possible to define the barcode positioning through
// the top and left parameters unless center parameter is true.
// In brief, when center parameter equals true, left and top parameters has no effect.
// Percent parameter represents the Barcode's width/height inside the cell.
// i.e. Percent: 75 means that the Barcode will take up 75% of Col's width
// There is a constraint in the proportion defined, height cannot be greater than 20% of
// the width, and height cannot be smaller than 10% of the width.
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

// ExamplePdfMaroto_RegisterFooter demonstrates how to register footer.
// For register footer in Maroto you need to call method RegisterFooter
// that receives a closure.
// In this closure you are free to set any components you want to compose
// your footer.
// In this example there is a signature and a text with right align.
// It is important to remember that it is recommended to create Row's and
// Col's if necessary.
func ExamplePdfMaroto_RegisterFooter() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(func() {
				m.Signature("lorem ipsum dolor")
			})
			m.Col(func() {
				m.Text(time.Now().Format("02-January-2006"), props.Text{Align: consts.Right})
			})
		})
	})

	// Do more things or not and save...
}

// ExamplePdfMaroto_RegisterHeader demonstrates how to register header.
// For register header in Maroto you need to call method RegisterHeader
// that receives a closure.
// In this closure you are free to set any components you want to compose
// your header.
// In this example there is a two texts with different props and one image.
// It is important to remember that it is recommended to create Row's and
// Col's if necessary.
// A tip is to register the header immediately after the Maroto
// instantiation to make the code easier to read.
func ExamplePdfMaroto_RegisterHeader() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.RegisterHeader(func() {
		m.Row(10, func() {
			m.Col(func() {
				m.Text("lorem ipsum dolor", props.Text{Align: consts.Left})
			})
			m.Col(func() {
				_ = m.FileImage("internal/assets/images/frontpage.png")
			})
			m.Col(func() {
				m.Text(time.Now().Format("02-January-2006"),
					props.Text{Align: consts.Right})
			})
		})
	})

	// Do more things or not and save...
}

// ExamplePdfMaroto_SetPageMargins demonstrates how to set custom page margins.
func ExamplePdfMaroto_SetPageMargins() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetPageMargins(10, 60, 10)

	// Do more things or not and save...
}

// ExamplePdfMaroto_GetPageSize demonstrates how to obtain the current page size (width and height)
func ExamplePdfMaroto_GetPageSize() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Get
	width, height := m.GetPageSize()
	fmt.Println(width)
	fmt.Println(height)

	// Do more things and save...
}

// ExamplePdfMaroto_GetCurrentPage demonstrates how to obtain the current page index
func ExamplePdfMaroto_GetCurrentPage() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Index here will be 0
	_ = m.GetCurrentPage()

	// Add Rows, Cols and Components

	// Index here will not be 0
	_ = m.GetCurrentPage()

	// Do more things and save...
}

// ExamplePdfMaroto_GetCurrentOffset demonstrates how to obtain the current write offset
// i.e the height of cursor adding content in the pdf
func ExamplePdfMaroto_GetCurrentOffset() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Offset here will be 0
	_ = m.GetCurrentOffset()

	// Add Rows, Cols and Components until maroto add a new page

	// Offset here will not be 0
	_ = m.GetCurrentOffset()

	// Add Rows, Cols and Components to maroto add a new page

	// Offset here will be 0
	_ = m.GetCurrentOffset()

	// Do more things and save...
}

// ExamplePdfMaroto_GetPageMargins demonstrates how to obtain the current page margins
func ExamplePdfMaroto_GetPageMargins() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Get
	left, top, right, bottom := m.GetPageMargins()
	fmt.Println(left)
	fmt.Println(top)
	fmt.Println(right)
	fmt.Println(bottom)

	// Do more things and save...
}
