package pdf

import (
	"bytes"
	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Maroto is the principal abstraction to create a PDF document.
type Maroto interface {
	// Grid System
	Row(height float64, closure func())
	Col(width uint, closure func())
	ColSpace(uint)
	ColSpaces(qtd int)

	// Registers
	RegisterHeader(closure func())
	RegisterFooter(closure func())

	// Helpers
	SetBorder(on bool)
	SetBackgroundColor(color color.Color)
	GetBorder() bool
	GetPageSize() (width float64, height float64)
	GetCurrentPage() int
	GetCurrentOffset() float64
	SetPageMargins(left, top, right float64)
	GetPageMargins() (left float64, top float64, right float64, bottom float64)

	// Outside Col/Row Components
	TableList(header []string, contents [][]string, prop ...props.TableList)
	Line(spaceHeight float64)

	// Inside Col/Row Components
	Text(text string, prop ...props.Text)
	FileImage(filePathName string, prop ...props.Rect) (err error)
	Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) (err error)
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
	TableListHelper           internal.TableList
	pageIndex                 int
	offsetY                   float64
	marginTop                 float64
	rowHeight                 float64
	xColOffset                float64
	colWidth                  float64
	backgroundColor           color.Color
	colsClosures              []func()
	headerClosure             func()
	footerClosure             func()
	footerHeight              float64
	headerFooterContextActive bool
	calculationMode           bool
	debugMode                 bool
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

	tableList := internal.NewTableList(text, font)

	maroto := &PdfMaroto{
		Pdf:             fpdf,
		Math:            math,
		Font:            font,
		TextHelper:      text,
		SignHelper:      signature,
		Image:           image,
		Code:            code,
		TableListHelper: tableList,
		pageSize:        pageSize,
		orientation:     orientation,
		calculationMode: false,
		backgroundColor: color.NewWhite(),
	}

	maroto.TableListHelper.BindGrid(maroto)

	maroto.Font.SetFamily(consts.Arial)
	maroto.Font.SetStyle(consts.Bold)
	maroto.Font.SetSize(16)
	maroto.debugMode = false

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

// GetCurrentOffset obtain the current offset in y axis
func (s *PdfMaroto) GetCurrentOffset() float64 {
	return s.offsetY
}

// SetPageMargins overrides default margins (10,10,10)
// the new page margin will affect all PDF pages
func (s *PdfMaroto) SetPageMargins(left, top, right float64) {
	if top <= 10 {
		s.Pdf.SetMargins(left, top, right)
	} else {
		s.marginTop = top - 10
		s.Pdf.SetMargins(left, 10, right)
	}
}

// GetPageMargins returns the set page margins. Comes in order of Left, Top, Right, Bottom
// Default page margins is left: 10, top: 10, right: 10
func (s *PdfMaroto) GetPageMargins() (left float64, top float64, right float64, bottom float64) {
	return s.Pdf.GetMargins()
}

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (s *PdfMaroto) Signature(label string, prop ...props.Font) {
	signProp := props.Font{}
	if len(prop) > 0 {
		signProp = prop[0]
	}

	signProp.MakeValid()

	yColOffset := s.offsetY + s.rowHeight

	s.SignHelper.AddSpaceFor(label, signProp.ToTextProp(consts.Center, 0.0, false, 0), s.colWidth, yColOffset, s.xColOffset)
}

// TableList create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (s *PdfMaroto) TableList(header []string, contents [][]string, prop ...props.TableList) {
	s.TableListHelper.Create(header, contents, prop...)
}

// SetBorder enable the draw of lines in every cell.
// Draw borders in all columns created.
func (s *PdfMaroto) SetBorder(on bool) {
	s.debugMode = on
}

// SetBackgroundColor define the background color of the PDF.
// This method can be used to toggle background from rows
func (s *PdfMaroto) SetBackgroundColor(color color.Color) {
	s.backgroundColor = color
	s.Pdf.SetFillColor(s.backgroundColor.Red, s.backgroundColor.Green, s.backgroundColor.Blue)
}

// GetBorder return the actual border value.
func (s *PdfMaroto) GetBorder() bool {
	return s.debugMode
}

// GetPageSize return the actual page size
func (s *PdfMaroto) GetPageSize() (width float64, height float64) {
	return s.Pdf.GetPageSize()
}

// Line draw a line from margin left to margin right
// in the currently row.
func (s *PdfMaroto) Line(spaceHeight float64) {
	s.Row(spaceHeight, func() {
		s.Col(0, func() {
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
			s.headerFooterContextActive = true
			s.footer()
			s.headerFooterContextActive = false
			s.offsetY = 0
			s.pageIndex++
		}
	}

	// If is a new page, add the header
	if !s.headerFooterContextActive {
		if s.offsetY == 0 {
			s.headerFooterContextActive = true
			s.header()
			s.headerFooterContextActive = false
		}
	}

	s.rowHeight = height
	s.xColOffset = 0

	// This closure has only maroto.Cols, which are
	// not executed firstly, they are added to colsClosures
	// and this enable us to know how many cols will be added
	// and calculate the width from the cells
	closure()

	// Execute the codes inside the Cols
	//for _, colClosure := range s.colsClosures {
	//	colClosure()
	//}

	s.colsClosures = nil
	s.offsetY += s.rowHeight
	s.Pdf.Ln(s.rowHeight)
}

// Col create a column inside a row and enable to add
// components inside.
func (s *PdfMaroto) Col(width uint, closure func()) {
	// Array will be executed only in the Row context
	//s.colsClosures = append(s.colsClosures, func() {
	if width == 0 {
		width = 12
	}

	percent := float64(width) / float64(12)

	widthPerCol := s.Math.GetWidthPerCol(percent)
	s.colWidth = widthPerCol
	s.createColSpace(widthPerCol)
	closure()
	s.xColOffset += s.colWidth
	//})
}

// ColSpace create an empty column inside a row.
func (s *PdfMaroto) ColSpace(width uint) {
	s.colsClosures = append(s.colsClosures, func() {
		widthPerCol := s.Math.GetWidthPerCol(float64(len(s.colsClosures)))
		s.createColSpace(widthPerCol)
		s.xColOffset++
	})
}

// ColSpaces create some empty columns inside a row.
func (s *PdfMaroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		s.ColSpace(0)
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

	yColOffset := s.offsetY + textProp.Top

	s.TextHelper.Add(text, textProp, yColOffset, s.xColOffset, s.colWidth)
}

// FileImage add an Image reading from disk inside a cell.
// Defining Image properties.
func (s *PdfMaroto) FileImage(filePathName string, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	sumOfyOffsets := s.offsetY + rectProp.Top

	return s.Image.AddFromFile(filePathName, sumOfyOffsets, s.xColOffset, s.colWidth, s.rowHeight, rectProp)
}

// Base64Image add an Image reading byte slices inside a cell.
// Defining Image properties.
func (s *PdfMaroto) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	sumOfyOffsets := s.offsetY + rectProp.Top

	return s.Image.AddFromBase64(base64, sumOfyOffsets, s.xColOffset, s.colWidth, s.rowHeight, rectProp, extension)
}

// Barcode create an barcode inside a cell.
func (s *PdfMaroto) Barcode(code string, prop ...props.Barcode) (err error) {
	barcodeProp := props.Barcode{}
	if len(prop) > 0 {
		barcodeProp = prop[0]
	}

	barcodeProp.MakeValid()

	sumOfyOffsets := s.offsetY + barcodeProp.Top

	err = s.Code.AddBar(code, sumOfyOffsets, s.xColOffset, s.colWidth, s.rowHeight, barcodeProp)

	return
}

// QrCode create a qrcode inside a cell.
func (s *PdfMaroto) QrCode(code string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	yColOffset := s.offsetY + rectProp.Top

	s.Code.AddQr(code, yColOffset, s.xColOffset, s.colWidth, s.rowHeight, rectProp)
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

func (s *PdfMaroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if s.debugMode {
		border = "1"
	}

	s.Pdf.CellFormat(actualWidthPerCol, s.rowHeight, "", border, 0.0, "C", !s.backgroundColor.IsWhite(), 0.0, "")
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

func (s *PdfMaroto) footer() {
	backgroundColor := s.backgroundColor
	s.SetBackgroundColor(color.NewWhite())

	_, pageHeight := s.Pdf.GetPageSize()
	_, top, _, bottom := s.Pdf.GetMargins()

	totalOffsetY := int(s.offsetY + s.footerHeight)
	maxOffsetPage := int(pageHeight - bottom - top)

	s.Row(float64(maxOffsetPage-totalOffsetY), func() {
		s.ColSpace(12)
	})

	if s.footerClosure != nil {
		s.footerClosure()
	}

	s.SetBackgroundColor(backgroundColor)
}

func (s *PdfMaroto) header() {
	backgroundColor := s.backgroundColor
	s.SetBackgroundColor(color.NewWhite())

	s.Row(s.marginTop, func() {
		s.ColSpace(12)
	})

	if s.headerClosure != nil {
		s.headerClosure()
	}

	s.SetBackgroundColor(backgroundColor)
}
