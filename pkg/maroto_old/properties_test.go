package maroto_test

import (
	"github.com/johnfercher/maroto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectProp_MakeValid(t *testing.T) {
	cases := []struct {
		name     string
		rectProp maroto.RectProp
		assert   func(t *testing.T, prop maroto.RectProp)
	}{
		{
			"When percent is less than zero, should become 100",
			maroto.RectProp{
				Percent: -2,
			},
			func(t *testing.T, prop maroto.RectProp) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When percent is greater than 100, should become 100",
			maroto.RectProp{
				Percent: 102,
			},
			func(t *testing.T, prop maroto.RectProp) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When is center, top and left should become 0",
			maroto.RectProp{
				Center: true,
				Top:    5,
				Left:   5,
			},
			func(t *testing.T, prop maroto.RectProp) {
				assert.Equal(t, prop.Top, 0.0)
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When left is less than 0, should become 0",
			maroto.RectProp{
				Left: -5,
			},
			func(t *testing.T, prop maroto.RectProp) {
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When top is less than 0, should become 0",
			maroto.RectProp{
				Top: -5,
			},
			func(t *testing.T, prop maroto.RectProp) {
				assert.Equal(t, prop.Top, 0.0)
			},
		},
	}

	for _, c := range cases {
		c.rectProp.MakeValid()
		c.assert(t, c.rectProp)
	}
}

func TestBarcodeProp_MakeValid(t *testing.T) {
	cases := []struct {
		name        string
		barcodeProp maroto.BarcodeProp
		assert      func(t *testing.T, prop maroto.BarcodeProp)
	}{
		{
			"When percent is less than zero, should become 100",
			maroto.BarcodeProp{
				Percent: -2,
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When percent is greater than 100, should become 100",
			maroto.BarcodeProp{
				Percent: 102,
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When is center, top and left should become 0",
			maroto.BarcodeProp{
				Center: true,
				Top:    5,
				Left:   5,
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Top, 0.0)
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When left is less than 0, should become 0",
			maroto.BarcodeProp{
				Left: -5,
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When top is less than 0, should become 0",
			maroto.BarcodeProp{
				Top: -5,
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Top, 0.0)
			},
		},
		{
			"When proportion.width less than 0",
			maroto.BarcodeProp{
				Proportion: maroto.Proportion{
					Width: -5,
				},
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Proportion.Width, 1.0)
			},
		},
		{
			"When proportion.height less than 0",
			maroto.BarcodeProp{
				Proportion: maroto.Proportion{
					Height: -5,
				},
			},
			func(t *testing.T, prop maroto.BarcodeProp) {
				assert.Equal(t, prop.Proportion.Height, 0.33)
			},
		},
	}

	for _, c := range cases {
		c.barcodeProp.MakeValid()
		c.assert(t, c.barcodeProp)
	}
}

func TestTextProp_MakeValid(t *testing.T) {
	cases := []struct {
		name     string
		fontProp *maroto.TextProp
		assert   func(t *testing.T, prop *maroto.TextProp)
	}{
		{
			"When family is not defined, should define arial",
			&maroto.TextProp{
				Family: "",
			},
			func(t *testing.T, prop *maroto.TextProp) {
				assert.Equal(t, prop.Family, maroto.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&maroto.TextProp{
				Style: "",
			},
			func(t *testing.T, prop *maroto.TextProp) {
				assert.Equal(t, prop.Style, maroto.Normal)
			},
		},
		{
			"When size is zero, should define 10.0",
			&maroto.TextProp{
				Size: 0.0,
			},
			func(t *testing.T, prop *maroto.TextProp) {
				assert.Equal(t, prop.Size, 10.0)
			},
		},
		{
			"When align is not defined, should define Left",
			&maroto.TextProp{
				Align: "",
			},
			func(t *testing.T, prop *maroto.TextProp) {
				assert.Equal(t, prop.Align, maroto.Left)
			},
		},
		{
			"When top is less than 0, should become 0",
			&maroto.TextProp{
				Top: -5.0,
			},
			func(t *testing.T, prop *maroto.TextProp) {
				assert.Equal(t, prop.Top, 0.0)
			},
		},
	}

	for _, c := range cases {
		c.fontProp.MakeValid()
		c.assert(t, c.fontProp)
	}
}

func TestFontProp_MakeValid(t *testing.T) {
	cases := []struct {
		name          string
		signatureProp *maroto.FontProp
		assert        func(t *testing.T, prop *maroto.FontProp)
	}{
		{
			"When family is not defined, should define arial",
			&maroto.FontProp{
				Family: "",
			},
			func(t *testing.T, prop *maroto.FontProp) {
				assert.Equal(t, prop.Family, maroto.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&maroto.FontProp{
				Style: "",
			},
			func(t *testing.T, prop *maroto.FontProp) {
				assert.Equal(t, prop.Style, maroto.Bold)
			},
		},
		{
			"When size is zero, should define 10.0",
			&maroto.FontProp{
				Size: 0.0,
			},
			func(t *testing.T, prop *maroto.FontProp) {
				assert.Equal(t, prop.Size, 8.0)
			},
		},
	}

	for _, c := range cases {
		// Act
		c.signatureProp.MakeValid()

		// Assert
		c.assert(t, c.signatureProp)
	}
}

func TestTableListProp_MakeValid(t *testing.T) {
	cases := []struct {
		name          string
		tableListProp *maroto.TableListProp
		assert        func(t *testing.T, m *maroto.TableListProp)
	}{
		{
			"When HeaderProp/ContentProp is not defined",
			&maroto.TableListProp{
				HeaderProp:  maroto.FontProp{},
				ContentProp: maroto.FontProp{},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderProp.Size, 10.0)
				assert.Equal(t, m.HeaderProp.Family, maroto.Arial)
				assert.Equal(t, m.HeaderProp.Style, maroto.Bold)
				assert.Equal(t, m.ContentProp.Size, 10.0)
				assert.Equal(t, m.ContentProp.Family, maroto.Arial)
				assert.Equal(t, m.ContentProp.Style, maroto.Normal)
			},
		},
		{
			"When HeaderProp.Size is 0.0",
			&maroto.TableListProp{
				HeaderProp: maroto.FontProp{
					Size: 0.0,
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderProp.Size, 10.0)
			},
		},
		{
			"When HeaderProp.Style is empty",
			&maroto.TableListProp{
				HeaderProp: maroto.FontProp{
					Style: "",
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderProp.Style, maroto.Bold)
			},
		},
		{
			"When HeaderProp.Family is empty",
			&maroto.TableListProp{
				HeaderProp: maroto.FontProp{
					Family: "",
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderProp.Family, maroto.Arial)
			},
		},
		{
			"When HeaderHeight is 0.0",
			&maroto.TableListProp{
				HeaderHeight: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderHeight, 7.0)
			},
		},
		{
			"When Align is empty",
			&maroto.TableListProp{
				Align: "",
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.Align, maroto.Left)
			},
		},
		{
			"When ContentProp.Size is 0.0",
			&maroto.TableListProp{
				ContentProp: maroto.FontProp{
					Size: 0.0,
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.ContentProp.Size, 10.0)
			},
		},
		{
			"When ContentProp.Style is empty",
			&maroto.TableListProp{
				HeaderProp: maroto.FontProp{
					Style: "",
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.ContentProp.Style, maroto.Normal)
			},
		},
		{
			"When ContentProp.Family is empty",
			&maroto.TableListProp{
				HeaderProp: maroto.FontProp{
					Family: "",
				},
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.ContentProp.Family, maroto.Arial)
			},
		},
		{
			"When ContentHeight is 0.0",
			&maroto.TableListProp{
				ContentHeight: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.ContentHeight, 5.0)
			},
		},
		{
			"When HeaderContentSpace is 0.0",
			&maroto.TableListProp{
				HeaderContentSpace: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HeaderContentSpace, 4.0)
			},
		},
	}

	for _, c := range cases {
		// Act
		c.tableListProp.MakeValid()

		// Assert
		c.assert(t, c.tableListProp)
	}
}
