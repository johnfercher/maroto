package maroto

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
)

const tableListTag = "header"

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(height float64, closure func())
	Col(closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Helpers
	SetDebugMode(on bool)
	GetDebugMode() bool
	GetPageSize() (float64, float64)

	// Outside Col/Row Components
	TableList(header []string, contents [][]string, prop *TableListProp)
	Line(spaceHeight float64)

	// Inside Col/Row Components
	Text(text string, prop *TextProp)
	FileImage(filePathName string, prop *RectProp)
	Base64Image(base64 string, extension Extension, prop *RectProp)
	Barcode(code string, prop *RectProp) error
	QrCode(code string, prop *RectProp)
	Signature(label string, prop *SignatureProp)

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
func (self *PdfMaroto) Signature(label string, prop *SignatureProp) {
	if prop == nil {
		prop = &SignatureProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfYOffsets := self.offsetY + self.rowHeight

	self.SignHelper.AddSpaceFor(label, prop.Family, prop.Style, prop.Size, qtdCols, sumOfYOffsets, self.rowColCount)
}

// Create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (self *PdfMaroto) TableList(header []string, contents [][]string, prop *TableListProp) {
	if header == nil || len(header) == 0 {
		return
	}

	if contents == nil || len(contents) == 0 {
		return
	}

	if prop == nil {
		prop = &TableListProp{}
	}

	prop.MakeValid()

	self.Row(prop.HHeight, func() {
		headerMarginTop := 2.0
		qtdCols := float64(len(header))

		for i, h := range header {
			hs := h
			is := i

			self.Col(func() {
				if headerMarginTop > self.rowHeight {
					headerMarginTop = self.rowHeight
				}

				reason := hs

				sumOyYOffesets := headerMarginTop + self.offsetY + 2.5

				self.TextHelper.Add(reason, prop.HFontFamily, prop.HFontStyle, prop.HFontSize, sumOyYOffesets, prop.Align, float64(is), qtdCols)
			})
		}
	})

	self.Row(prop.Space, func() {
		self.ColSpace()
	})

	contentMarginTop := 2.0

	for _, content := range contents {
		self.Row(prop.CHeight, func() {
			for j, c := range content {
				cs := c
				js := j
				hs := float64(len(header))
				sumOyYOffesets := contentMarginTop + self.offsetY + 2.0

				self.Col(func() {
					self.TextHelper.Add(cs, prop.CFontFamily, prop.CFontStyle, prop.CFontSize, sumOyYOffesets, prop.Align, float64(js), hs)
				})
			}
		})
	}
}

// Enable debug mode.
// Draw borders in all columns created.
func (self *PdfMaroto) SetDebugMode(on bool) {
	self.DebugMode = on
}

// Get actual debug mode.
func (self *PdfMaroto) GetDebugMode() bool {
	return self.DebugMode
}

// Get actual page size
func (self *PdfMaroto) GetPageSize() (float64, float64) {
	return self.Pdf.GetPageSize()
}

// Draw a line from margin left to margin right
// in the currently row.
func (self *PdfMaroto) Line(spaceHeight float64) {
	self.Row(spaceHeight, func() {
		self.Col(func() {
			width, _ := self.Pdf.GetPageSize()
			left, top, right, _ := self.Pdf.GetMargins()

			self.Pdf.Line(left, self.offsetY+top+(spaceHeight/2.0), width-right, self.offsetY+top+(spaceHeight/2.0))
		})
	})
}

// Add a row and enable add columns inside the row.
func (self *PdfMaroto) Row(height float64, closure func()) {
	self.rowHeight = height
	self.rowColCount = 0

	_, pageHeight := self.Pdf.GetPageSize()
	_, top, _, bottom := self.Pdf.GetMargins()

	if self.offsetY > pageHeight-bottom-top-self.rowHeight {
		self.offsetY = 0
	}

	closure()

	for _, colClosure := range self.colsClosures {
		colClosure()
	}

	self.colsClosures = nil
	self.offsetY += self.rowHeight
	self.Pdf.Ln(self.rowHeight)
}

// Create a column inside a row and enable to add
// components inside.
func (self *PdfMaroto) Col(closure func()) {
	self.colsClosures = append(self.colsClosures, func() {
		widthPerCol := self.Math.GetWidthPerCol(float64(len(self.colsClosures)))
		self.createColSpace(widthPerCol)
		closure()
		self.rowColCount++
	})
}

// Create an empty column inside a row.
func (self *PdfMaroto) ColSpace() {
	self.colsClosures = append(self.colsClosures, func() {
		widthPerCol := self.Math.GetWidthPerCol(float64(len(self.colsClosures)))
		self.createColSpace(widthPerCol)
		self.rowColCount++
	})
}

// Create some empty columns.
func (self *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		self.ColSpace()
	}
}

// Add a Text inside a column.
func (self *PdfMaroto) Text(text string, prop *TextProp) {
	if prop == nil {
		prop = &TextProp{}
	}

	prop.MakeValid()

	if prop.Top > self.rowHeight {
		prop.Top = self.rowHeight
	}

	sumOfYOffsets := prop.Top + self.offsetY

	self.TextHelper.Add(text, prop.Family, prop.Style, prop.Size, sumOfYOffsets, prop.Align, self.rowColCount, float64(len(self.colsClosures)))
}

// Add an Image reading from disk inside a column.
// Defining Image properties.
func (self *PdfMaroto) FileImage(filePathName string, prop *RectProp) {
	if prop == nil {
		prop = &RectProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + prop.Top

	if prop.Center {
		self.Image.AddFromFile(filePathName, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	} else {
		self.Image.AddFromFile(filePathName, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	}
}

// Add an Image reading byte slices.
// Defining Image properties.
func (self *PdfMaroto) Base64Image(base64 string, extension Extension, prop *RectProp) {
	if prop == nil {
		prop = &RectProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + prop.Top

	if prop.Center {
		self.Image.AddFromBase64(base64, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent, extension)
	} else {
		self.Image.AddFromBase64(base64, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent, extension)
	}
}

// Save pdf in disk.
func (self *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	err = self.Pdf.OutputFileAndClose(filePathName)
	return
}

// Get PDF in byte slices
func (self *PdfMaroto) Output() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := self.Pdf.Output(&buffer)
	return buffer, err
}

func (self *PdfMaroto) Barcode(code string, prop *RectProp) (err error) {
	if prop == nil {
		prop = &RectProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + prop.Top

	if prop.Center {
		err = self.Code.AddBar(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	} else {
		err = self.Code.AddBar(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	}

	return
}

func (self *PdfMaroto) QrCode(code string, prop *RectProp) {
	if prop == nil {
		prop = &RectProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + prop.Top

	if prop.Center {
		self.Code.AddQr(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	} else {
		self.Code.AddQr(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, prop.Percent)
	}
}

func (self *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if self.DebugMode {
		border = "1"
	}

	self.Pdf.CellFormat(actualWidthPerCol, self.rowHeight, "", border, 0.0, "C", false, 0.0, "")
}
