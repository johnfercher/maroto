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

func TestSignatureProp_MakeValid(t *testing.T) {
	cases := []struct {
		name          string
		signatureProp *maroto.SignatureProp
		assert        func(t *testing.T, prop *maroto.SignatureProp)
	}{
		{
			"When family is not defined, should define arial",
			&maroto.SignatureProp{
				Family: "",
			},
			func(t *testing.T, prop *maroto.SignatureProp) {
				assert.Equal(t, prop.Family, maroto.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&maroto.SignatureProp{
				Style: "",
			},
			func(t *testing.T, prop *maroto.SignatureProp) {
				assert.Equal(t, prop.Style, maroto.Bold)
			},
		},
		{
			"When size is zero, should define 10.0",
			&maroto.SignatureProp{
				Size: 0.0,
			},
			func(t *testing.T, prop *maroto.SignatureProp) {
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
			"When HFontSize is 0.0",
			&maroto.TableListProp{
				HFontSize: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HFontSize, 10.0)
			},
		},
		{
			"When HFontStyle is empty",
			&maroto.TableListProp{
				HFontStyle: "",
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HFontStyle, maroto.Bold)
			},
		},
		{
			"When HFontFamily is empty",
			&maroto.TableListProp{
				HFontFamily: "",
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HFontFamily, maroto.Arial)
			},
		},
		{
			"When HHeight is 0.0",
			&maroto.TableListProp{
				HHeight: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.HHeight, 7.0)
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
			"When CFontSize is 0.0",
			&maroto.TableListProp{
				CFontSize: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.CFontSize, 10.0)
			},
		},
		{
			"When CFontStyle is empty",
			&maroto.TableListProp{
				CFontStyle: "",
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.CFontStyle, maroto.Normal)
			},
		},
		{
			"When CFontFamily is empty",
			&maroto.TableListProp{
				CFontFamily: "",
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.CFontFamily, maroto.Arial)
			},
		},
		{
			"When CHeight is 0.0",
			&maroto.TableListProp{
				CHeight: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.CHeight, 5.0)
			},
		},
		{
			"When Space is 0.0",
			&maroto.TableListProp{
				Space: 0.0,
			},
			func(t *testing.T, m *maroto.TableListProp) {
				assert.Equal(t, m.Space, 4.0)
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
