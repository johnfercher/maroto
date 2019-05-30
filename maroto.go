package maroto

import (
	"bytes"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

type Maroto interface {
	// Grid System
	Row(label string, height float64, closure func())
	Col(label string, closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Features
	RowTableList(label string, header []string, contents [][]string)
	SetDebugMode(on bool)
	Text(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align HorizontalAlign)
	Image(filePathName string, marginTop float64)
	Barcode(code string, width float64, height float64, marginTop float64) error
	QrCode(code string)
	Sign(label string, fontFamily Family, fontStyle Style, fontSize float64)
	Line()

	// File System
	OutputFileAndClose(filePathName string) error
	Output() (bytes.Buffer, error)
}

type maroto struct {
	fpdf         gofpdf.Pdf
	math         Math
	font         Font
	text         Text
	sign         Sign
	image        Image
	offsetY      float64
	rowHeight    float64
	rowColCount  float64
	colsClosures []func()
	debugMode    bool
}

func (m *maroto) Output() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := m.fpdf.Output(&buffer)
	return buffer, err
}

func NewMaroto(orientation Orientation, pageSize PageSize) Maroto {
	fpdfOrientation := "P"
	fpdfPageSize := "A4"

	if orientation == Vertical {
		fpdfOrientation = "P"
	}

	if pageSize == A4 {
		fpdfPageSize = "A4"
	}

	fpdf := gofpdf.New(fpdfOrientation, "mm", fpdfPageSize, "")
	fpdf.SetMargins(10, 10, 10)

	_math := NewMath(fpdf)
	_font := NewFont(fpdf, 16, Arial, Bold)
	_text := NewText(fpdf, _math, _font)

	_sign := NewSign(fpdf, _math, _text)

	_image := NewImage(fpdf, _math)

	maroto := &maroto{
		fpdf:  fpdf,
		math:  _math,
		font:  _font,
		text:  _text,
		sign:  _sign,
		image: _image,
	}

	maroto.font.SetFamily(Arial)
	maroto.font.SetStyle(Bold)
	maroto.font.SetSize(16)
	maroto.debugMode = false

	maroto.fpdf.AddPage()

	return maroto
}

func (m *maroto) Sign(label string, fontFamily Family, fontStyle Style, fontSize float64) {
	qtdCols := float64(len(m.colsClosures))
	sumOfYOffsets := m.offsetY + m.rowHeight

	m.sign.Sign(label, fontFamily, fontStyle, fontSize, qtdCols, sumOfYOffsets, m.rowColCount)
}

func (m *maroto) RowTableList(label string, header []string, contents [][]string) {
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

func (m *maroto) SetDebugMode(on bool) {
	m.debugMode = on
}

func (m *maroto) Line() {
	m.Row("", 1, func() {
		m.Col("", func() {
			width, _ := m.fpdf.GetPageSize()
			left, top, right, _ := m.fpdf.GetMargins()

			m.fpdf.Line(left, m.offsetY+top, width-right, m.offsetY+top)
		})
	})
}

func (m *maroto) Row(label string, height float64, closure func()) {
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

func (m *maroto) Col(label string, closure func()) {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		closure()
		m.rowColCount++
	})

}

func (m *maroto) ColSpace() {
	m.colsClosures = append(m.colsClosures, func() {
		widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
		m.createColSpace(widthPerCol)
		m.rowColCount++
	})
}

func (m *maroto) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		m.colsClosures = append(m.colsClosures, func() {
			widthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))
			m.createColSpace(widthPerCol)
			m.rowColCount++
		})
	}
}

func (m *maroto) Text(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align HorizontalAlign) {
	if marginTop > m.rowHeight {
		marginTop = m.rowHeight
	}

	sumOfYOffsets := marginTop + m.offsetY

	m.text.Add(text, fontFamily, fontStyle, fontSize, sumOfYOffsets, align, m.rowColCount, float64(len(m.colsClosures)))
}

func (m *maroto) Image(filePathName string, marginTop float64) {
	qtdCols := float64(len(m.colsClosures))
	sumOfyOffsets := m.offsetY + marginTop

	m.image.AddFromPath(filePathName, sumOfyOffsets, m.rowColCount, qtdCols, m.rowHeight)
}
func (m *maroto) OutputFileAndClose(filePathName string) (err error) {
	err = m.fpdf.OutputFileAndClose(filePathName)
	return
}

func (m *maroto) Barcode(code string, width float64, height float64, marginTop float64) (err error) {
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

func (m *maroto) QrCode(code string) {
	key := barcode.RegisterQR(m.fpdf, code, qr.H, qr.Unicode)

	actualWidthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))

	qrSide := actualWidthPerCol

	if m.rowHeight < qrSide {
		qrSide = m.rowHeight
	}

	barcode.Barcode(m.fpdf, key, actualWidthPerCol*m.rowColCount, 0, qrSide, qrSide, false)
}

func (m *maroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if m.debugMode {
		border = "1"
	}

	m.fpdf.CellFormat(actualWidthPerCol, m.rowHeight, "", border, 0, "C", false, 0, "")
}
