package pdf_test

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// ExampleNewMaroto demonstrates how to create maroto.
func ExampleNewMaroto() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do things
	m.GetPageMargins()

	// Do more things and save...
}

// ExampleNewMaroto demonstrates how to create maroto with custom page size.
func ExampleNewMarotoCustomSize() {
	m := pdf.NewMarotoCustomSize(consts.Landscape, "C6", "mm", 114.0, 162.0)

	// Do things
	m.GetPageMargins()

	// Do more things and save...
}

// ExamplePdfMaroto_AddPage how to force add a new page.
func ExamplePdfMaroto_AddPage() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Add rows, cols and components
	m.AddPage()

	// Add rows, col and components in a new page
	// Do more things and save...
}

// ExamplePdfMaroto_Row demonstrates how to define a row.
func ExamplePdfMaroto_Row() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	// Warning: There is no way to use a row inside a row or a row inside a col.
	m.Row(rowHeight, func() {
		m.Col(12, func() {
			// Add a component
		})
	})

	// Warning: There is no way to use a row inside a row or a row inside a col.
	m.Row(rowHeight, func() {
		m.Col(12, func() {
			// Add another component
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Col demonstrates how to add
// an useful column.
func ExamplePdfMaroto_Col() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	// Warning: The sum of all column gridSize cannot extrapolate 12
	// Warning: There is no way to use a row inside a row or a row inside a col.
	m.Row(rowHeight, func() {
		m.Col(12, func() {
			// Add Image, Text, Signature, QrCode or Barcode...
		})
	})

	// Warning: The sum of all column gridSize cannot extrapolate 12
	// Warning: There is no way to use a row inside a row or a row inside a col.
	m.Row(rowHeight, func() {
		m.Col(6, func() {
			// Add Image, Text, Signature, QrCode or Barcode...
		})
		m.Col(6, func() {
			// Add Image, Text, Signature, QrCode or Barcode...
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_ColSpace demonstrates how to add
// an empty column inside a row.
func ExamplePdfMaroto_ColSpace() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	// Warning: The sum of all column gridSize cannot extrapolate 12
	m.Row(rowHeight, func() {
		m.ColSpace(12)
	})

	// Warning: The sum of all column gridSize cannot extrapolate 12
	m.Row(rowHeight, func() {
		m.ColSpace(6)
		m.ColSpace(6)
	})

	// Do more things and save...
}

// ExamplePdfMaroto_RegisterHeader demonstrates how to register header.
func ExamplePdfMaroto_RegisterHeader() {
	// For register header in Maroto you need to call method RegisterHeader
	// that receives a closure.
	// In this closure you are free to set any components you want to compose
	// your header.
	// In this example there is a two texts with different props and one image.
	// It is important to remember that it is recommended to create Row's and
	// Col's if necessary.
	// You have to register the header immediately after the Maroto

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.RegisterHeader(func() {
		m.Row(10, func() {
			m.Col(3, func() {
				m.Text("lorem ipsum dolor", props.Text{Align: consts.Left})
			})
			m.Col(3, func() {
				_ = m.FileImage("internal/assets/images/frontpage.png")
			})
			m.Col(3, func() {
				m.Text(time.Now().Format("02-January-2006"),
					props.Text{Align: consts.Right})
			})
		})
	})

	// Do more things or not and save...
}

// ExamplePdfMaroto_RegisterFooter demonstrates how to register footer.
func ExamplePdfMaroto_RegisterFooter() {
	// For register footer you need to call method RegisterFooter
	// that receives a closure.
	// In this closure you are free to set any components you want to compose
	// your footer.
	// In this example there is a signature and a text with right align.
	// It is important to remember that it is recommended to create Row's and
	// Col's if necessary.
	// You have to register the footer immediately after the Maroto
	// All footers will be rendered at the bottom of all pages

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(6, func() {
				m.Signature("lorem ipsum dolor")
			})
			m.Col(6, func() {
				m.Text(time.Now().Format("02-January-2006"), props.Text{Align: consts.Right})
			})
		})
	})

	// Do more things or not and save...
}

// ExamplePdfMaroto_TableList demonstrates how to add a table
// with multiple rows and columns.
func ExamplePdfMaroto_TableList() {
	// Not passing this table list prop will lead the method to use all the follow values.
	// Not passing HeaderProp.Size, make the method use 10.
	// Not passing HeaderProp.Family, make the method use arial.
	// Not passing HeaderProp.Style, make the method use bold.
	// Not passing HeaderProp.GridSizes, make the method use an array with same length
	// Not passing HeaderProp.Color, make the method use a black font
	// of header array, the values will be perfectly divided to make all columns with the same size.
	// Not passing Align, make the method to use left.
	// Not passing ContentProp.Size, make the method use 10.
	// Not passing ContentProp.Family, make the method use arial.
	// Not passing ContentProp.Style, make the method use normal.
	// Not passing ContentProp.Color, make the method use a black font
	// Not passing Content.GridSizes, make the method use an array with same length
	// of content array in the first line, the values will be perfectly divided to make all columns with the same size.
	// Not passing HeaderContentSpace, will make the method use 4.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	headers := []string{"Header1", "Header2"}
	contents := [][]string{
		{"Content1", "Content2"},
		{"Content3", "Content3"},
	}

	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      11.0,
			GridSizes: []uint{3, 9},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      10.0,
			GridSizes: []uint{3, 9},
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

	// TableList have to be used at same level as row
	m.Row(10, func() {
		m.Col(12, func() {
			// Add a component
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Line demonstrates how to draw a line
// separator.
func ExamplePdfMaroto_Line() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Not passing Prop.Color make line as black.
	// Not passing width make width as 0.1.
	// Not passing style make style as solid.
	m.Line(1.0, props.Line{
		Color: color.Color{
			Red:   255,
			Green: 100,
			Blue:  50,
		},
		Style: consts.Dotted,
		Width: 1.0,
	})

	// Do more things and save...
}

// ExamplePdfMaroto_Text demonstrates how to add
// a Text inside a col.
func ExamplePdfMaroto_Text() {
	// Not passing the text prop will lead to use all the follow default values.
	// Not passing family, makes the method use arial.
	// Not passing style, makes the method use normal.
	// Not passing size, makes the method use 10.0.
	// Not passing align, makes the method use left.
	// Not passing extrapolate, makes the method use false.
	// Not passing color, makes the method use the current color.
	// Top cannot be less than 0.
	// VerticalPadding cannot be less than 0.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(12, func() {
			m.Text("TextContent", props.Text{
				Size:            12.0,
				Style:           consts.BoldItalic,
				Family:          consts.Courier,
				Align:           consts.Center,
				Top:             1.0,
				Extrapolate:     false,
				VerticalPadding: 1.0,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_FileImage demonstrates how add an Image
// reading from disk.
func ExamplePdfMaroto_FileImage() {
	// When props.Rect is nil, method make Image fulfill the context cell.
	// When center is true, left and top has no effect.
	// Percent represents the width/height of the Image inside the cell,
	// Ex: 85, means that Image will have width of 85% of column width.
	// When center is false, is possible to manually positioning the Image.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(12, func() {
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
func ExamplePdfMaroto_Base64Image() {
	// When props.Rect is nil, method make Image fulfill the context cell.
	// When center is true, left and top has no effect.
	// Percent represents the width/height of the Image inside the cell,
	// Ex: 85, means that Image will have width of 85% of column width.
	// When center is false, is possible to manually positioning the Image.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	// Bytes of the image loaded
	bytes := []byte{1, 2, 3}
	base64String := base64.StdEncoding.EncodeToString(bytes)

	m.Row(rowHeight, func() {
		m.Col(12, func() {
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

// ExamplePdfMaroto_Barcode demonstrates how to place a barcode inside
// a Col.
func ExamplePdfMaroto_Barcode() {
	// Passing nil on barcode props parameter implies the Barcode fills it's
	// context cell depending on it's size.
	// It's possible to define the barcode positioning through
	// the top and left parameters unless center parameter is true.
	// In brief, when center parameter equals true, left and top parameters has no effect.
	// Percent parameter represents the Barcode's width/height inside the cell.
	// i.e. Percent: 75 means that the Barcode will take up 75% of Col's width
	// There is a constraint in the proportion defined, height cannot be greater than 20% of
	// the width, and height cannot be smaller than 10% of the width.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do a lot of things on rows and columns...

	m.Col(12, func() {
		_ = m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
			Percent:    75,
			Proportion: props.Proportion{Width: 50, Height: 10},
			Center:     true,
		})
	})

	// do more things...
}

// ExamplePdfMaroto_QrCode demonstrates how to add
// a QR Code inside a Col.
func ExamplePdfMaroto_QrCode() {
	// Passing nil on rectProps makes
	// the QR Code fills the context cell.
	// When center is true, left and top has no effect.
	// Percent represents the width/height of the QR Code inside the cell.
	// i.e. 80 means that the QR Code will take up 80% of Col's width
	// When center is false, positioning of the QR Code can be done through
	// left and top.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(12, func() {
			m.QrCode("https://godoc.org/github.com/johnfercher/maroto", props.Rect{
				Left:    5,
				Top:     5,
				Center:  false,
				Percent: 80,
			})
		})
	})
}

// ExamplePdfMaroto_DataMatrixCode demonstrates how to add
// a DataMatrixCode inside a Col.
func ExamplePdfMaroto_DataMatrixCode() {
	// Passing nil on rectProps makes
	// the DataMatrixCode fill the context cell.
	// When center is true, left and top has no effect.
	// Percent represents the width/height of the DataMatrixCode inside the cell.
	// i.e. 80 means that the DataMatrixCode will take up 80% of Col's width
	// When center is false, positioning of the DataMatrixCode can be done through
	// left and top.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(12, func() {
			m.DataMatrixCode("https://godoc.org/github.com/johnfercher/maroto", props.Rect{
				Left:    5,
				Top:     5,
				Center:  false,
				Percent: 80,
			})
		})
	})
}

// ExamplePdfMaroto_Signature demonstrates how to add
// a Signature space inside a col.
func ExamplePdfMaroto_Signature() {
	// Not passing this font prop will lead the method to use all the follow values.
	// Not passing family, make method use arial.
	// Not passing style, make method use normal.
	// Not passing size, make method use 10.0.
	// Not passing color, make method use the current color.

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	rowHeight := 5.0

	m.Row(rowHeight, func() {
		m.Col(12, func() {
			m.Signature("LabelForSignature", props.Font{
				Size:   12.0,
				Style:  consts.BoldItalic,
				Family: consts.Courier,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
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
// base64 string from PDF.
func ExamplePdfMaroto_Output() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Do a lot of things on rows and columns...

	_, err := m.Output()
	if err != nil {
		return
	}
}

// ExamplePdfMaroto_SetBorder demonstrates how to
// enable the line drawing in every cell.
func ExamplePdfMaroto_SetBorder() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(true)

	// Add some Rows, Cols, Lines and etc...
	// Here will be drawn borders in every cell.

	m.SetBorder(false)

	// Add some Rows, Cols, Lines and etc...
	// Here will not be drawn borders.

	// Do more things and save...
}

// ExamplePdfMaroto_SetFirstPageNb demonstrates
// how to use SetFirstPageNb method.
func ExamplePdfMaroto_SetFirstPageNb() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Ths will set first page index to 1.
	m.SetFirstPageNb(1)
}

// ExamplePdfMaroto_SetAliasNbPages demonstrates
// how to use SetAliasNbPages method.
func ExamplePdfMaroto_SetAliasNbPages() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Ths will create placeholder "{nbs}" to use in texts for total count of pages.
	m.SetAliasNbPages("{nbs}")

	// This will create a row with full width column and inside it a text that will display the total number of pages.
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Total number of pages: {nbs}")
		})
	})
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

	// This Row will be filled with the color.
	m.Row(20, func() {
		m.Col(12, func() {
			// Add components.
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	// Note: The default value is White (255, 255, 255), if maroto see this value it will ignore not will the cell with any color.

	// This Row will not be filled with the color.
	m.Row(20, func() {
		m.Col(12, func() {
			// Add components.
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_GetBorder demonstrates how to
// obtain the actual borders status.
func ExamplePdfMaroto_GetBorder() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// false.
	_ = m.GetBorder()

	m.SetBorder(true)

	// true.
	_ = m.GetBorder()

	// Do more things and save...
}

// ExamplePdfMaroto_GetPageSize demonstrates how to obtain the current page size (width and height).
func ExamplePdfMaroto_GetPageSize() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Get
	width, height := m.GetPageSize()
	fmt.Println(width)
	fmt.Println(height)

	// Do more things and save...
}

// ExamplePdfMaroto_GetCurrentPage demonstrates how to obtain the current page index.
func ExamplePdfMaroto_GetCurrentPage() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Index here will be 0.
	_ = m.GetCurrentPage()

	// Add Rows, Cols and Components.

	// Index here will not be 0.
	_ = m.GetCurrentPage()

	// Do more things and save...
}

// ExamplePdfMaroto_GetCurrentOffset demonstrates how to obtain the current write offset
// i.e the height of cursor adding content in the pdf.
func ExamplePdfMaroto_GetCurrentOffset() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Offset here will be 0.
	_ = m.GetCurrentOffset()

	// Add Rows, Cols and Components until maroto add a new page.

	// Offset here will not be 0.
	_ = m.GetCurrentOffset()

	// Add Rows, Cols and Components to maroto add a new page.

	// Offset here will be 0.
	_ = m.GetCurrentOffset()

	// Do more things and save...
}

// ExamplePdfMaroto_SetPageMargins demonstrates how to set custom page margins.
func ExamplePdfMaroto_SetPageMargins() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetPageMargins(10, 60, 10)

	// Do more things or not and save...
}

// ExamplePdfMaroto_GetPageMargins demonstrates how to obtain the current page margins.
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

// ExamplePdfMaroto_AddUTF8Font demonstrates how to add a custom utf8 font.
func ExamplePdfMaroto_AddUTF8Font() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Define font to all styles.
	m.AddUTF8Font("CustomArial", consts.Normal, "internal/assets/fonts/arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.Italic, "internal/assets/fonts/arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.Bold, "internal/assets/fonts/arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.BoldItalic, "internal/assets/fonts/arial-unicode-ms.ttf")

	m.Row(10, func() {
		m.Col(12, func() {
			// Use style.
			m.Text("CustomUtf8Font", props.Text{
				Family: "CustomArial",
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_SetFontLocation demonstrates how to add a custom utf8 font from custom location.
func ExamplePdfMaroto_SetFontLocation() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Define custom location. It might be an absolute path as well as a relative path.
	m.SetFontLocation("internal/assets/fonts/")

	// Define font to all styles.
	m.AddUTF8Font("CustomArial", consts.Normal, "arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.Italic, "arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.Bold, "arial-unicode-ms.ttf")
	m.AddUTF8Font("CustomArial", consts.BoldItalic, "arial-unicode-ms.ttf")

	m.Row(10, func() {
		m.Col(12, func() {
			// Use style.
			m.Text("CustomUtf8Font", props.Text{
				Family: "CustomArial",
			})
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_AddUTF8Font demonstrates how to define a custom font to your pdf.
func TestPdfMaroto_SetDefaultFontFamily(t *testing.T) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Default font family is Arial.

	// Change the default to Courier.
	m.SetDefaultFontFamily(consts.Courier)

	m.Row(10, func() {
		m.Col(12, func() {
			// This will be in courier.
			m.Text("CustomUtf8Font")
		})
	})

	// Do more things and save...
}

// ExamplePdfMaroto_AddUTF8Font demonstrates how to obtain the current default font family.
func ExamplePdfMaroto_GetDefaultFontFamily() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Will return Arial.
	_ = m.GetDefaultFontFamily()

	// Change the default to Courier.
	m.SetDefaultFontFamily(consts.Courier)

	// Will return Courier.
	_ = m.GetDefaultFontFamily()

	// Do more things and save...
}

// ExamplePdfMaroto_SetProtection demonstrates how to define a protection to pdf.
func ExamplePdfMaroto_SetProtection() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetProtection(1, "userPassword", "ownerPassword")

	// Do more things and save...
}

// ExamplePdfMaroto_SetCompression demonstrates how to disable compression
// By default compression is enabled.
func ExamplePdfMaroto_SetCompression() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetCompression(false)

	// Do more things and save...
}
