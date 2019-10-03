package pdf

import (
	"bytes"
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
func (s *PdfMaroto) RegisterHeader(closure func()) {
	s.headerClosure = closure
}

// RegisterFooter define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (s *PdfMaroto) RegisterFooter(closure func()) {
	s.footerClosure = closure

	// calculation mode execute all row flow but
	// only to calculate the sum of heights
	s.calculationMode = true
	closure()
	s.calculationMode = false
}

// GetCurrentPage obtain the current page index
// this can be used inside a RegisterFooter/RegisterHeader
// to draw the current page, or to another purposes
func (s *PdfMaroto) GetCurrentPage() int {
	return s.pageIndex
}

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (s *PdfMaroto) Signature(label string, prop ...props.Font) {
	signProp := props.Font{}
	if len(prop) > 0 {
		signProp = prop[0]
	}

	signProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfYOffsets := s.offsetY + s.rowHeight

	s.SignHelper.AddSpaceFor(label, signProp.ToTextProp(consts.Center, 0.0), qtdCols, sumOfYOffsets, s.rowColCount)
}

// TableList create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (s *PdfMaroto) TableList(header []string, contents [][]string, prop ...props.TableList) {
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

	s.Row(tableProp.HeaderHeight, func() {
		headerMarginTop := 2.0
		qtdCols := float64(len(header))

		for i, h := range header {
			hs := h
			is := i

			s.Col(func() {
				if headerMarginTop > s.rowHeight {
					headerMarginTop = s.rowHeight
				}

				reason := hs

				sumOyYOffesets := headerMarginTop + s.offsetY + 2.5

				s.TextHelper.Add(reason, tableProp.HeaderProp.ToTextProp(tableProp.Align, 0.0), sumOyYOffesets, float64(is), qtdCols)
			})
		}
	})

	s.Row(tableProp.HeaderContentSpace, func() {
		s.ColSpace()
	})

	contentMarginTop := 2.0

	for _, content := range contents {
		s.Row(tableProp.ContentHeight, func() {
			for j, c := range content {
				cs := c
				js := j
				hs := float64(len(header))
				sumOyYOffesets := contentMarginTop + s.offsetY + 2.0

				s.Col(func() {
					s.TextHelper.Add(cs, tableProp.ContentProp.ToTextProp(tableProp.Align, 0.0), sumOyYOffesets, float64(js), hs)
				})
			}
		})
	}
}

// SetBorder enable the draw of lines in every cell.
// Draw borders in all columns created.
func (s *PdfMaroto) SetBorder(on bool) {
	s.DebugMode = on
}

// GetBorder return the actual border value.
func (s *PdfMaroto) GetBorder() bool {
	return s.DebugMode
}

// GetPageSize return the actual page size
func (s *PdfMaroto) GetPageSize() (float64, float64) {
	return s.Pdf.GetPageSize()
}

// Line draw a line from margin left to margin right
// in the currently row.
func (s *PdfMaroto) Line(spaceHeight float64) {
	s.Row(spaceHeight, func() {
		s.Col(func() {
			width, _ := s.Pdf.GetPageSize()
			left, top, right, _ := s.Pdf.GetMargins()

			s.Pdf.Line(left, s.offsetY+top+(spaceHeight/2.0), width-right, s.offsetY+top+(spaceHeight/2.0))
		})
	})
}

// Row define a row and enable add columns inside the row.
func (s *PdfMaroto) Row(height float64, closure func()) {
	// Used to calculate the height of the footer
	if s.calculationMode {
		s.footerHeight += height
		return
	}

	_, pageHeight := s.Pdf.GetPageSize()
	_, top, _, bottom := s.Pdf.GetMargins()

	totalOffsetY := int(s.offsetY + height + s.footerHeight)
	maxOffsetPage := int(pageHeight - bottom - top)

	// Note: The headerFooterContextActive is needed to avoid recursive
	// calls without end, because footerClosure and headerClosure actually
	// have Row calls too.

	// If the new cell to be added pass the useful space counting the
	// height of the footer, add the footer
	if totalOffsetY > maxOffsetPage {
		if !s.headerFooterContextActive {
			if s.footerClosure != nil {
				s.headerFooterContextActive = true
				s.footerClosure()
				s.headerFooterContextActive = false
			}
			s.offsetY = 0
			s.pageIndex++
		}
	}

	// If is a new page, add the header
	if !s.headerFooterContextActive && s.headerClosure != nil {
		if s.offsetY == 0 {
			s.headerFooterContextActive = true
			s.headerClosure()
			s.headerFooterContextActive = false
		}
	}

	s.rowHeight = height
	s.rowColCount = 0

	// This closure has only maroto.Cols, which are
	// not executed firstly, they are added to colsClosures
	// and this enable us to know how many cols will be added
	// and calculate the width from the cells
	closure()

	// Execute the codes inside the Cols
	for _, colClosure := range s.colsClosures {
		colClosure()
	}

	s.colsClosures = nil
	s.offsetY += s.rowHeight
	s.Pdf.Ln(s.rowHeight)
}

// Col create a column inside a row and enable to add
// components inside.
func (s *PdfMaroto) Col(closure func()) {
	s.colsClosures = append(s.colsClosures, func() {
		widthPerCol := s.Math.GetWidthPerCol(float64(len(s.colsClosures)))
		s.createColSpace(widthPerCol)
		closure()
		s.rowColCount++
	})
}

// ColSpace create an empty column inside a row.
func (s *PdfMaroto) ColSpace() {
	s.colsClosures = append(s.colsClosures, func() {
		widthPerCol := s.Math.GetWidthPerCol(float64(len(s.colsClosures)))
		s.createColSpace(widthPerCol)
		s.rowColCount++
	})
}

// ColSpaces create some empty columns inside a row.
func (s *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		s.ColSpace()
	}
}

// Text create a text inside a cell.
func (s *PdfMaroto) Text(text string, prop ...props.Text) {
	textProp := props.Text{}
	if len(prop) > 0 {
		textProp = prop[0]
	}

	textProp.MakeValid()

	if textProp.Top > s.rowHeight {
		textProp.Top = s.rowHeight
	}

	sumOfYOffsets := textProp.Top + s.offsetY

	s.TextHelper.Add(text, textProp, sumOfYOffsets, s.rowColCount, float64(len(s.colsClosures)))
}

// FileImage add an Image reading from disk inside a cell.
// Defining Image properties.
func (s *PdfMaroto) FileImage(filePathName string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top

	if rectProp.Center {
		s.Image.AddFromFile(filePathName, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent)
	} else {
		s.Image.AddFromFile(filePathName, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent)
	}
}

// Base64Image add an Image reading byte slices inside a cell.
// Defining Image properties.
func (s *PdfMaroto) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top

	if rectProp.Center {
		s.Image.AddFromBase64(base64, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent, extension)
	} else {
		s.Image.AddFromBase64(base64, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent, extension)
	}
}

// OutputFileAndClose save pdf in disk.
func (s *PdfMaroto) OutputFileAndClose(filePathName string) (err error) {
	s.drawLastFooter()
	err = s.Pdf.OutputFileAndClose(filePathName)

	return
}

// Output extract PDF in byte slices
func (s *PdfMaroto) Output() (bytes.Buffer, error) {
	s.drawLastFooter()
	var buffer bytes.Buffer
	err := s.Pdf.Output(&buffer)
	return buffer, err
}

// Barcode create an barcode inside a cell.
func (s *PdfMaroto) Barcode(code string, prop ...props.Barcode) (err error) {
	barcodeProp := props.Barcode{}
	if len(prop) > 0 {
		barcodeProp = prop[0]
	}

	barcodeProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + barcodeProp.Top

	if barcodeProp.Center {
		err = s.Code.AddBar(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, barcodeProp.Percent, barcodeProp.Proportion.Height/barcodeProp.Proportion.Width)
	} else {
		err = s.Code.AddBar(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, barcodeProp.Percent, barcodeProp.Proportion.Height/barcodeProp.Proportion.Width)
	}

	return
}

// QrCode create a qrcode inside a cell.
func (s *PdfMaroto) QrCode(code string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top

	if rectProp.Center {
		s.Code.AddQr(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent)
	} else {
		s.Code.AddQr(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp.Percent)
	}
}

func (s *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if s.DebugMode {
		border = "1"
	}

	s.Pdf.CellFormat(actualWidthPerCol, s.rowHeight, "", border, 0.0, "C", false, 0.0, "")
}

func (s *PdfMaroto) drawLastFooter() {
	if s.footerClosure != nil {
		_, pageHeight := s.Pdf.GetPageSize()
		_, top, _, bottom := s.Pdf.GetMargins()

		if s.offsetY+s.footerHeight < pageHeight-bottom-top {
			s.headerFooterContextActive = true
			s.footerClosure()
			s.headerFooterContextActive = false
		}
	}
}
