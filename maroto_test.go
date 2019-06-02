package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/stretchr/testify/assert"
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
