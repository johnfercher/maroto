package maroto

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/johnfercher/maroto/enums"
	"github.com/johnfercher/maroto/font"
	"github.com/johnfercher/maroto/math"
	"github.com/johnfercher/maroto/sign"
	"github.com/johnfercher/maroto/text"
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
	Text(text string, fontFamily font.Family, fontStyle font.Style, fontSize float64, marginTop float64, align enums.HorizontalAlign)
	Image(filePathName string, marginTop float64)
	Barcode(code string, width float64, height float64, marginTop float64) error
	QrCode(code string)
	Sign(label string, fontFamily font.Family, fontStyle font.Style, fontSize float64)
	Line()

	// File System
	OutputFileAndClose(filePathName string) error
}

type maroto struct {
	fpdf         gofpdf.Pdf
	math         math.Math
	font         font.Font
	text         text.Text
	sign         sign.Sign
	offsetY      float64
	rowHeight    float64
	rowColCount  float64
	colsClosures []func()
	debugMode    bool
}

func NewMaroto(orientation enums.Orientation, pageSize enums.PageSize) Maroto {
	fpdfOrientation := "P"
	fpdfPageSize := "A4"

	if orientation == enums.Vertical {
		fpdfOrientation = "P"
	}

	if pageSize == enums.A4 {
		fpdfPageSize = "A4"
	}

	fpdf := gofpdf.New(fpdfOrientation, "mm", fpdfPageSize, "")
	fpdf.SetMargins(10, 10, 10)

	_math := math.NewMath(fpdf)
	_font := font.NewFont(fpdf, 16, font.Arial, font.Bold)
	_text := text.NewText(fpdf, _math, _font)

	_sign := sign.NewSign(fpdf, _math, _text)

	maroto := &maroto{
		fpdf: fpdf,
		math: _math,
		font: _font,
		text: _text,
		sign: _sign,
	}

	maroto.font.SetFamily(font.Arial)
	maroto.font.SetStyle(font.Bold)
	maroto.font.SetSize(16)
	maroto.debugMode = false

	return maroto
}

func (m *maroto) Sign(label string, fontFamily font.Family, fontStyle font.Style, fontSize float64) {
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

				m.text.Add(reason, font.Arial, font.Bold, 10, sumOyYOffesets, enums.Left, float64(is), qtdCols)
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
					m.text.Add(cs, font.Arial, font.Normal, 10, sumOyYOffesets, enums.Left, float64(js), hs)
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
	m.fpdf.Ln(m.rowHeight)
	m.rowHeight = height
	m.rowColCount = 0

	if m.fpdf.PageCount() == 0 {
		m.fpdf.AddPage()
	}

	closure()

	for _, colClosure := range m.colsClosures {
		colClosure()
	}

	m.colsClosures = nil
	m.offsetY += m.rowHeight
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

func (m *maroto) Text(text string, fontFamily font.Family, fontStyle font.Style, fontSize float64, marginTop float64, align enums.HorizontalAlign) {
	if marginTop > m.rowHeight {
		marginTop = m.rowHeight
	}

	sumOfYOffsets := marginTop + m.offsetY

	m.text.Add(text, fontFamily, fontStyle, fontSize, sumOfYOffsets, align, m.rowColCount, float64(len(m.colsClosures)))
}

func (m *maroto) Image(filePathName string, marginTop float64) {
	var opt gofpdf.ImageOptions
	actualWidthPerCol := m.math.GetWidthPerCol(float64(len(m.colsClosures)))

	left, top, _, _ := m.fpdf.GetMargins()

	m.fpdf.ImageOptions(filePathName, actualWidthPerCol*m.rowColCount+left, m.offsetY+marginTop+top, actualWidthPerCol, 0, false, opt, 0, "")
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

	barcode.Barcode(m.fpdf, barcode.Register(bcode), actualWidthPerCol*m.rowColCount+((actualWidthPerCol-width)/2)+left, marginTop+top, width, height, false)
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

	return
}

func (m *maroto) createColSpace(actualWidthPerCol float64) {
	border := ""

	if m.debugMode {
		border = "1"
	}

	m.fpdf.CellFormat(actualWidthPerCol, m.rowHeight, "", border, 0, "C", false, 0, "")
}
