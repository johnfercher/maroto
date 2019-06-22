package maroto

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
)

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(height float64, closure func())
	RegisterHeader(closure func())
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
	Pdf                       gofpdf.Pdf
	Math                      Math
	Font                      Font
	TextHelper                Text
	SignHelper                Signature
	Image                     Image
	Code                      Code
	offsetY                   float64
	rowHeight                 float64
	rowColCount               float64
	colsClosures              []func()
	headerClosure             func()
	headerFooterContextActive bool
	DebugMode                 bool
}

// RegisterHeader define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (self *PdfMaroto) RegisterHeader(closure func()) {
	self.headerClosure = closure
}

// NewMaroto create a Maroto instance returning a pointer to PdfMaroto
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

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (self *PdfMaroto) Signature(label string, prop *SignatureProp) {
	if prop == nil {
		prop = &SignatureProp{}
	}

	prop.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfYOffsets := self.offsetY + self.rowHeight

	self.SignHelper.AddSpaceFor(label, prop.Family, prop.Style, prop.Size, qtdCols, sumOfYOffsets, self.rowColCount)
}

// TableList create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (self *PdfMaroto) TableList(header []string, contents [][]string, prop *TableListProp) {
	if len(header) == 0 {
		return
	}

	if len(contents) == 0 {
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

// SetDebugMode enable debug mode.
// Draw borders in all columns created.
func (self *PdfMaroto) SetDebugMode(on bool) {
	self.DebugMode = on
}

// GetDebugMode return the actual debug mode.
func (self *PdfMaroto) GetDebugMode() bool {
	return self.DebugMode
}

// GetPageSize return the actual page size
func (self *PdfMaroto) GetPageSize() (float64, float64) {
	return self.Pdf.GetPageSize()
}

// Line draw a line from margin left to margin right
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

// Row define a row and enable add columns inside the row.
func (self *PdfMaroto) Row(height float64, closure func()) {
	_, pageHeight := self.Pdf.GetPageSize()
	_, top, _, bottom := self.Pdf.GetMargins()

	if self.offsetY+height > pageHeight-bottom-top {
		self.offsetY = 0
	}

	if !self.headerFooterContextActive && self.headerClosure != nil {
		if self.offsetY == 0 {
			self.headerFooterContextActive = true
			self.headerClosure()
			self.headerFooterContextActive = false
		}
	}

	self.rowHeight = height
	self.rowColCount = 0

	closure()

	for _, colClosure := range self.colsClosures {
		colClosure()
	}

	self.colsClosures = nil
	self.offsetY += self.rowHeight
	self.Pdf.Ln(self.rowHeight)
}

// Col create a column inside a row and enable to add
// components inside.
func (self *PdfMaroto) Col(closure func()) {
	self.colsClosures = append(self.colsClosures, func() {
		widthPerCol := self.Math.GetWidthPerCol(float64(len(self.colsClosures)))
		self.createColSpace(widthPerCol)
		closure()
		self.rowColCount++
	})
}

// ColSpace create an empty column inside a row.
func (self *PdfMaroto) ColSpace() {
	self.colsClosures = append(self.colsClosures, func() {
		widthPerCol := self.Math.GetWidthPerCol(float64(len(self.colsClosures)))
		self.createColSpace(widthPerCol)
		self.rowColCount++
	})
}

// ColSpace create some empty columns inside a row.
func (self *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		self.ColSpace()
	}
}

// Text create a text inside a cell.
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

// FileImage add an Image reading from disk inside a cell.
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

// Base64Image add an Image reading byte slices inside a cell.
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

// OutputFileAndClose save pdf in disk.
func (self *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	err = self.Pdf.OutputFileAndClose(filePathName)
	return
}

// Output extract PDF in byte slices
func (self *PdfMaroto) Output() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := self.Pdf.Output(&buffer)
	return buffer, err
}

// Barcode create an barcode inside a cell.
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

// QrCode create a qrcode inside a cell.
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
