package maroto_test

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/pkg/errors"
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
	// Arrange
	m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

	// Assert & Act
	assert.False(t, m.GetDebugMode())
	m.SetDebugMode(true)

	// Assert
	assert.True(t, m.GetDebugMode())
}

func TestPdfMaroto_Signature(t *testing.T) {
	cases := []struct {
		name      string
		signature func() *mocks.Signature
		assert    func(t *testing.T, signature *mocks.Signature)
		act       func(m maroto.Maroto)
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
				signature.AssertCalled(t, "AddSpaceFor", "Signature1", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Signature("Signature1", nil)
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
				signature.AssertCalled(t, "AddSpaceFor", "Signature2", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "Signature3", maroto.Courier, maroto.BoldItalic, 9.5, 1.0, 40.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Signature("Signature2", nil)
						m.Signature("Signature3", &maroto.SignatureProp{
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
				signature.AssertCalled(t, "AddSpaceFor", "Signature4", maroto.Arial, maroto.Bold, 8.0, 2.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "Signature5", maroto.Courier, maroto.BoldItalic, 9.5, 2.0, 40.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Signature("Signature4", nil)
					})
					m.Col(func() {
						m.Signature("Signature5", &maroto.SignatureProp{
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
				signature.AssertCalled(t, "AddSpaceFor", "Signature6", maroto.Arial, maroto.Bold, 8.0, 1.0, 40.0, 0.0)
				signature.AssertCalled(t, "AddSpaceFor", "Signature7", maroto.Courier, maroto.BoldItalic, 9.5, 1.0, 80.0, 0.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Signature("Signature6", nil)
					})
				})
				m.Row(40, func() {
					m.Col(func() {
						m.Signature("Signature7", &maroto.SignatureProp{
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
		c.act(m)

		// Assert
		c.assert(t, signature)
	}
}

func TestPdfMaroto_Text(t *testing.T) {
	cases := []struct {
		name   string
		assert func(t *testing.T, signature *mocks.Text)
		act    func(m maroto.Maroto)
	}{
		{
			"One text inside one column, inside a row, without props",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 1)
				text.AssertCalled(t, "Add", "Text1", maroto.Arial, maroto.Normal, 10.0, 0.0, maroto.Left, 0.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text1", nil)
					})
				})
			},
		},
		{
			"Two different text inside one colum, inside one row",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 2)
				text.AssertCalled(t, "Add", "Text2", maroto.Arial, maroto.Normal, 10.0, 0.0, maroto.Left, 0.0, 1.0)
				text.AssertCalled(t, "Add", "Text3", maroto.Courier, maroto.BoldItalic, 9.5, 5.0, maroto.Center, 0.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text2", nil)
						m.Text("Text3", &maroto.TextProp{
							Family: maroto.Courier,
							Style:  maroto.BoldItalic,
							Size:   9.5,
							Align:  maroto.Center,
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
				text.AssertCalled(t, "Add", "Text4", maroto.Arial, maroto.Normal, 10.0, 0.0, maroto.Left, 0.0, 2.0)
				text.AssertCalled(t, "Add", "Text5", maroto.Helvetica, maroto.Italic, 8.5, 4.4, maroto.Center, 1.0, 2.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text4", nil)
					})
					m.Col(func() {
						m.Text("Text5", &maroto.TextProp{
							Family: maroto.Helvetica,
							Style:  maroto.Italic,
							Size:   8.5,
							Top:    4.4,
							Align:  maroto.Center,
						})
					})
				})
			},
		},
		{
			"Two different text with different columns, inside one row",
			func(t *testing.T, text *mocks.Text) {
				text.AssertNumberOfCalls(t, "Add", 2)
				text.AssertCalled(t, "Add", "Text6", maroto.Arial, maroto.Normal, 10.0, 0.0, maroto.Left, 0.0, 1.0)
				text.AssertCalled(t, "Add", "Text7", maroto.Courier, maroto.BoldItalic, 9.5, 40.0, maroto.Left, 0.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text6", nil)
					})
				})
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text7", &maroto.TextProp{
							Family: maroto.Courier,
							Style:  maroto.BoldItalic,
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
				text.AssertCalled(t, "Add", "Text8", maroto.Arial, maroto.Normal, 10.0, 40.0, maroto.Left, 0.0, 1.0)
			},
			func(m maroto.Maroto) {
				m.Row(40, func() {
					m.Col(func() {
						m.Text("Text8", &maroto.TextProp{
							Top: 50,
						})
					})
				})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		text := baseTextTest()
		pdf := basePdfTest()
		math := baseMathTest()
		m := newMarotoTest(pdf, math, nil, text, nil, nil, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, text)
	}
}

func TestPdfMaroto_FileImage(t *testing.T) {
	cases := []struct {
		name   string
		image  func() *mocks.Image
		assert func(t *testing.T, image *mocks.Image)
		act    func(m maroto.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 1)
				image.AssertCalled(t, "AddFromFile", "Image1", 0.0, 0.0, 1.0, 20.0, 100.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.FileImage("Image1", nil)
					})
				})
			},
		},
		{
			"Two images inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image2", 4.0, 0.0, 1.0, 20.0, 40.0)
				image.AssertCalled(t, "AddFromFile", "Image3", 0.0, 0.0, 1.0, 20.0, 40.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.FileImage("Image2", &maroto.RectProp{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						m.FileImage("Image3", &maroto.RectProp{
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
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image4", 4.5, 0.0, 2.0, 20.0, 55.0)
				image.AssertCalled(t, "AddFromFile", "Image5", 0.0, 1.0, 2.0, 20.0, 53.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.FileImage("Image4", &maroto.RectProp{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(func() {
						m.FileImage("Image5", &maroto.RectProp{
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
				image.On("AddFromFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromFile", 2)
				image.AssertCalled(t, "AddFromFile", "Image6", 8.5, 0.0, 1.0, 20.0, 66.0)
				image.AssertCalled(t, "AddFromFile", "Image7", 20.0, 0.0, 1.0, 20.0, 98.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.FileImage("Image6", &maroto.RectProp{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(func() {
						m.FileImage("Image7", &maroto.RectProp{
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
		pdf := basePdfTest()
		math := baseMathTest()
		image := c.image()

		m := newMarotoTest(pdf, math, nil, nil, nil, image, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, image)
	}
}

func TestPdfMaroto_Base64Image(t *testing.T) {
	cases := []struct {
		name   string
		image  func() *mocks.Image
		assert func(t *testing.T, image *mocks.Image)
		act    func(m maroto.Maroto)
	}{
		{
			"One code inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 1)
				image.AssertCalled(t, "AddFromBase64", "Image1", 0.0, 0.0, 1.0, 20.0, 100.0, maroto.Jpg)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.Base64Image("Image1", maroto.Jpg, nil)
					})
				})
			},
		},
		{
			"Two images inside a col inside a row",
			func() *mocks.Image {
				image := &mocks.Image{}
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image2", 4.0, 0.0, 1.0, 20.0, 40.0, maroto.Png)
				image.AssertCalled(t, "AddFromBase64", "Image3", 0.0, 0.0, 1.0, 20.0, 40.0, maroto.Jpg)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.Base64Image("Image2", maroto.Png, &maroto.RectProp{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						m.Base64Image("Image3", maroto.Jpg, &maroto.RectProp{
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
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image4", 4.5, 0.0, 2.0, 20.0, 55.0, maroto.Png)
				image.AssertCalled(t, "AddFromBase64", "Image5", 0.0, 1.0, 2.0, 20.0, 53.0, maroto.Jpg)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.Base64Image("Image4", maroto.Png, &maroto.RectProp{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(func() {
						m.Base64Image("Image5", maroto.Jpg, &maroto.RectProp{
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
				image.On("AddFromBase64", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return image
			},
			func(t *testing.T, image *mocks.Image) {
				image.AssertNumberOfCalls(t, "AddFromBase64", 2)
				image.AssertCalled(t, "AddFromBase64", "Image6", 8.5, 0.0, 1.0, 20.0, 66.0, maroto.Png)
				image.AssertCalled(t, "AddFromBase64", "Image7", 20.0, 0.0, 1.0, 20.0, 98.0, maroto.Jpg)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.Base64Image("Image6", maroto.Png, &maroto.RectProp{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(func() {
						m.Base64Image("Image7", maroto.Jpg, &maroto.RectProp{
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
		pdf := basePdfTest()
		math := baseMathTest()
		image := c.image()

		m := newMarotoTest(pdf, math, nil, nil, nil, image, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, image)
	}
}

func TestPdfMaroto_QrCode(t *testing.T) {
	cases := []struct {
		name   string
		code   func() *mocks.Code
		assert func(t *testing.T, image *mocks.Code)
		act    func(m maroto.Maroto)
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
				code.AssertCalled(t, "AddQr", "Code1", 0.0, 0.0, 1.0, 20.0, 100.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.QrCode("Code1", nil)
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
				code.AssertCalled(t, "AddQr", "Code2", 4.0, 0.0, 1.0, 20.0, 40.0)
				code.AssertCalled(t, "AddQr", "Code3", 0.0, 0.0, 1.0, 20.0, 40.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.QrCode("Code2", &maroto.RectProp{
							Left:    2.0,
							Top:     4.0,
							Percent: 40.0,
						})
						m.QrCode("Code3", &maroto.RectProp{
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
				code.AssertCalled(t, "AddQr", "Code4", 4.5, 0.0, 2.0, 20.0, 55.0)
				code.AssertCalled(t, "AddQr", "Code5", 0.0, 1.0, 2.0, 20.0, 53.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.QrCode("Code4", &maroto.RectProp{
							Left:    4.0,
							Top:     4.5,
							Percent: 55.0,
						})
					})
					m.Col(func() {
						m.QrCode("Code5", &maroto.RectProp{
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
				code.AssertCalled(t, "AddQr", "Code6", 8.5, 0.0, 1.0, 20.0, 66.0)
				code.AssertCalled(t, "AddQr", "Code7", 20.0, 0.0, 1.0, 20.0, 98.0)
			},
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.Col(func() {
						m.QrCode("Code6", &maroto.RectProp{
							Left:    7.0,
							Top:     8.5,
							Percent: 66.0,
						})
					})
				})
				m.Row(20, func() {
					m.Col(func() {
						m.QrCode("Code7", &maroto.RectProp{
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
		pdf := basePdfTest()
		math := baseMathTest()
		code := c.code()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, code)

		// Act
		c.act(m)

		// Assert
		c.assert(t, code)
	}
}

func TestPdfMaroto_Row(t *testing.T) {
	cases := []struct {
		name                 string
		act                  func(m maroto.Maroto, calledTimes *int)
		assertRowCalledTimes func(t *testing.T, calledTimes int)
		assertPdfCalls       func(t *testing.T, pdf *mocks.Pdf)
	}{
		{
			"One row",
			func(m maroto.Maroto, calledTimes *int) {
				m.Row(30, func() {
					*calledTimes++
				})
			},
			func(t *testing.T, calledTimes int) {
				assert.Equal(t, calledTimes, 1)
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "Ln", 1)
				pdf.AssertCalled(t, "Ln", 30.0)
			},
		},
		{
			"Two rows",
			func(m maroto.Maroto, calledTimes *int) {
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
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 2)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 2)
				pdf.AssertNumberOfCalls(t, "Ln", 2)

				pdf.AssertCalled(t, "Ln", 30.0)
				pdf.AssertCalled(t, "Ln", 40.0)
			},
		},
		{
			"Three rows",
			func(m maroto.Maroto, calledTimes *int) {
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
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 3)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 3)
				pdf.AssertNumberOfCalls(t, "Ln", 3)

				pdf.AssertCalled(t, "Ln", 30.0)
				pdf.AssertCalled(t, "Ln", 40.0)
				pdf.AssertCalled(t, "Ln", 10.0)
			},
		},
		{
			"Rows to add new page",
			func(m maroto.Maroto, calledTimes *int) {
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
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 3)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 3)
				pdf.AssertNumberOfCalls(t, "Ln", 3)

				pdf.AssertCalled(t, "Ln", 50.0)
				pdf.AssertCalled(t, "Ln", 40.0)
				pdf.AssertCalled(t, "Ln", 45.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := basePdfTest()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)
		calledTimes := 0

		// Act
		c.act(m, &calledTimes)

		// Assert
		c.assertRowCalledTimes(t, calledTimes)
		c.assertPdfCalls(t, pdf)
	}
}

func TestPdfMaroto_Line(t *testing.T) {
	cases := []struct {
		name   string
		pdf    func() *mocks.Pdf
		assert func(t *testing.T, pdf *mocks.Pdf)
		act    func(m maroto.Maroto)
	}{
		{
			"One line",
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 2)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 2)

				pdf.AssertNumberOfCalls(t, "Line", 1)
				pdf.AssertCalled(t, "Line", 10.0, 10.5, 90.0, 10.5)
			},
			func(m maroto.Maroto) {
				m.Line(1.0)
			},
		},
		{
			"Two lines",
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 4)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 4)

				pdf.AssertNumberOfCalls(t, "Line", 2)
				pdf.AssertCalled(t, "Line", 10.0, 11.0, 90.0, 11.0)
				pdf.AssertCalled(t, "Line", 10.0, 14.0, 90.0, 14.0)
			},
			func(m maroto.Maroto) {
				m.Line(2.0)
				m.Line(4.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, pdf)
	}
}

func TestPdfMaroto_ColSpace(t *testing.T) {
	cases := []struct {
		name   string
		act    func(m maroto.Maroto)
		assert func(t *testing.T, pdf *mocks.Pdf)
	}{
		{
			"One ColSpace inside one Row",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpace()
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 1)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpace inside one Row",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpace()
					m.ColSpace()
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 2)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpace inside two Rows",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpace()
				})
				m.Row(35.0, func() {
					m.ColSpace()
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 2)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
				pdf.AssertCalled(t, "CellFormat", 20, 35, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"ColSpace with Debug",
			func(m maroto.Maroto) {
				m.SetDebugMode(true)
				m.Row(40.0, func() {
					m.ColSpace()
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 1)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "1", 0, "C", false, 0, "")
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := basePdfTest()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, pdf)
	}
}

func TestPdfMaroto_ColSpaces(t *testing.T) {
	cases := []struct {
		name   string
		act    func(m maroto.Maroto)
		assert func(t *testing.T, pdf *mocks.Pdf)
	}{
		{
			"One ColSpaces inside one Row",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpaces(2)
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 2)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpaces inside one Row",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpaces(2)
					m.ColSpaces(2)
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 4)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"Two ColSpaces inside two Rows",
			func(m maroto.Maroto) {
				m.Row(40.0, func() {
					m.ColSpaces(2)
				})
				m.Row(35.0, func() {
					m.ColSpaces(2)
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 4)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "", 0, "C", false, 0, "")
				pdf.AssertCalled(t, "CellFormat", 20, 35, "", "", 0, "C", false, 0, "")
			},
		},
		{
			"ColSpaces with Debug",
			func(m maroto.Maroto) {
				m.SetDebugMode(true)
				m.Row(40.0, func() {
					m.ColSpaces(2)
				})
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "CellFormat", 2)
				pdf.AssertCalled(t, "CellFormat", 20, 40, "", "1", 0, "C", false, 0, "")
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := basePdfTest()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)

		// Act
		c.act(m)

		// Assert
		c.assert(t, pdf)
	}
}

func TestPdfMaroto_Output(t *testing.T) {
	cases := []struct {
		name           string
		pdf            func() *mocks.Pdf
		assertPdfCalls func(t *testing.T, pdf *mocks.Pdf)
		assertBytes    func(t *testing.T, bytes bytes.Buffer)
		assertError    func(t *testing.T, err error)
	}{
		{
			"When Output returns an error",
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("Output", mock.Anything).Return(errors.New("AnyError"))
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Output", 1)
			},
			func(t *testing.T, bytes bytes.Buffer) {
				assert.Nil(t, bytes.Bytes())
			},
			func(t *testing.T, err error) {
				assert.Equal(t, err.Error(), "AnyError")
			},
		},
		{
			"When Output not returns an error",
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("Output", mock.Anything).Return(nil)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Output", 1)
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
		pdf := c.pdf()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)

		// Act
		bytes, err := m.Output()

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertBytes(t, bytes)
		c.assertError(t, err)
	}
}

func TestPdfMaroto_OutputFileAndClose(t *testing.T) {
	cases := []struct {
		name           string
		pdf            func() *mocks.Pdf
		assertPdfCalls func(t *testing.T, pdf *mocks.Pdf)
		assertBytes    func(t *testing.T, bytes bytes.Buffer)
		assertError    func(t *testing.T, err error)
	}{
		{
			"When OutputFileAndClose returns an error",
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("OutputFileAndClose", mock.Anything).Return(errors.New("AnyError"))
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "OutputFileAndClose", 1)
				pdf.AssertCalled(t, "OutputFileAndClose", "AnyName")
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
			func() *mocks.Pdf {
				pdf := basePdfTest()
				pdf.On("OutputFileAndClose", mock.Anything).Return(nil)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "OutputFileAndClose", 1)
				pdf.AssertCalled(t, "OutputFileAndClose", "AnyName")
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
		pdf := c.pdf()
		math := baseMathTest()

		m := newMarotoTest(pdf, math, nil, nil, nil, nil, nil)

		// Act
		err := m.OutputFileAndClose("AnyName")

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertError(t, err)
	}
}

func TestPdfMaroto_RegisterHeader(t *testing.T) {
	cases := []struct {
		name        string
		headerCalls int
		act         func(m maroto.Maroto)
		assert      func(t *testing.T, headerCalls int)
	}{
		{
			"Always execute header once",
			0,
			func(m maroto.Maroto) {
				m.Row(20, func() {
					m.ColSpace()
				})
			},
			func(t *testing.T, headerCalls int) {
				assert.Equal(t, headerCalls, 1)
			},
		},
		{
			"Execute twice when create a second page",
			0,
			func(m maroto.Maroto) {
				header, contents := getContents()
				m.TableList(header, contents, nil)
			},
			func(t *testing.T, headerCalls int) {
				assert.Equal(t, headerCalls, 2)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := basePdfTest()
		math := baseMathTest()
		text := baseTextTest()
		headerCalls := c.headerCalls

		m := newMarotoTest(pdf, math, nil, text, nil, nil, nil)

		m.RegisterHeader(func() {
			headerCalls++
		})

		// Act
		c.act(m)

		// Assert
		c.assert(t, headerCalls)
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

func baseTextTest() *mocks.Text {
	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	return text
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
