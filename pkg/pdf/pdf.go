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
	ColSpace(gridSize uint)

	// Registers
	RegisterHeader(closure func())
	RegisterFooter(closure func())

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

	// Helpers
	AddPage()
	SetBorder(on bool)
	SetBackgroundColor(color color.Color)
	GetBorder() bool
	GetPageSize() (width float64, height float64)
	GetCurrentPage() int
	GetCurrentOffset() float64
	SetPageMargins(left, top, right float64)
	GetPageMargins() (left float64, top float64, right float64, bottom float64)
	AddUTF8Font(familyStr, styleStr, fileStr string)
	SetProtection(actionFlag byte, userPassStr, ownerPassStr string)
}

// PdfMaroto is the principal structure which implements Maroto abstraction
type PdfMaroto struct {
	Pdf                       *gofpdf.Fpdf
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
	headerClosure             func()
	footerClosure             func()
	footerHeight              float64
	headerFooterContextActive bool
	calculationMode           bool
	debugMode                 bool
	orientation               consts.Orientation
	pageSize                  consts.PageSize
}

// NewMarotoCustomSize creates a Maroto instance returning a pointer to PdfMaroto
// Receive an Orientation and a PageSize.
// Use if custom page size is needed. Otherwise use NewMaroto() shorthand if using page sizes from consts.Pagesize.
// If using custom width and height, pageSize is just a string value for the format and takes no effect.
// Width and height inputs are measurements of the page in Portrait orientation.
func NewMarotoCustomSize(orientation consts.Orientation, pageSize consts.PageSize, unitStr string, width, height float64) Maroto {
	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: string(orientation),
		UnitStr:        unitStr,
		SizeStr:        string(pageSize),
		Size: gofpdf.SizeType{
			Wd: width,
			Ht: height,
		},
		FontDirStr: "",
	})
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

// NewMaroto create a Maroto instance returning a pointer to PdfMaroto
// Receive an Orientation and a PageSize.
// Shorthand when using a preset page size from consts.PageSize
func NewMaroto(orientation consts.Orientation, pageSize consts.PageSize) Maroto {
	return NewMarotoCustomSize(orientation, pageSize, "mm", 0, 0)
}

// AddPage adds a new page in the PDF
func (s *PdfMaroto) AddPage() {
	_, pageHeight := s.Pdf.GetPageSize()
	_, top, _, bottom := s.Pdf.GetMargins()

	totalOffsetY := int(s.offsetY + s.footerHeight)
	maxOffsetPage := int(pageHeight - bottom - top)

	s.Row(float64(maxOffsetPage-totalOffsetY), func() {
		s.ColSpace(12)
	})
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
	if top > 10 {
		s.marginTop = top - 10
	}

	s.Pdf.SetMargins(left, 10.0, right)
}

// GetPageMargins returns the set page margins. Comes in order of Left, Top, Right, Bottom
// Default page margins is left: 10, top: 10, right: 10
func (s *PdfMaroto) GetPageMargins() (left float64, top float64, right float64, bottom float64) {
	left, top, right, bottom = s.Pdf.GetMargins()
	top += s.marginTop

	return
}

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (s *PdfMaroto) Signature(label string, prop ...props.Font) {
	signProp := props.Font{}
	if len(prop) > 0 {
		signProp = prop[0]
	}

	signProp.MakeValid()

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY,
		Width:  s.colWidth,
		Height: s.rowHeight,
	}

	s.SignHelper.AddSpaceFor(label, cell, signProp.ToTextProp(consts.Center, 0.0, false, 0))
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

	// This closure has the Cols to be executed
	closure()

	s.offsetY += s.rowHeight
	s.Pdf.Ln(s.rowHeight)
}

// Col create a column inside a row and enable to add
// components inside.
func (s *PdfMaroto) Col(width uint, closure func()) {
	if width == 0 {
		width = 12
	}

	percent := float64(width) / float64(12)

	pageWidth, _ := s.Pdf.GetPageSize()
	left, _, right, _ := s.Pdf.GetMargins()
	widthPerCol := (pageWidth - right - left) * percent

	s.colWidth = widthPerCol
	s.createColSpace(widthPerCol)

	// This closure has the components to be executed
	closure()

	s.xColOffset += s.colWidth
}

// ColSpace create an empty column inside a row.
func (s *PdfMaroto) ColSpace(gridSize uint) {
	s.Col(gridSize, func() {})
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

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY + textProp.Top,
		Width:  s.colWidth,
		Height: 0,
	}

	s.TextHelper.Add(text, cell, textProp)
}

// FileImage add an Image reading from disk inside a cell.
// Defining Image properties.
func (s *PdfMaroto) FileImage(filePathName string, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY + rectProp.Top,
		Width:  s.colWidth,
		Height: s.rowHeight,
	}

	return s.Image.AddFromFile(filePathName, cell, rectProp)
}

// Base64Image add an Image reading byte slices inside a cell.
// Defining Image properties.
func (s *PdfMaroto) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY + rectProp.Top,
		Width:  s.colWidth,
		Height: s.rowHeight,
	}

	return s.Image.AddFromBase64(base64, cell, rectProp, extension)
}

// Barcode create an barcode inside a cell.
func (s *PdfMaroto) Barcode(code string, prop ...props.Barcode) (err error) {
	barcodeProp := props.Barcode{}
	if len(prop) > 0 {
		barcodeProp = prop[0]
	}

	barcodeProp.MakeValid()

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY + barcodeProp.Top,
		Width:  s.colWidth,
		Height: s.rowHeight,
	}

	err = s.Code.AddBar(code, cell, barcodeProp)

	return
}

// QrCode create a qrcode inside a cell.
func (s *PdfMaroto) QrCode(code string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	cell := internal.Cell{
		X:      s.xColOffset,
		Y:      s.offsetY + rectProp.Top,
		Width:  s.colWidth,
		Height: s.rowHeight,
	}

	s.Code.AddQr(code, cell, rectProp)
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

	s.Pdf.CellFormat(actualWidthPerCol, s.rowHeight, "", border, 0, "C", !s.backgroundColor.IsWhite(), 0, "")
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
func (s *PdfMaroto) AddUTF8Font(familyStr, styleStr, fileStr string) {
	s.Pdf.AddUTF8Font(familyStr, styleStr, fileStr)
}

func (s *PdfMaroto) SetProtection(actionFlag byte, userPassStr, ownerPassStr string) {
	s.Pdf.SetProtection(actionFlag, userPassStr, ownerPassStr)
}
