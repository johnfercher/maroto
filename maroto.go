package maroto

import (
	"bytes"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(label string, height float64, closure func())
	Col(label string, closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Components
	RowTableList(label string, header []string, contents [][]string)
	SetDebugMode(on bool)
	Text(text string, marginTop float64, fontProp *FontProp)
	FileImage(filePathName string, rectProp *RectProp)
	Base64Image(base64 string, extension Extension, rectProp *RectProp)
	Barcode(code string, width float64, height float64, marginTop float64) error
	QrCode(code string)
	Signature(label string, signatureProp *SignatureProp)
	Line()

	// File System
	OutputFileAndClose(filePathName string) error
	Output() (bytes.Buffer, error)
}

// PdfMaroto is the principal structure which implements Maroto abstraction
type PdfMaroto struct {
	fpdf         gofpdf.Pdf
	math         Math
	font         Font
	text         Text
	signature    Signature
	image        Image
	offsetY      float64
	rowHeight    float64
	rowColCount  float64
	colsClosures []func()
	debugMode    bool
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

	maroto := &PdfMaroto{
		fpdf:      fpdf,
		math:      math,
		font:      font,
		text:      text,
		signature: signature,
		image:     image,
	}

	maroto.font.SetFamily(Arial)
	maroto.font.SetStyle(Bold)
	maroto.font.SetSize(16)
	maroto.debugMode = false

	maroto.fpdf.AddPage()

	return maroto
}

// Add a signature space with a label text inside a column.
// Create a line with the width from a column
// and add a text at the bottom of the line.
func (m *PdfMaroto) Signature(label string, signatureProp *SignatureProp) {
	if signatureProp == nil {
		signatureProp = &SignatureProp{}
	}

	signatureProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfYOffsets := m.offsetY + m.rowHeight

	m.signature.AddSpaceFor(label, signatureProp.Family, signatureProp.Style, signatureProp.Size, qtdCols, sumOfYOffsets, m.rowColCount)
}

func (m *PdfMaroto) RowTableList(label string, header []string, contents [][]string) {
	headerHeight := 7.0

	m.Row("", headerHeight, func() {
		headerMarginTop := 2.0
		qtdCols := float64(len(header))

		for i, h := range header {
			hs := h
			is := i

			m.Col("", func() {
				if headerMarginTop > m.rowHeight {
					headerMarginTop = m.rowHeight
				}

				reason := hs

				sumOyYOffesets := headerMarginTop + m.offsetY + 2.5

				m.text.Add(reason, Arial, Bold, 10, sumOyYOffesets, Left, float64(is), qtdCols)
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
				hs := float64(len(header))
				sumOyYOffesets := contentMarginTop + m.offsetY + 2.0

				m.Col("", func() {
					m.text.Add(cs, Arial, Normal, 10, sumOyYOffesets, Left, float64(js), hs)
				})
			}
		})
	}
}

// Enable debug mode.
// Draw borders in all columns created.
func (m *PdfMaroto) SetDebugMode(on bool) {
	m.debugMode = on
}

// Draw a line from margin left to margin right
// in the currently row.
func (m *PdfMaroto) Line() {
	m.Row("", 1, func() {
		m.Col("", func() {
			width, _ := m.fpdf.GetPageSize()
			left, top, right, _ := m.fpdf.GetMargins()

			m.fpdf.Line(left, m.offsetY+top, width-right, m.offsetY+top)
		})
	})
}

// Add a row and enable add columns inside the row.
func (m *PdfMaroto) Row(label string, height float64, closure func()) {
	m.rowHeight = height
	m.rowColCount = 0

	_, pageHeight := m.fpdf.GetPageSize()
	_, top, _, bottom := m.fpdf.GetMargins()

	if m.offsetY > pageHeight-bottom-top-m.rowHeight {
		m.offsetY = 0
	}

	closure()

	for _, colClosure := range m.colsClosures {
		colClosure()
	}

	m.colsClosures = nil
	m.offsetY += m.rowHeight
	m.fpdf.Ln(m.rowHeight)
}

// Create a column inside a row and enable to add
// components inside.
func (m *PdfMaroto) Col(label string, closure func()) {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		closure()
		m.rowColCount++
	})

}

// Create an empty column inside a row.
func (m *PdfMaroto) ColSpace() {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		m.rowColCount++
	})
}

// Create some empty columns.
func (m *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		m.colsClosures = append(m.colsClosures, func() {
			widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
			m.createColSpace(widthPerCol)
			m.rowColCount++
		})
	}
}

// Add a text inside a column.
func (m *PdfMaroto) Text(text string, marginTop float64, fontProp *FontProp) {
	if fontProp == nil {
		fontProp = &FontProp{}
	}

	fontProp.MakeValid()

	if marginTop > m.rowHeight {
		marginTop = m.rowHeight
	}

	sumOfYOffsets := marginTop + m.offsetY

	m.text.Add(text, fontProp.Family, fontProp.Style, fontProp.Size, sumOfYOffsets, fontProp.Align, m.rowColCount, float64(len(m.colsClosures)))
}

// Add an image reading from disk inside a column.
// Defining image properties.
func (m *PdfMaroto) FileImage(filePathName string, rectProp *RectProp) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))

	if rectProp.Center {
		m.image.AddFromFile(filePathName, m.offsetY, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	} else {
		m.image.AddFromFile(filePathName, m.offsetY, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent)
	}
}

// Add an image reading byte slices.
// Defining image properties.
func (m *PdfMaroto) Base64Image(base64 string, extension Extension, rectProp *RectProp) {
	if rectProp == nil {
		rectProp = &RectProp{}
	}

	rectProp.MakeValid()

	qtdCols := float64(len(m.colsClosures))
	sumOfyOffsets := m.offsetY + rectProp.Top

	if rectProp.Center {
		m.image.AddFromBase64(base64, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent, extension)
	} else {
		m.image.AddFromBase64(base64, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight, rectProp.Percent, extension)
	}
}

// Save pdf in disk.
func (m *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	err = m.fpdf.OutputFileAndClose(filePathName)
	return
}

// Get PDF in byte slices
func (m *PdfMaroto) Output() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := m.fpdf.Output(&buffer)
	return buffer, err
}

func (m *PdfMaroto) Barcode(code string, width float64, height float64, marginTop float64) (err error) {
	bcode, err := code128.Encode(code)

	if err != nil {
		return
	}

	actualWidthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))

	if width > actualWidthPerCol {
		width = actualWidthPerCol
	}

	if height > m.rowHeight {
		height = m.rowHeight
	}

	left, top, _, _ := m.fpdf.GetMargins()

	sumOffsetY := marginTop + top + m.offsetY

	barcode.Barcode(m.fpdf, barcode.Register(bcode), actualWidthPerCol*m.rowColCount+((actualWidthPerCol-width)/2)+left, sumOffsetY, width, height, false)
	return
}

func (m *PdfMaroto) QrCode(code string) {
	key := barcode.RegisterQR(m.fpdf, code, qr.H, qr.Unicode)

	actualWidthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))

	qrSide := actualWidthPerCol

	if m.rowHeight < qrSide {
		qrSide = m.rowHeight
	}

	barcode.Barcode(m.fpdf, key, actualWidthPerCol*m.rowColCount, 0, qrSide, qrSide, false)
}

func (m *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if m.debugMode {
		border = "1"
	}

	m.fpdf.CellFormat(actualWidthPerCol, m.rowHeight, "", border, 0, "C", false, 0, "")
}
