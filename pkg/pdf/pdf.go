package pdf

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(height float64, closure func())
	Col(closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Registers
	RegisterHeader(closure func())
	RegisterFooter(closure func())

	// Helpers
	SetBorder(on bool)
	GetBorder() bool
	GetPageSize() (float64, float64)
	GetCurrentLine() float64
	GetCurrentPage() int

	// Outside Col/Row Components
	TableList(header []string, contents [][]string, prop ...props.TableList)
	Line(spaceHeight float64)

	// Inside Col/Row Components
	Text(text string, prop ...props.Text)
	FileImage(filePathName string, prop ...props.Rect)
	Base64Image(base64 string, extension consts.Extension, prop ...props.Rect)
	Barcode(code string, prop ...props.Barcode) error
	QrCode(code string, prop ...props.Rect)
	Signature(label string, prop ...props.Font)

	// File System
	OutputFileAndClose(filePathName string) error
	Output() (bytes.Buffer, error)
}

// PdfMaroto is the principal structure which implements Maroto abstraction
type PdfMaroto struct {
	Pdf                       gofpdf.Pdf
	Math                      internal.Math
	Font                      internal.Font
	TextHelper                internal.Text
	SignHelper                internal.Signature
	Image                     internal.Image
	Code                      internal.Code
	pageIndex                 int
	offsetY                   float64
	rowHeight                 float64
	rowColCount               float64
	colsClosures              []func()
	headerClosure             func()
	footerClosure             func()
	footerHeight              float64
	headerFooterContextActive bool
	calculationMode           bool
	DebugMode                 bool
	orientation               consts.Orientation
	pageSize                  consts.PageSize
}

// NewMaroto create a Maroto instance returning a pointer to PdfMaroto
// Receive an Orientation and a PageSize.
func NewMaroto(orientation consts.Orientation, pageSize consts.PageSize) Maroto {
	fpdf := gofpdf.New(string(orientation), "mm", string(pageSize), "")
	fpdf.SetMargins(10, 10, 10)

	math := internal.NewMath(fpdf)
	font := internal.NewFont(fpdf, 16, consts.Arial, consts.Bold)
	text := internal.NewText(fpdf, math, font)

	signature := internal.NewSignature(fpdf, math, text)

	image := internal.NewImage(fpdf, math)

	code := internal.NewCode(fpdf, math)

	maroto := &PdfMaroto{
		Pdf:             fpdf,
		Math:            math,
		Font:            font,
		TextHelper:      text,
		SignHelper:      signature,
		Image:           image,
		Code:            code,
		pageSize:        pageSize,
		orientation:     orientation,
		calculationMode: false,
	}

	maroto.Font.SetFamily(consts.Arial)
	maroto.Font.SetStyle(consts.Bold)
	maroto.Font.SetSize(16)
	maroto.DebugMode = false

	maroto.Pdf.AddPage()

	return maroto
}

// RegisterHeader define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (self *PdfMaroto) RegisterHeader(closure func()) {
	self.headerClosure = closure
}

// RegisterFooter define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (self *PdfMaroto) RegisterFooter(closure func()) {
	self.footerClosure = closure
	self.calculationMode = true
	closure()
	fmt.Println(self.footerHeight)
	self.calculationMode = false
}

// GetCurrentLine return the current line i.e
// the current y offset given the origin in the
// pdf page
func (self *PdfMaroto) GetCurrentLine() float64 {
	return self.offsetY
}

func (self *PdfMaroto) GetCurrentPage() int {
	return self.pageIndex
}

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (self *PdfMaroto) Signature(label string, prop ...props.Font) {
	signProp := props.Font{}
	if len(prop) > 0 {
		signProp = prop[0]
	}

	signProp.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfYOffsets := self.offsetY + self.rowHeight

	self.SignHelper.AddSpaceFor(label, signProp.ToTextProp(consts.Center, 0.0), qtdCols, sumOfYOffsets, self.rowColCount)
}

// TableList create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (self *PdfMaroto) TableList(header []string, contents [][]string, prop ...props.TableList) {
	if len(header) == 0 {
		return
	}

	if len(contents) == 0 {
		return
	}

	tableProp := props.TableList{}
	if len(prop) > 0 {
		tableProp = prop[0]
	}

	tableProp.MakeValid()

	self.Row(tableProp.HeaderHeight, func() {
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

				self.TextHelper.Add(reason, tableProp.HeaderProp.ToTextProp(tableProp.Align, 0.0), sumOyYOffesets, float64(is), qtdCols)
			})
		}
	})

	self.Row(tableProp.HeaderContentSpace, func() {
		self.ColSpace()
	})

	contentMarginTop := 2.0

	for _, content := range contents {
		self.Row(tableProp.ContentHeight, func() {
			for j, c := range content {
				cs := c
				js := j
				hs := float64(len(header))
				sumOyYOffesets := contentMarginTop + self.offsetY + 2.0

				self.Col(func() {
					self.TextHelper.Add(cs, tableProp.ContentProp.ToTextProp(tableProp.Align, 0.0), sumOyYOffesets, float64(js), hs)
				})
			}
		})
	}
}

// SetBorder enable the draw of lines in every cell.
// Draw borders in all columns created.
func (self *PdfMaroto) SetBorder(on bool) {
	self.DebugMode = on
}

// GetBorder return the actual border value.
func (self *PdfMaroto) GetBorder() bool {
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
	if self.calculationMode {
		self.footerHeight += height
		return
	}

	_, pageHeight := self.Pdf.GetPageSize()
	_, top, _, bottom := self.Pdf.GetMargins()

	if self.offsetY+height+self.footerHeight > pageHeight-bottom-top {
		if !self.headerFooterContextActive && self.footerClosure != nil {
			self.headerFooterContextActive = true
			self.footerClosure()
			self.headerFooterContextActive = false
			self.offsetY = 0
		}
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
func (self *PdfMaroto) Text(text string, prop ...props.Text) {
	textProp := props.Text{}
	if len(prop) > 0 {
		textProp = prop[0]
	}

	textProp.MakeValid()

	if textProp.Top > self.rowHeight {
		textProp.Top = self.rowHeight
	}

	sumOfYOffsets := textProp.Top + self.offsetY

	self.TextHelper.Add(text, textProp, sumOfYOffsets, self.rowColCount, float64(len(self.colsClosures)))
}

// FileImage add an Image reading from disk inside a cell.
// Defining Image properties.
func (self *PdfMaroto) FileImage(filePathName string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + rectProp.Top

	if rectProp.Center {
		self.Image.AddFromFile(filePathName, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent)
	} else {
		self.Image.AddFromFile(filePathName, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent)
	}
}

// Base64Image add an Image reading byte slices inside a cell.
// Defining Image properties.
func (self *PdfMaroto) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + rectProp.Top

	if rectProp.Center {
		self.Image.AddFromBase64(base64, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent, extension)
	} else {
		self.Image.AddFromBase64(base64, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent, extension)
	}
}

// OutputFileAndClose save pdf in disk.
func (self *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	self.drawLastFooter()
	err = self.Pdf.OutputFileAndClose(filePathName)

	return
}

// Output extract PDF in byte slices
func (self *PdfMaroto) Output() (bytes.Buffer, error) {
	self.drawLastFooter()
	var buffer bytes.Buffer
	err := self.Pdf.Output(&buffer)
	return buffer, err
}

// Barcode create an barcode inside a cell.
func (self *PdfMaroto) Barcode(code string, prop ...props.Barcode) (err error) {
	barcodeProp := props.Barcode{}
	if len(prop) > 0 {
		barcodeProp = prop[0]
	}

	barcodeProp.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + barcodeProp.Top

	if barcodeProp.Center {
		err = self.Code.AddBar(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, barcodeProp.Percent, barcodeProp.Proportion.Height/barcodeProp.Proportion.Width)
	} else {
		err = self.Code.AddBar(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, barcodeProp.Percent, barcodeProp.Proportion.Height/barcodeProp.Proportion.Width)
	}

	return
}

// QrCode create a qrcode inside a cell.
func (self *PdfMaroto) QrCode(code string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(self.colsClosures))
	sumOfyOffsets := self.offsetY + rectProp.Top

	if rectProp.Center {
		self.Code.AddQr(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent)
	} else {
		self.Code.AddQr(code, sumOfyOffsets, self.rowColCount, qtdCols, self.rowHeight, rectProp.Percent)
	}
}

func (self *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if self.DebugMode {
		border = "1"
	}

	self.Pdf.CellFormat(actualWidthPerCol, self.rowHeight, "", border, 0.0, "C", false, 0.0, "")
}

func (self *PdfMaroto) drawLastFooter() {
	if self.footerClosure != nil {
		_, pageHeight := self.Pdf.GetPageSize()
		_, top, _, bottom := self.Pdf.GetMargins()

		if self.offsetY+self.footerHeight < pageHeight-bottom-top {
			self.headerFooterContextActive = true
			self.footerClosure()
			self.headerFooterContextActive = false
		}
	}
}

func (self *PdfMaroto) footerHeightCalcRow(height float64) {
	self.footerHeight += height
}
