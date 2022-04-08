package pdf_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPdf(t *testing.T) {
	cases := []struct {
		name        string
		orientation consts.Orientation
		pageSize    consts.PageSize
		assert      func(t *testing.T, m pdf.Maroto)
	}{
		{
			"When portrait and A4",
			consts.Portrait,
			consts.A4,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 210.0, 0.1)
				assert.InDelta(t, height, 297.0, 0.1)
			},
		},
		{
			"When portrait and A3",
			consts.Portrait,
			consts.A3,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 297.0, 0.1)
				assert.InDelta(t, height, 419.9, 0.1)
			},
		},
		{
			"When portrait and A5",
			consts.Portrait,
			consts.A5,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 148.4, 0.1)
				assert.InDelta(t, height, 210.0, 0.1)
			},
		},
		{
			"When portrait and Legal",
			consts.Portrait,
			consts.Legal,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 215.9, 0.1)
				assert.InDelta(t, height, 355.6, 0.1)
			},
		},
		{
			"When portrait and Letter",
			consts.Portrait,
			consts.Letter,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 215.9, 0.1)
				assert.InDelta(t, height, 279.4, 0.1)
			},
		},
		{
			"When landscape and A4",
			consts.Landscape,
			consts.A4,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 210.0, 0.1)
				assert.InDelta(t, width, 297.0, 0.1)
			},
		},
		{
			"When landscape and A3",
			consts.Landscape,
			consts.A3,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 297.0, 0.1)
				assert.InDelta(t, width, 419.9, 0.1)
			},
		},
		{
			"When landscape and A5",
			consts.Landscape,
			consts.A5,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 148.4, 0.1)
				assert.InDelta(t, width, 210.0, 0.1)
			},
		},
		{
			"When landscape and Legal",
			consts.Landscape,
			consts.Legal,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 215.9, 0.1)
				assert.InDelta(t, width, 355.6, 0.1)
			},
		},
		{
			"When landscape and Letter",
			consts.Landscape,
			consts.Letter,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 215.9, 0.1)
				assert.InDelta(t, width, 279.4, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Act
		m := pdf.NewMaroto(c.orientation, c.pageSize)

		// Assert
		c.assert(t, m)
	}
}

func TestNewCustomSizeFpdf(t *testing.T) {
	cases := []struct {
		name        string
		orientation consts.Orientation
		pageSize    consts.PageSize
		unit        string
		height      float64
		width       float64
		assert      func(t *testing.T, m pdf.Maroto)
	}{
		{
			"When portrait and C6",
			consts.Portrait,
			"C6",
			"mm",
			162.0,
			114.0,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 114.0, 0.1)
				assert.InDelta(t, height, 162.0, 0.1)
			},
		},
		{
			"When landscape and C6",
			consts.Landscape,
			"C6",
			"mm",
			162.0,
			114.0,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 162.0, 0.1)
				assert.InDelta(t, height, 114.0, 0.1)
			},
		},
		{
			"When portrait and B3 and given in cm",
			consts.Portrait,
			"B3",
			"cm",
			50.0,
			35.3,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 35.3, 0.1)
				assert.InDelta(t, height, 50.0, 0.1)
			},
		},
		{
			"When landdscape and B3 and given in cm",
			consts.Landscape,
			"B3",
			"cm",
			50.0,
			35.3,
			func(t *testing.T, m pdf.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*pdf.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 50.0, 0.1)
				assert.InDelta(t, height, 35.3, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Act
		m := pdf.NewMarotoCustomSize(c.orientation, c.pageSize, c.unit, c.width, c.height)

		// Assert
		c.assert(t, m)
	}
}

func TestFpdfMaroto_SetGetDebugMode(t *testing.T) {
	// Arrange
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Assert & Act
	assert.False(t, m.GetBorder())
	m.SetBorder(true)

	// Assert
	assert.True(t, m.GetBorder())
}

func TestFpdfMaroto_SetFirstPageNb(t *testing.T) {
	// Arrange
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Assert & Act
	assert.Equal(t, 0, m.GetCurrentPage())
	m.SetFirstPageNb(1)

	// Assert
	assert.Equal(t, 1, m.GetCurrentPage())
}

func TestFpdfMaroto_SetAliasNbPages(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.SetAliasNbPages("{nbs}")

	// Assert
	Fpdf.AssertCalled(t, "AliasNbPages", "{nbs}")
}

func TestFpdfMaroto_Signature(t *testing.T) {
	cases := []struct {
		name      string
		signature func() *mocks.Signature
		assert    func(t *testing.T, signature *mocks.Signature)
		act       func(m pdf.Maroto)
	}{
		{
			"Calculate mode",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNotCalled(t, "AddSpaceFor")
			},
			func(m pdf.Maroto) {
				m.RegisterFooter(func() {
					m.Row(40, func() {
						m.Col(12, func() {
							m.Signature("Signature1")
						})
					})
				})
			},
		},
		{
			"One signature inside one column, inside a row, without props",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 1)
				signature.AssertCalled(t, "AddSpaceFor", "Signature1", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Arial,
					Style:  consts.Bold,
					Size:   8.0,
					Align:  consts.Center,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(0, func() {
						m.Signature("Signature1")
					})
				})
			},
		},
		{
			"Two different signatures inside one colum, inside one row",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 2)
				signature.AssertCalled(t, "AddSpaceFor", "Signature2", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Arial,
					Style:  consts.Bold,
					Size:   8.0,
					Align:  consts.Center,
				})
				signature.AssertCalled(t, "AddSpaceFor", "Signature3", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9.5,
					Align:  consts.Center,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Signature("Signature2")
						m.Signature("Signature3", props.Font{
							Family: consts.Courier,
							Style:  consts.BoldItalic,
							Size:   9.5,
						})
					})
				})
			},
		},
		{
			"Two different signatures with different columns, inside one row",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 2)
				signature.AssertCalled(t, "AddSpaceFor", "Signature4", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  40.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Arial,
					Style:  consts.Bold,
					Size:   8.0,
					Align:  consts.Center,
				})
				signature.AssertCalled(t, "AddSpaceFor", "Signature5", internal.Cell{
					X:      40.0,
					Y:      0.0,
					Width:  40.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9.5,
					Align:  consts.Center,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(6, func() {
						m.Signature("Signature4")
					})
					m.Col(6, func() {
						m.Signature("Signature5", props.Font{
							Family: consts.Courier,
							Style:  consts.BoldItalic,
							Size:   9.5,
						})
					})
				})
			},
		},
		{
			"Two different signatures with different columns, inside one row",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 2)
				signature.AssertCalled(t, "AddSpaceFor", "Signature6", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  40.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Arial,
					Style:  consts.Bold,
					Size:   8.0,
					Align:  consts.Center,
				})
				signature.AssertCalled(t, "AddSpaceFor", "Signature7", internal.Cell{
					X:      0.0,
					Y:      40.0,
					Width:  40.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9.5,
					Align:  consts.Center,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(6, func() {
						m.Signature("Signature6")
					})
				})
				m.Row(40, func() {
					m.Col(6, func() {
						m.Signature("Signature7", props.Font{
							Family: consts.Courier,
							Style:  consts.BoldItalic,
							Size:   9.5,
						})
					})
				})
			},
		},
		{
			"Custom color signature",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 1)
				signature.AssertCalled(t, "AddSpaceFor", "Signature1", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 40.0,
				}, props.Text{
					Family: consts.Arial,
					Style:  consts.Bold,
					Size:   8.0,
					Align:  consts.Center,
					Color: color.Color{
						Red:   20,
						Green: 20,
						Blue:  20,
					},
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(0, func() {
						m.Signature("Signature1", props.Font{
							Color: color.Color{Red: 20, Green: 20, Blue: 20},
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		signature := c.signature()
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		tableList := baseTableList()
		m := newMarotoTest(Fpdf, math, nil, nil, signature, nil, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, signature)
	}
}

func TestFpdfMaroto_Text(t *testing.T) {
	cases := []struct {
		name   string
		assert func(t *testing.T, signature *mocks.Text)
		act    func(m pdf.Maroto)
	}{
		{
			"One text inside one column, inside a row, without props",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 1)
				text.AssertCalled(t, "Add", "Text1", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Style:       consts.Normal,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        10.0,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text1")
					})
				})
			},
		},
		{
			"Two different text inside one colum, inside one row",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 2)
				text.AssertCalled(t, "Add", "Text2", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Style:       consts.Normal,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        10.0,
				})
				text.AssertCalled(t, "Add", "Text3", internal.Cell{
					X:      0.0,
					Y:      5.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Courier,
					Style:       consts.BoldItalic,
					Align:       consts.Center,
					Top:         5.0,
					Extrapolate: false,
					Size:        9.5,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text2")
						m.Text("Text3", props.Text{
							Family: consts.Courier,
							Style:  consts.BoldItalic,
							Size:   9.5,
							Align:  consts.Center,
							Top:    5.0,
						})
					})
				})
			},
		},
		{
			"Two different text with different columns, inside one row",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 2)
				text.AssertCalled(t, "Add", "Text4", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Style:       consts.Normal,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        10.0,
				})
				text.AssertCalled(t, "Add", "Text5", internal.Cell{
					X:      80.0,
					Y:      4.4,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Helvetica,
					Style:       consts.Italic,
					Align:       consts.Center,
					Top:         4.4,
					Extrapolate: false,
					Size:        8.5,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text4")
					})
					m.Col(12, func() {
						m.Text("Text5", props.Text{
							Family: consts.Helvetica,
							Style:  consts.Italic,
							Size:   8.5,
							Top:    4.4,
							Align:  consts.Center,
						})
					})
				})
			},
		},
		{
			"Two different text with different columns, inside one row",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 2)
				text.AssertCalled(t, "Add", "Text6", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Style:       consts.Normal,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        10.0,
				})
				text.AssertCalled(t, "Add", "Text7", internal.Cell{
					X:      0.0,
					Y:      40.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Courier,
					Style:       consts.BoldItalic,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        9.5,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text6")
					})
				})
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text7", props.Text{
							Family: consts.Courier,
							Style:  consts.BoldItalic,
							Size:   9.5,
						})
					})
				})
			},
		},
		{
			"When top is greater than row height",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 1)
				text.AssertCalled(t, "Add", "Text8", internal.Cell{
					X:      0.0,
					Y:      40.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Align:       consts.Left,
					Top:         40.0,
					Extrapolate: false,
					Size:        10.0,
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text8", props.Text{
							Top: 50,
						})
					})
				})
			},
		},
		{
			"custom color",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 1)
				text.AssertCalled(t, "Add", "Text1", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 0.0,
				}, props.Text{
					Family:      consts.Arial,
					Style:       consts.Normal,
					Align:       consts.Left,
					Top:         0.0,
					Extrapolate: false,
					Size:        10.0,
					Color: color.Color{
						Red:   20,
						Green: 20,
						Blue:  20,
					},
				})
			},
			func(m pdf.Maroto) {
				m.Row(40, func() {
					m.Col(12, func() {
						m.Text("Text1", props.Text{
							Color: color.Color{
								Red:   20,
								Green: 20,
								Blue:  20,
							},
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		text := baseTextTest()
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		tableList := baseTableList()
		m := newMarotoTest(Fpdf, math, nil, text, nil, nil, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, text)
	}
}

func TestFpdfMaroto_FileImage(t *testing.T) {
	cases := []struct {
		name   string
		image  func() *mocks.Image
		assert func(t *testing.T, image *mocks.Image)
		act    func(m pdf.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 1)
				image.AssertCalled(t, "AddFromFile", "Image1", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 100.0,
					Center:  false,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.FileImage("Image1")
					})
				})
			},
		},
		{
			"Two images inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image2", internal.Cell{X: 0.0, Y: 4.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    2,
					Top:     4,
					Percent: 40,
					Center:  false,
				})
				image.AssertCalled(t, "AddFromFile", "Image3", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 40,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						_ = m.FileImage("Image2", props.Rect{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						_ = m.FileImage("Image3", props.Rect{
							Percent: 40.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two images inside two cols inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image4", internal.Cell{X: 0.0, Y: 4.5, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    4,
					Top:     4.5,
					Percent: 55,
					Center:  false,
				})
				image.AssertCalled(t, "AddFromFile", "Image5", internal.Cell{X: 80.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 53,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						_ = m.FileImage("Image4", props.Rect{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(0, func() {
						_ = m.FileImage("Image5", props.Rect{
							Percent: 53.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two images inside one col inside two rows",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image6", internal.Cell{X: 0.0, Y: 8.5, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    7,
					Top:     8.5,
					Percent: 66,
					Center:  false,
				})
				image.AssertCalled(t, "AddFromFile", "Image7", internal.Cell{X: 0.0, Y: 20.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 98,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.FileImage("Image6", props.Rect{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.FileImage("Image7", props.Rect{
							Percent: 98.0,
							Center:  true,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		image := c.image()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, image, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, image)
	}
}

func TestFpdfMaroto_Base64Image(t *testing.T) {
	cases := []struct {
		name   string
		image  func() *mocks.Image
		assert func(t *testing.T, image *mocks.Image)
		act    func(m pdf.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 1)
				image.AssertCalled(t, "AddFromBase64", "Image1", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 100,
					Center:  false,
				}, consts.Jpg)
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.Base64Image("Image1", consts.Jpg)
					})
				})
			},
		},
		{
			"Two images inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image2", internal.Cell{X: 0.0, Y: 4.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    2,
					Top:     4,
					Percent: 40,
					Center:  false,
				}, consts.Png)
				image.AssertCalled(t, "AddFromBase64", "Image3", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 40,
					Center:  true,
				}, consts.Jpg)
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						_ = m.Base64Image("Image2", consts.Png, props.Rect{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						_ = m.Base64Image("Image3", consts.Jpg, props.Rect{
							Percent: 40.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two images inside two cols inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image4", internal.Cell{X: 0.0, Y: 4.5, Width: 40.0, Height: 20.0}, props.Rect{
					Left:    4,
					Top:     4.5,
					Percent: 55,
					Center:  false,
				}, consts.Png)
				image.AssertCalled(t, "AddFromBase64", "Image5", internal.Cell{X: 40.0, Y: 0.0, Width: 40.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 53,
					Center:  true,
				}, consts.Jpg)
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(6, func() {
						_ = m.Base64Image("Image4", consts.Png, props.Rect{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(6, func() {
						_ = m.Base64Image("Image5", consts.Jpg, props.Rect{
							Percent: 53.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two images inside one col inside two rows",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(nil)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image6", internal.Cell{X: 0.0, Y: 8.5, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    7,
					Top:     8.5,
					Percent: 66,
					Center:  false,
				}, consts.Png)
				image.AssertCalled(t, "AddFromBase64", "Image7", internal.Cell{X: 0.0, Y: 20.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    0,
					Top:     0,
					Percent: 98,
					Center:  true,
				}, consts.Jpg)
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.Base64Image("Image6", consts.Png, props.Rect{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.Base64Image("Image7", consts.Jpg, props.Rect{
							Percent: 98.0,
							Center:  true,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		image := c.image()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, image, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, image)
	}
}

// nolint:dupl // QrCode test
func TestFpdfMaroto_QrCode(t *testing.T) {
	cases := []struct {
		name   string
		code   func() *mocks.Code
		assert func(t *testing.T, image *mocks.Code)
		act    func(m pdf.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddQr", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddQr", 1)
				code.AssertCalled(t, "AddQr", "Code1", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0},
					props.Rect{Percent: 100, Center: false})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.QrCode("Code1")
					})
				})
			},
		},
		{
			"Two codes inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddQr", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddQr", 2)
				code.AssertCalled(t, "AddQr", "Code2", internal.Cell{X: 0.0, Y: 4.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    2.0,
					Top:     4.0,
					Percent: 40.0,
				})
				code.AssertCalled(t, "AddQr", "Code3", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Percent: 40.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.QrCode("Code2", props.Rect{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						m.QrCode("Code3", props.Rect{
							Percent: 40.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two codes inside two cols inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddQr", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddQr", 2)
				code.AssertCalled(t, "AddQr", "Code4", internal.Cell{X: 0.0, Y: 4.5, Width: 40.0, Height: 20.0}, props.Rect{
					Left:    4.0,
					Top:     4.5,
					Percent: 55.0,
				})
				code.AssertCalled(t, "AddQr", "Code5", internal.Cell{X: 40.0, Y: 0.0, Width: 40.0, Height: 20.0}, props.Rect{
					Percent: 53.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(6, func() {
						m.QrCode("Code4", props.Rect{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(6, func() {
						m.QrCode("Code5", props.Rect{
							Percent: 53.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two codes inside one col inside two rows",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddQr", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddQr", 2)
				code.AssertCalled(t, "AddQr", "Code6", internal.Cell{X: 0.0, Y: 8.5, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    7.0,
					Top:     8.5,
					Percent: 66.0,
				})
				code.AssertCalled(t, "AddQr", "Code7", internal.Cell{X: 0.0, Y: 20.0, Width: 80.0, Height: 20.0}, props.Rect{
					Percent: 98.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.QrCode("Code6", props.Rect{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(12, func() {
						m.QrCode("Code7", props.Rect{
							Percent: 98.0,
							Center:  true,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		code := c.code()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, code, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, code)
	}
}

func TestFpdfMaroto_Barcode(t *testing.T) {
	cases := []struct {
		name   string
		code   func() *mocks.Code
		assert func(t *testing.T, image *mocks.Code)
		act    func(m pdf.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddBar", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddBar", 1)
				code.AssertCalled(t, "AddBar", "Code1", internal.Cell{
					X:      0.0,
					Y:      0.0,
					Width:  80.0,
					Height: 20.0,
				}, props.Barcode{
					Percent: 100,
					Center:  false,
					Proportion: props.Proportion{
						Width:  1,
						Height: 0.2,
					},
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.Barcode("Code1", props.Barcode{Proportion: props.Proportion{Width: 1, Height: 0.2}})
					})
				})
			},
		},
		{
			"Two codes inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddBar", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddBar", 2)
				code.AssertCalled(t, "AddBar", "Code2", internal.Cell{X: 0.0, Y: 4.0, Width: 80.0, Height: 20.0}, props.Barcode{
					Left:       2.0,
					Top:        4.0,
					Percent:    40.0,
					Proportion: props.Proportion{Width: 1, Height: 0.2},
				})
				code.AssertCalled(t, "AddBar", "Code3", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Barcode{
					Percent:    40.0,
					Center:     true,
					Proportion: props.Proportion{Width: 1, Height: 0.2},
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(12, func() {
						_ = m.Barcode("Code2", props.Barcode{
							Left:       2.0,
							Top:        4.0,
							Percent:    40.0,
							Proportion: props.Proportion{Width: 1, Height: 0.2},
						})
						_ = m.Barcode("Code3", props.Barcode{
							Percent:    40.0,
							Center:     true,
							Proportion: props.Proportion{Width: 1, Height: 0.2},
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		code := c.code()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, code, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, code)
	}
}

// nolint:dupl // DataMatrixCod test
func TestFpdfMaroto_DataMatrixCode(t *testing.T) {
	cases := []struct {
		name   string
		code   func() *mocks.Code
		assert func(t *testing.T, image *mocks.Code)
		act    func(m pdf.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddDataMatrix", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddDataMatrix", 1)
				code.AssertCalled(t, "AddDataMatrix", "Code1",
					internal.Cell{
						X:      0.0,
						Y:      0.0,
						Width:  80.0,
						Height: 20.0,
					}, props.Rect{
						Percent: 100,
						Center:  false,
					})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.DataMatrixCode("Code1")
					})
				})
			},
		},
		{
			"Two codes inside a col inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddDataMatrix", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddDataMatrix", 2)
				code.AssertCalled(t, "AddDataMatrix", "Code2", internal.Cell{X: 0.0, Y: 4.0, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    2.0,
					Top:     4.0,
					Percent: 40.0,
				})
				code.AssertCalled(t, "AddDataMatrix", "Code3", internal.Cell{X: 0.0, Y: 0.0, Width: 80.0, Height: 20.0}, props.Rect{
					Percent: 40.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.DataMatrixCode("Code2", props.Rect{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						m.DataMatrixCode("Code3", props.Rect{
							Percent: 40.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two codes inside two cols inside a row",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddDataMatrix", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddDataMatrix", 2)
				code.AssertCalled(t, "AddDataMatrix", "Code4", internal.Cell{X: 0.0, Y: 4.5, Width: 40.0, Height: 20.0}, props.Rect{
					Left:    4.0,
					Top:     4.5,
					Percent: 55.0,
				})
				code.AssertCalled(t, "AddDataMatrix", "Code5", internal.Cell{X: 40.0, Y: 0.0, Width: 40.0, Height: 20.0}, props.Rect{
					Percent: 53.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(6, func() {
						m.DataMatrixCode("Code4", props.Rect{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(6, func() {
						m.DataMatrixCode("Code5", props.Rect{
							Percent: 53.0,
							Center:  true,
						})
					})
				})
			},
		},
		{
			"Two codes inside one col inside two rows",
			func() *mocks.Code {
				code := &mocks.Code{}
				code.On("AddDataMatrix", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return code
			},
			func(t *testing.T, code *mocks.Code) {
				code.AssertNumberOfCalls(t, "AddDataMatrix", 2)
				code.AssertCalled(t, "AddDataMatrix", "Code6", internal.Cell{X: 0.0, Y: 8.5, Width: 80.0, Height: 20.0}, props.Rect{
					Left:    7.0,
					Top:     8.5,
					Percent: 66.0,
				})
				code.AssertCalled(t, "AddDataMatrix", "Code7", internal.Cell{X: 0.0, Y: 20.0, Width: 80.0, Height: 20.0}, props.Rect{
					Percent: 98.0,
					Center:  true,
				})
			},
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.Col(0, func() {
						m.DataMatrixCode("Code6", props.Rect{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(12, func() {
						m.DataMatrixCode("Code7", props.Rect{
							Percent: 98.0,
							Center:  true,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		code := c.code()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, code, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, code)
	}
}

func TestFpdfMaroto_Row(t *testing.T) {
	cases := []struct {
		name                 string
		act                  func(m pdf.Maroto, calledTimes *int)
		assertRowCalledTimes func(t *testing.T, calledTimes int)
		assertFpdfCalls      func(t *testing.T, Fpdf *mocks.Fpdf)
	}{
		{
			"One row",
			func(m pdf.Maroto, calledTimes *int) {
				m.Row(30, func() {
					*calledTimes++
				})
			},
			func(t *testing.T, calledTimes int) {
				assert.Equal(t, calledTimes, 1)
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 3)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 3)
				Fpdf.AssertNumberOfCalls(t, "Ln", 2)
				Fpdf.AssertCalled(t, "Ln", 30.0)
			},
		},
		{
			"Two rows",
			func(m pdf.Maroto, calledTimes *int) {
				m.Row(30, func() {
					*calledTimes++
				})
				m.Row(40, func() {
					*calledTimes++
				})
			},
			func(t *testing.T, calledTimes int) {
				assert.Equal(t, calledTimes, 2)
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 4)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 4)
				Fpdf.AssertNumberOfCalls(t, "Ln", 3)

				Fpdf.AssertCalled(t, "Ln", 30.0)
				Fpdf.AssertCalled(t, "Ln", 40.0)
			},
		},
		{
			"Three rows",
			func(m pdf.Maroto, calledTimes *int) {
				m.Row(30, func() {
					*calledTimes++
				})
				m.Row(40, func() {
					*calledTimes++
				})
				m.Row(10, func() {
					*calledTimes++
				})
			},
			func(t *testing.T, calledTimes int) {
				assert.Equal(t, calledTimes, 3)
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 5)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 5)
				Fpdf.AssertNumberOfCalls(t, "Ln", 4)

				Fpdf.AssertCalled(t, "Ln", 30.0)
				Fpdf.AssertCalled(t, "Ln", 40.0)
				Fpdf.AssertCalled(t, "Ln", 10.0)
			},
		},
		{
			"Rows to add new page",
			func(m pdf.Maroto, calledTimes *int) {
				m.Row(50, func() {
					*calledTimes++
				})
				m.Row(40, func() {
					*calledTimes++
				})
				m.Row(45, func() {
					*calledTimes++
				})
			},
			func(t *testing.T, calledTimes int) {
				assert.Equal(t, calledTimes, 3)
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 10)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 10)
				Fpdf.AssertNumberOfCalls(t, "Ln", 6)

				Fpdf.AssertCalled(t, "Ln", 50.0)
				Fpdf.AssertCalled(t, "Ln", 40.0)
				Fpdf.AssertCalled(t, "Ln", 45.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, nil, tableList, nil)
		calledTimes := 0

		// Act
		c.act(m, &calledTimes)

		// Assert
		c.assertRowCalledTimes(t, calledTimes)
		c.assertFpdfCalls(t, Fpdf)
	}
}

func TestFpdfMaroto_Line(t *testing.T) {
	cases := []struct {
		name   string
		Fpdf   func() *mocks.Fpdf
		Line   func() *mocks.Line
		assert func(t *testing.T, Fpdf *mocks.Fpdf, line *mocks.Line)
		act    func(m pdf.Maroto)
	}{
		{
			"Line without prop",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				return Fpdf
			},
			func() *mocks.Line {
				line := new(mocks.Line)
				line.On("Draw", mock.Anything, mock.Anything)
				return line
			},
			func(t *testing.T, Fpdf *mocks.Fpdf, line *mocks.Line) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 5)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 5)

				line.AssertNumberOfCalls(t, "Draw", 1)
				line.AssertCalled(t, "Draw", internal.Cell{
					X:      10,
					Y:      10.5,
					Width:  90,
					Height: 10.5,
				}, props.Line{
					Color: color.Color{
						Red:   0,
						Green: 0,
						Blue:  0,
					},
					Style: consts.Solid,
					Width: 0.1,
				})
			},
			func(m pdf.Maroto) {
				m.Line(1.0)
			},
		},
		{
			"One solid line without color",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Line {
				line := new(mocks.Line)
				line.On("Draw", mock.Anything, mock.Anything)
				return line
			},
			func(t *testing.T, Fpdf *mocks.Fpdf, line *mocks.Line) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 5)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 5)

				line.AssertNumberOfCalls(t, "Draw", 1)
				line.AssertCalled(t, "Draw", internal.Cell{
					X:      10,
					Y:      11,
					Width:  90,
					Height: 11,
				}, props.Line{
					Color: color.Color{
						Red:   255,
						Green: 100,
						Blue:  50,
					},
					Style: consts.Solid,
					Width: 0.1,
				})
			},
			func(m pdf.Maroto) {
				m.Line(2.0, props.Line{
					Color: color.Color{
						Red:   255,
						Green: 100,
						Blue:  50,
					},
				})
			},
		},
		// nolint:dupl // better this way
		{
			"One dashed line without color",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Line {
				line := new(mocks.Line)
				line.On("Draw", mock.Anything, mock.Anything)
				return line
			},
			func(t *testing.T, Fpdf *mocks.Fpdf, line *mocks.Line) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 5)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 5)

				line.AssertNumberOfCalls(t, "Draw", 1)
				line.AssertCalled(t, "Draw", internal.Cell{
					X:      10,
					Y:      11,
					Width:  90,
					Height: 11,
				}, props.Line{
					Color: color.Color{
						Red:   0,
						Green: 0,
						Blue:  0,
					},
					Style: consts.Dashed,
					Width: 0.1,
				})
			},
			func(m pdf.Maroto) {
				m.Line(2.0, props.Line{
					Style: consts.Dashed,
				})
			},
		},
		// nolint:dupl // better this way
		{
			"One dotted line without color",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Line {
				line := new(mocks.Line)
				line.On("Draw", mock.Anything, mock.Anything)
				return line
			},
			func(t *testing.T, Fpdf *mocks.Fpdf, line *mocks.Line) {
				Fpdf.AssertNumberOfCalls(t, "GetMargins", 5)
				Fpdf.AssertNumberOfCalls(t, "GetPageSize", 5)

				line.AssertNumberOfCalls(t, "Draw", 1)
				line.AssertCalled(t, "Draw", internal.Cell{
					X:      10,
					Y:      11,
					Width:  90,
					Height: 11,
				}, props.Line{
					Color: color.Color{
						Red:   0,
						Green: 0,
						Blue:  0,
					},
					Style: consts.Dotted,
					Width: 0.1,
				})
			},
			func(m pdf.Maroto) {
				m.Line(2.0, props.Line{
					Style: consts.Dotted,
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		line := c.Line()
		math := baseMathTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, nil, tableList, line)

		// Act
		c.act(m)

		// Assert
		c.assert(t, Fpdf, line)
	}
}

func TestFpdfMaroto_ColSpace(t *testing.T) {
	cases := []struct {
		name   string
		act    func(m pdf.Maroto)
		assert func(t *testing.T, Fpdf *mocks.Fpdf)
	}{
		{
			"One ColSpace inside one Row",
			func(m pdf.Maroto) {
				m.Row(40.0, func() {
					m.ColSpace(12)
				})
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "CellFormat", 2)
				Fpdf.AssertCalled(t, "CellFormat", 80, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpace inside one Row",
			func(m pdf.Maroto) {
				m.Row(40.0, func() {
					m.ColSpace(5)
					m.ColSpace(7)
				})
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "CellFormat", 3)
				Fpdf.AssertCalled(t, "CellFormat", 33, 40, "", "", 0, "C", false, 0, "")
				Fpdf.AssertCalled(t, "CellFormat", 46, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpace inside two Rows",
			func(m pdf.Maroto) {
				m.Row(33.0, func() {
					m.ColSpace(12)
				})
				m.Row(35.0, func() {
					m.ColSpace(12)
				})
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertCalled(t, "CellFormat", 80, 33, "", "", 0, "C", false, 0, "")
				Fpdf.AssertCalled(t, "CellFormat", 80, 35, "", "", 0, "C", false, 0, "")
				Fpdf.AssertNumberOfCalls(t, "CellFormat", 3)
			},
		},
		{
			"ColSpace with Border",
			func(m pdf.Maroto) {
				m.SetBorder(true)
				m.Row(23.0, func() {
					m.ColSpace(12)
				})
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "CellFormat", 2)
				Fpdf.AssertCalled(t, "CellFormat", 80, 23, "", "1", 0, "C", false, 0, "")
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, Fpdf)
	}
}

func TestFpdfMaroto_Output(t *testing.T) {
	cases := []struct {
		name              string
		Fpdf              func() *mocks.Fpdf
		hasFooter         bool
		assertFpdfCalls   func(t *testing.T, Fpdf *mocks.Fpdf)
		assertBytes       func(t *testing.T, bytes bytes.Buffer)
		assertError       func(t *testing.T, err error)
		assertFooterCalls func(t *testing.T, footerCalls int)
	}{
		{
			"When Output returns an error",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Output", mock.Anything).Return(errors.New("AnyError"))
				return Fpdf
			},
			false,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Output", 1)
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Equal(t, err.Error(), "AnyError")
			},
			func(t *testing.T, footerCalls int) {
				assert.Zero(t, footerCalls)
			},
		},
		{
			"When Output not returns an error",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Output", mock.Anything).Return(nil)
				return Fpdf
			},
			false,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Output", 1)
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			func(t *testing.T, footerCalls int) {
				assert.Zero(t, footerCalls)
			},
		},
		{
			"When Output has footer",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("Output", mock.Anything).Return(nil)
				return Fpdf
			},
			true,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Output", 1)
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			func(t *testing.T, footerCalls int) {
				assert.NotZero(t, footerCalls)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		math := baseMathTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, nil, tableList, nil)
		footerCalls := 0

		// Act
		if c.hasFooter {
			m.RegisterFooter(func() {
				footerCalls++
			})
		}

		bytes, err := m.Output()

		// Assert
		c.assertFpdfCalls(t, Fpdf)
		c.assertFooterCalls(t, footerCalls)
		c.assertBytes(t, bytes)
		c.assertError(t, err)
	}
}

func TestFpdfMaroto_OutputFileAndClose(t *testing.T) {
	cases := []struct {
		name            string
		Fpdf            func() *mocks.Fpdf
		assertFpdfCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertBytes     func(t *testing.T, bytes bytes.Buffer)
		assertError     func(t *testing.T, err error)
	}{
		{
			"When OutputFileAndClose returns an error",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("OutputFileAndClose", mock.Anything).Return(errors.New("AnyError"))
				return Fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "OutputFileAndClose", 1)
				Fpdf.AssertCalled(t, "OutputFileAndClose", "AnyName")
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Equal(t, err.Error(), "AnyError")
			},
		},
		{
			"When OutputFileAndClose not returns an error",
			func() *mocks.Fpdf {
				Fpdf := baseFpdfTest(10, 10, 10)
				Fpdf.On("OutputFileAndClose", mock.Anything).Return(nil)
				return Fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "OutputFileAndClose", 1)
				Fpdf.AssertCalled(t, "OutputFileAndClose", "AnyName")
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		math := baseMathTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, nil, nil, nil, nil, nil, tableList, nil)

		// Act
		err := m.OutputFileAndClose("AnyName")

		// Assert
		c.assertFpdfCalls(t, Fpdf)
		c.assertError(t, err)
	}
}

func newMarotoTest(fFpdf *mocks.Fpdf, math *mocks.Math, font *mocks.Font, text *mocks.Text,
	signature *mocks.Signature, image *mocks.Image, code *mocks.Code, tableList *mocks.TableList,
	line *mocks.Line,
) pdf.Maroto {
	m := &pdf.PdfMaroto{
		Pdf:             fFpdf,
		Math:            math,
		Font:            font,
		TextHelper:      text,
		SignHelper:      signature,
		Image:           image,
		Code:            code,
		TableListHelper: tableList,
		LineHelper:      line,
	}

	m.SetDefaultFontFamily(consts.Arial)
	m.SetBackgroundColor(color.NewWhite())

	return m
}

func baseFpdfTest(left, top, right float64) *mocks.Fpdf {
	Fpdf := &mocks.Fpdf{}
	Fpdf.On("GetPageSize").Return(100.0, 100.0)
	Fpdf.On("GetMargins").Return(left, top, right, 0.0)
	Fpdf.On("CellFormat", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	Fpdf.On("Ln", mock.Anything)
	Fpdf.On("AliasNbPages", mock.Anything)
	Fpdf.On("GetFontSize").Return(1.0, 1.0)
	Fpdf.On("SetMargins", mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"))
	Fpdf.On("SetFillColor", mock.Anything, mock.Anything, mock.Anything)
	Fpdf.On("AddUTF8Font", mock.Anything, mock.Anything, mock.Anything)
	Fpdf.On("SetFontLocation", mock.Anything)
	Fpdf.On("SetProtection", mock.Anything, mock.Anything, mock.Anything)
	Fpdf.On("SetCompression", mock.Anything)
	return Fpdf
}

func baseMathTest() *mocks.Math {
	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(20.0)
	return math
}

func baseTextTest() *mocks.Text {
	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	return text
}

func baseFontTest() *mocks.Font {
	font := &mocks.Font{}
	font.On("GetFontSize").Return(1.0, 1.0)
	return font
}

func baseTableList() *mocks.TableList {
	tableList := &mocks.TableList{}
	tableList.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	return tableList
}

func getContents() ([]string, [][]string) {
	header := []string{"j = 0", "j = 1", "j = 2", "j = 4"}

	contents := [][]string{}
	for i := 0; i < 20; i++ {
		content := []string{}
		for j := 0; j < 4; j++ {
			content = append(content, fmt.Sprintf("i = %d, j = %d", i, j))
		}
		contents = append(contents, content)
	}

	return header, contents
}

func TestFpdfMaroto_RegisterFooter(t *testing.T) {
	cases := []struct {
		name      string
		act       func(m pdf.Maroto)
		hasFooter bool
		assert    func(t *testing.T, footerCalls int)
	}{
		{
			"Always execute footer once",
			func(m pdf.Maroto) {
			},
			true,
			func(t *testing.T, footerCalls int) {
				assert.Equal(t, footerCalls, 1)
			},
		},
		{
			"Execute 6 times when create a 6 pages",
			func(m pdf.Maroto) {
				headers, contents := getContents()
				m.Row(20, func() {
					for _, header := range headers {
						m.Col(0, func() {
							m.Text(header)
						})
					}
				})

				for _, content := range contents {
					m.Row(20, func() {
						for _, txt := range content {
							m.Col(12, func() {
								m.Text(txt)
							})
						}
					})
				}
			},
			true,
			func(t *testing.T, footerCalls int) {
				assert.Equal(t, footerCalls, 6)
			},
		},
		{
			"When footer is nil, not execute",
			func(m pdf.Maroto) {
				header, contents := getContents()
				m.TableList(header, contents)
			},
			false,
			func(t *testing.T, footerCalls int) {
				assert.Equal(t, footerCalls, 0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		text := baseTextTest()
		font := baseFontTest()
		tableList := baseTableList()
		footerCalls := 0

		m := newMarotoTest(Fpdf, math, font, text, nil, nil, nil, tableList, nil)

		if c.hasFooter {
			m.RegisterFooter(func() {
				footerCalls++
			})
		}

		// Act
		c.act(m)

		// Assert
		c.assert(t, footerCalls)
	}
}

func TestFpdfMaroto_RegisterHeader(t *testing.T) {
	cases := []struct {
		name       string
		act        func(m pdf.Maroto)
		hasClosure bool
		assert     func(t *testing.T, headerCalls int)
	}{
		{
			"Always execute header once when add something",
			func(m pdf.Maroto) {
				m.Row(20, func() {
					m.ColSpace(0)
				})
			},
			true,
			func(t *testing.T, headerCalls int) {
				assert.Equal(t, headerCalls, 1)
			},
		},
		{
			"Execute 6 times when create a 6 pages",
			func(m pdf.Maroto) {
				headers, contents := getContents()
				m.Row(20, func() {
					for _, header := range headers {
						m.Col(12, func() {
							m.Text(header)
						})
					}
				})

				for _, content := range contents {
					m.Row(20, func() {
						for _, txt := range content {
							m.Col(12, func() {
								m.Text(txt)
							})
						}
					})
				}
			},
			true,
			func(t *testing.T, headerCalls int) {
				assert.Equal(t, headerCalls, 6)
			},
		},
		{
			"When header is nil not execute",
			func(m pdf.Maroto) {
				header, contents := getContents()
				m.TableList(header, contents)
			},
			false,
			func(t *testing.T, headerCalls int) {
				assert.Equal(t, headerCalls, 0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		text := baseTextTest()
		font := baseFontTest()
		tableList := baseTableList()

		headerCalls := 0

		m := newMarotoTest(Fpdf, math, font, text, nil, nil, nil, tableList, nil)

		if c.hasClosure {
			m.RegisterHeader(func() {
				headerCalls++
			})
		}

		// Act
		c.act(m)

		// Assert
		c.assert(t, headerCalls)
	}
}

func TestFpdfMaroto_GetCurrentPage(t *testing.T) {
	cases := []struct {
		name   string
		act    func(m pdf.Maroto)
		assert func(t *testing.T, pageIndex int)
	}{
		{
			"When create page index should be 1",
			func(m pdf.Maroto) {
			},
			func(t *testing.T, pageIndex int) {
				assert.Equal(t, pageIndex, 0)
			},
		},
		{
			"When has a secund page, page index should be 2",
			func(m pdf.Maroto) {
				headers, contents := getContents()
				m.Row(20, func() {
					for _, header := range headers {
						m.Col(uint(12/len(headers)), func() {
							m.Text(header)
						})
					}
				})

				for _, content := range contents {
					m.Row(20, func() {
						for _, txt := range content {
							m.Col(uint(12/len(contents)), func() {
								m.Text(txt)
							})
						}
					})
				}
			},
			func(t *testing.T, pageIndex int) {
				assert.Equal(t, pageIndex, 5)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(10, 10, 10)
		math := baseMathTest()
		text := baseTextTest()
		font := baseFontTest()
		tableList := baseTableList()

		m := newMarotoTest(Fpdf, math, font, text, nil, nil, nil, tableList, nil)

		// Act
		c.act(m)

		// Assert
		pageIndex := m.GetCurrentPage()
		c.assert(t, pageIndex)
	}
}

func TestFpdfMaroto_GetCurrentPage_WhenCreateOffsetIsZero(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10, 10, 10)
	math := baseMathTest()
	text := baseTextTest()
	font := baseFontTest()
	tableList := baseTableList()

	m := newMarotoTest(Fpdf, math, font, text, nil, nil, nil, tableList, nil)

	// Act
	offset := m.GetCurrentOffset()

	// Assert
	assert.Zero(t, offset)
}

func TestFpdfMaroto_GetCurrentPage_WhenIsNotZero(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10, 10, 10)
	math := baseMathTest()
	text := baseTextTest()
	font := baseFontTest()
	tableList := baseTableList()

	m := newMarotoTest(Fpdf, math, font, text, nil, nil, nil, tableList, nil)

	m.Row(20, func() {
		m.Col(0, func() {
			m.Text("test")
		})
	})

	// Act
	offset := m.GetCurrentOffset()

	// Assert
	assert.Equal(t, offset, float64(20))
}

func TestFpdfMaroto_SetPageMargins(t *testing.T) {
	cases := []struct {
		name   string
		act    func(m pdf.Maroto)
		assert func(t *testing.T, m *mocks.Fpdf)
	}{
		{
			"Set page margins should override default, top greater than 10",
			func(m pdf.Maroto) {
				m.SetPageMargins(12.3, 19.3, 0.0)
			},
			func(t *testing.T, m *mocks.Fpdf) {
				m.AssertCalled(t, "SetMargins", 12.3, 10.0, 0.0)
			},
		},
		{
			"Set page margins should override default, top less than 10",
			func(m pdf.Maroto) {
				m.SetPageMargins(12.3, 9.0, 0.0)
			},
			func(t *testing.T, m *mocks.Fpdf) {
				m.AssertCalled(t, "SetMargins", 12.3, 10.0, 0.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := baseFpdfTest(12.3, 19.3, 0)

		m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, Fpdf)
	}
}

func TestFpdfMaroto_SetBackgroundColor(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(12.3, 19.3, 0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)
	white := color.NewWhite()

	// Act
	m.SetBackgroundColor(white)

	// Assert
	Fpdf.AssertCalled(t, "SetFillColor", white.Red, white.Green, white.Blue)
	Fpdf.AssertNumberOfCalls(t, "SetFillColor", 2)
}

func TestFpdfMaroto_GetPageMargins(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(12.3, 19.3, 0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	left, top, right, bottom := m.GetPageMargins()

	// Assert
	assert.Equal(t, 12.3, left)
	assert.Equal(t, 19.3, top)
	assert.Equal(t, 0.0, right)
	assert.Equal(t, 0.0, bottom)
}

func TestFpdfMaroto_AddPage(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.AddPage()

	// Assert
	Fpdf.AssertCalled(t, "CellFormat", 80, 0, "", "", 0, "C", false, 0, "")
	Fpdf.AssertCalled(t, "SetFillColor", 255, 255, 255)
	Fpdf.AssertCalled(t, "Ln", 90.0)
}

func TestPdfMaroto_AddUTF8Font(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.AddUTF8Font("family", "style", "file")

	// Assert
	Fpdf.AssertCalled(t, "AddUTF8Font", "family", "style", "file")
}

func TestPdfMaroto_SetFontLocation(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.SetFontLocation("/opt/fonts")

	// Assert
	Fpdf.AssertCalled(t, "SetFontLocation", "/opt/fonts")
}

func TestPdfMaroto_SetProtection(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.SetProtection(0, "userPassStr", "ownerPassStr")

	// Assert
	Fpdf.AssertCalled(t, "SetProtection", byte(0), "userPassStr", "ownerPassStr")
}

func TestPdfMaroto_SetGetDefaultFontFamily(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.SetDefaultFontFamily("family")
	family := m.GetDefaultFontFamily()

	// Assert
	assert.Equal(t, family, "family")
}

func TestPdfMaroto_SetCompression(t *testing.T) {
	// Arrange
	Fpdf := baseFpdfTest(10.0, 10.0, 10.0)
	m := newMarotoTest(Fpdf, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	m.SetCompression(false)

	// Assert
	Fpdf.AssertCalled(t, "SetCompression", false)
}
