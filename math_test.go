package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMath(t *testing.T) {
	math := maroto.NewMath(&mocks.Pdf{})

	assert.NotNil(t, math)
	assert.Equal(t, fmt.Sprintf("%T", math), "*maroto.math")
}

func TestMath_GetWidthPerCol(t *testing.T) {
	cases := []struct {
		name        string
		qtdCols     float64
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertWidth func(t *testing.T, width float64)
	}{
		{
			"1 col, margins 10 10",
			1,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 190)
			},
		},
		{
			"2 col, margins 10 10",
			2,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 95)
			},
		},
		{
			"4 col, margins 10 10",
			4,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 47)
			},
		},
		{
			"1 col, margins 20 20",
			1,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 170)
			},
		},
		{
			"2 col, margins 20 20",
			2,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 85)
			},
		},
		{
			"4 col, margins 20 20",
			4,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 42)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := maroto.NewMath(pdf)

		// Act
		width := math.GetWidthPerCol(c.qtdCols)

		// Assert
		c.assertWidth(t, width)
		c.assertCalls(t, pdf)
	}
}
