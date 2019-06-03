package maroto_test

import (
	"github.com/johnfercher/maroto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectProp_MakeValid(t *testing.T) {
	cases := []struct {
		name       string
		rectProp   maroto.RectProp
		assertRect func(t *testing.T, prop maroto.RectProp)
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
		c.assertRect(t, c.rectProp)
	}
}

func TestTextProp_MakeValid(t *testing.T) {
	cases := []struct {
		name           string
		fontProp       *maroto.TextProp
		assertFontProp func(t *testing.T, prop *maroto.TextProp)
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
		c.assertFontProp(t, c.fontProp)
	}
}

func TestSignatureProp_MakeValid(t *testing.T) {
	cases := []struct {
		name                string
		signatureProp       *maroto.SignatureProp
		assertSignatureProp func(t *testing.T, prop *maroto.SignatureProp)
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
		c.signatureProp.MakeValid()
		c.assertSignatureProp(t, c.signatureProp)
	}
}
