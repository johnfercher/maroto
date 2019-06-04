package maroto

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
)

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(label string, height float64, closure func())
	Col(label string, closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Helpers
	SetDebugMode(on bool)
	GetDebugMode() bool
	GetPageSize() (float64, float64)

	// Components
	RowTableList(label string, headers []string, contents [][]string)
	Text(text string, fontProp *TextProp)
	FileImage(filePathName string, rectProp *RectProp)
	Base64Image(base64 string, extension Extension, rectProp *RectProp)
	Barcode(code string, rectProp *RectProp) error
	QrCode(code string, rectProp *RectProp)
	Signature(label string, signatureProp *SignatureProp)
	Line()

	// File System
	OutputFileAndClose(filePathName string) error
	Output() (bytes.Buffer, error)
}

// PdfMaroto is the principal structure which implements Maroto abstraction
type PdfMaroto struct {
	Pdf          gofpdf.Pdf
	Math         Math
	Font         Font
	TextHelper   Text
	SignHelper   Signature
	Image        Image
	Code         Code
	offsetY      float64
	rowHeight    float64
	rowColCount  float64
	colsClosures []func()
	DebugMode    bool
}

// Create a Maroto instance returning a pointer to PdfMaroto
// Receive an Orientation and a PageSize.
func NewMaroto(orientation Orientation, pageSize PageSize) Maroto {
	fpdf := gofpdf.New(string(orientation), "mm", string(pageSize), "")
	fpdf.SetMargins(10, 10, 10)

	math := NewMath(fpdf)
	font := NewFont(fpdf, 16, Arial, Bold)
	text := NewText(fpdf, math, font)

	signature := NewSignature(fpdf, math, text)

	image := NewImage(fpdf, math)

	code := NewCode(fpdf, math)

	maroto := &PdfMaroto{
		Pdf:        fpdf,
		Math:       math,
		Font:       font,
		TextHelper: text,
		SignHelper: signature,
		Image:      image,
		Code:       code,
	}

	maroto.Font.SetFamily(Arial)
	maroto.Font.SetStyle(Bold)
	maroto.Font.SetSize(16)
	maroto.DebugMode = false

	maroto.Pdf.AddPage()

	return maroto
}

// Add a Signature space with a label TextHelper inside a column.
// Create a line with the width from a column
// and add a Text at the bottom of the line.
func (m *PdfMaroto) Signature(label string, signatureProp *SignatureProp) {
	if signatureProp == nil {
		signatureProp = &SignatureProp{}
	}

	signatureProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfYOffsets := m.offsetY + m.rowHeight

	m.SignHelper.AddSpaceFor(label, signatureProp.Family, signatureProp.Style, signatureProp.Size, qtdCols, sumOfYOffsets, m.rowColCount)
}

// Create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (m *PdfMaroto) RowTableList(label string, headers []string, contents [][]string) {
	headerHeight := 7.0

	m.Row("", headerHeight, func() {
		headerMarginTop := 2.0
		qtdCols := float64(len(headers))

		for i, h := range headers {
			hs := h
			is := i

			m.Col("", func() {
				if headerMarginTop > m.rowHeight {
					headerMarginTop = m.rowHeight
				}

				reason := hs

				sumOyYOffesets := headerMarginTop + m.offsetY + 2.5

				m.TextHelper.Add(reason, Arial, Bold, 10, sumOyYOffesets, Left, float64(is), qtdCols)
			})
		}
	})

	m.Row("", 4.0, func() {
		m.ColSpace()
	})

	contentHeight := 5.0
	contentMarginTop := 2.0

	for _, content := range contents {
		m.Row("", contentHeight, func() {
			for j, c := range content {
				cs := c
				js := j
				hs := float64(len(headers))
				sumOyYOffesets := contentMarginTop + m.offsetY + 2.0

				m.Col("", func() {
					m.TextHelper.Add(cs, Arial, Normal, 10, sumOyYOffesets, Left, float64(js), hs)
				})
			}
		})
	}
}

// Enable debug mode.
// Draw borders in all columns created.
func (m *PdfMaroto) SetDebugMode(on bool) {
	m.DebugMode = on
}

// Get actual debug mode.
func (m *PdfMaroto) GetDebugMode() bool {
	return m.DebugMode
}

// Get actual page size
func (m *PdfMaroto) GetPageSize() (float64, float64) {
	return m.Pdf.GetPageSize()
}

// Draw a line from margin left to margin right
// in the currently row.
func (m *PdfMaroto) Line() {
	m.Row("", 1, func() {
		m.Col("", func() {
			width, _ := m.Pdf.GetPageSize()
			left, top, right, _ := m.Pdf.GetMargins()

			m.Pdf.Line(left, m.offsetY+top, width-right, m.offsetY+top)
		})
	})
}

// Add a row and enable add columns inside the row.
func (m *PdfMaroto) Row(label string, height float64, closure func()) {
	m.rowHeight = height
	m.rowColCount = 0

	_, pageHeight := m.Pdf.GetPageSize()
	_, top, _, bottom := m.Pdf.GetMargins()

	if m.offsetY > pageHeight-bottom-top-m.rowHeight {
		m.offsetY = 0
	}

	closure()

	for _, colClosure := range m.colsClosures {
		colClosure()
	}

	m.colsClosures = nil
	m.offsetY += m.rowHeight
	m.Pdf.Ln(m.rowHeight)
}

// Create a column inside a row and enable to add
// components inside.
func (m *PdfMaroto) Col(label string, closure func()) {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.Math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		closure()
		m.rowColCount++
	})

}

// Create an empty column inside a row.
func (m *PdfMaroto) ColSpace() {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.Math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		m.rowColCount++
	})
}

// Create some empty columns.
func (m *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		m.ColSpace()
	}
}

// Add a Text inside a column.
func (m *PdfMaroto) Text(text string, fontProp *TextProp) {
	if fontProp == nil {
		fontProp = &TextProp{}
	}

	fontProp.MakeValid()

	if fontProp.Top > m.rowHeight {
		fontProp.Top = m.rowHeight
	}

	sumOfYOffsets := fontProp.Top + m.offsetY

	m.TextHelper.Add(text, fontProp.Family, fontProp.Style, fontProp.Size, sumOfYOffsets, fontProp.Align, m.rowColCount, float64(len(m.colsClosures)))
}

// Add an Image reading from disk inside a column.
// Defining Image properties.
func (m *PdfMaroto) FileImage(filePathName string, rectProp *RectProp) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))

	if rectProp.Center {
		m.Image.AddFromFile(filePathName, m.offsetY, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	} else {
		m.Image.AddFromFile(filePathName, m.offsetY, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	}
}

// Add an Image reading byte slices.
// Defining Image properties.
func (m *PdfMaroto) Base64Image(base64 string, extension Extension, rectProp *RectProp) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfyOffsets := m.offsetY + rectProp.Top

	if rectProp.Center {
		m.Image.AddFromBase64(base64, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent, extension)
	} else {
		m.Image.AddFromBase64(base64, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent, extension)
	}
}

// Save pdf in disk.
func (m *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	err = m.Pdf.OutputFileAndClose(filePathName)
	return
}

// Get PDF in byte slices
func (m *PdfMaroto) Output() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := m.Pdf.Output(&buffer)
	return buffer, err
}

func (m *PdfMaroto) Barcode(code string, rectProp *RectProp) (err error) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfyOffsets := m.offsetY + rectProp.Top

	err = m.Code.AddBar(code, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)

	return
}

func (m *PdfMaroto) QrCode(code string, rectProp *RectProp) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfyOffsets := m.offsetY + rectProp.Top

	if rectProp.Center {
		m.Code.AddQr(code, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	} else {
		m.Code.AddQr(code, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	}
}

func (m *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if m.DebugMode {
		border = "1"
	}

	m.Pdf.CellFormat(actualWidthPerCol, m.rowHeight, "", border, 0, "C", false, 0, "")
}
