package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewMaroto(t *testing.T) {
	cases := []struct {
		name        string
		orientation maroto.Orientation
		pageSize    maroto.PageSize
		assert      func(t *testing.T, m maroto.Maroto)
	}{
		{
			"When portrait and A4",
			maroto.Portrait,
			maroto.A4,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 210.0, 0.1)
				assert.InDelta(t, height, 297.0, 0.1)
			},
		},
		{
			"When portrait and A3",
			maroto.Portrait,
			maroto.A3,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 297.0, 0.1)
				assert.InDelta(t, height, 419.9, 0.1)
			},
		},
		{
			"When portrait and A5",
			maroto.Portrait,
			maroto.A5,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 148.4, 0.1)
				assert.InDelta(t, height, 210.0, 0.1)
			},
		},
		{
			"When portrait and Legal",
			maroto.Portrait,
			maroto.Legal,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 215.9, 0.1)
				assert.InDelta(t, height, 355.6, 0.1)
			},
		},
		{
			"When portrait and Letter",
			maroto.Portrait,
			maroto.Letter,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, width, 215.9, 0.1)
				assert.InDelta(t, height, 279.4, 0.1)
			},
		},
		{
			"When landscape and A4",
			maroto.Landscape,
			maroto.A4,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 210.0, 0.1)
				assert.InDelta(t, width, 297.0, 0.1)
			},
		},
		{
			"When landscape and A3",
			maroto.Landscape,
			maroto.A3,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 297.0, 0.1)
				assert.InDelta(t, width, 419.9, 0.1)
			},
		},
		{
			"When landscape and A5",
			maroto.Landscape,
			maroto.A5,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 148.4, 0.1)
				assert.InDelta(t, width, 210.0, 0.1)
			},
		},
		{
			"When landscape and Legal",
			maroto.Landscape,
			maroto.Legal,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 215.9, 0.1)
				assert.InDelta(t, width, 355.6, 0.1)
			},
		},
		{
			"When landscape and Letter",
			maroto.Landscape,
			maroto.Letter,
			func(t *testing.T, m maroto.Maroto) {
				assert.NotNil(t, m)
				assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.PdfMaroto")
				width, height := m.GetPageSize()
				assert.InDelta(t, height, 215.9, 0.1)
				assert.InDelta(t, width, 279.4, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Act
		m := maroto.NewMaroto(c.orientation, c.pageSize)

		// Assert
		c.assert(t, m)
	}

}

func TestPdfMaroto_SetGetDebugMode(t *testing.T) {
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	assert.False(t, m.GetDebugMode())
	m.SetDebugMode(true)

	assert.True(t, m.GetDebugMode())
}

func TestPdfMaroto_Signature(t *testing.T) {
	cases := []struct {
		name            string
		signature       func() *mocks.Signature
		assertSignature func(t *testing.T, signature *mocks.Signature)
		actSignature    func(m maroto.Maroto)
	}{
		{
			"One signature inside one column, inside a row, without props",
			func() *mocks.Signature {
				signature := &mocks.Signature{}
				signature.On("AddSpaceFor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return signature
			},
			func(t *testing.T, signature *mocks.Signature) {
				signature.AssertNumberOfCalls(t, "AddSpaceFor", 1)
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row("Row", 40, func() {
					m.Col("Col", func() {
						m.Signature("SignHelper", nil)
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
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper2", maroto.Courier, maroto.BoldItalic, 9.5, 1.0, 40.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row("Row", 40, func() {
					m.Col("Col", func() {
						m.Signature("SignHelper", nil)
						m.Signature("SignHelper2", &maroto.SignatureProp{
							Family: maroto.Courier,
							Style:  maroto.BoldItalic,
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
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper", maroto.Arial, maroto.Bold, 8.0, 2.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper2", maroto.Courier, maroto.BoldItalic, 9.5, 2.0, 40.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row("Row", 40, func() {
					m.Col("Col", func() {
						m.Signature("SignHelper", nil)
					})
					m.Col("Col", func() {
						m.Signature("SignHelper2", &maroto.SignatureProp{
							Family: maroto.Courier,
							Style:  maroto.BoldItalic,
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
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "SignHelper2", maroto.Courier, maroto.BoldItalic, 9.5, 1.0, 80.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row("Row", 40, func() {
					m.Col("Col", func() {
						m.Signature("SignHelper", nil)
					})
				})
				m.Row("Row", 40, func() {
					m.Col("Col", func() {
						m.Signature("SignHelper2", &maroto.SignatureProp{
							Family: maroto.Courier,
							Style:  maroto.BoldItalic,
							Size:   9.5,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		signature := c.signature()
		pdf := basePdfTest()
		math := baseMathTest()
		m := newMarotoTest(pdf, math, nil, nil, signature, nil, nil)

		// Act
		c.actSignature(m)

		// Assert
		c.assertSignature(t, signature)
	}
}

func newMarotoTest(fpdf *mocks.Pdf, math *mocks.Math, font *mocks.Font, text *mocks.Text, signature *mocks.Signature, image *mocks.Image, code *mocks.Code) maroto.Maroto {
	m := &maroto.PdfMaroto{
		Pdf:        fpdf,
		Math:       math,
		Font:       font,
		TextHelper: text,
		SignHelper: signature,
		Image:      image,
		Code:       code,
	}

	return m
}

func basePdfTest() *mocks.Pdf {
	pdf := &mocks.Pdf{}
	pdf.On("GetPageSize").Return(100.0, 100.0)
	pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
	pdf.On("CellFormat", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	pdf.On("Ln", mock.Anything)
	return pdf
}

func baseMathTest() *mocks.Math {
	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(20.0)
	return math
}
