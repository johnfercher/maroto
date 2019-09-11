package props_test

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectProp_MakeValid(t *testing.T) {
	cases := []struct {
		name     string
		rectProp props.Rect
		assert   func(t *testing.T, prop props.Rect)
	}{
		{
			"When percent is less than zero, should become 100",
			props.Rect{
				Percent: -2,
			},
			func(t *testing.T, prop props.Rect) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When percent is greater than 100, should become 100",
			props.Rect{
				Percent: 102,
			},
			func(t *testing.T, prop props.Rect) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When is center, top and left should become 0",
			props.Rect{
				Center: true,
				Top:    5,
				Left:   5,
			},
			func(t *testing.T, prop props.Rect) {
				assert.Equal(t, prop.Top, 0.0)
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When left is less than 0, should become 0",
			props.Rect{
				Left: -5,
			},
			func(t *testing.T, prop props.Rect) {
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When top is less than 0, should become 0",
			props.Rect{
				Top: -5,
			},
			func(t *testing.T, prop props.Rect) {
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
		barcodeProp props.Barcode
		assert      func(t *testing.T, prop props.Barcode)
	}{
		{
			"When percent is less than zero, should become 100",
			props.Barcode{
				Percent: -2,
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When percent is greater than 100, should become 100",
			props.Barcode{
				Percent: 102,
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Percent, 100.0)
			},
		},
		{
			"When is center, top and left should become 0",
			props.Barcode{
				Center: true,
				Top:    5,
				Left:   5,
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Top, 0.0)
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When left is less than 0, should become 0",
			props.Barcode{
				Left: -5,
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When top is less than 0, should become 0",
			props.Barcode{
				Top: -5,
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Top, 0.0)
			},
		},
		{
			"When proportion.width less than 0",
			props.Barcode{
				Proportion: props.Proportion{
					Width: -5,
				},
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Proportion.Width, 1.0)
			},
		},
		{
			"When proportion.height less than 0",
			props.Barcode{
				Proportion: props.Proportion{
					Height: -5,
				},
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Proportion.Height, 0.33)
			},
		},
	}

	for _, c := range cases {
		c.barcodeProp.MakeValid()
		c.assert(t, c.barcodeProp)
	}
}

func TestText_MakeValid(t *testing.T) {
	cases := []struct {
		name     string
		fontProp *props.Text
		assert   func(t *testing.T, prop *props.Text)
	}{
		{
			"When family is not defined, should define arial",
			&props.Text{
				Family: "",
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Family, consts.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&props.Text{
				Style: "",
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Style, consts.Normal)
			},
		},
		{
			"When size is zero, should define 10.0",
			&props.Text{
				Size: 0.0,
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Size, 10.0)
			},
		},
		{
			"When align is not defined, should define Left",
			&props.Text{
				Align: "",
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Align, consts.Left)
			},
		},
		{
			"When top is less than 0, should become 0",
			&props.Text{
				Top: -5.0,
			},
			func(t *testing.T, prop *props.Text) {
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
		signatureProp *props.Font
		assert        func(t *testing.T, prop *props.Font)
	}{
		{
			"When family is not defined, should define arial",
			&props.Font{
				Family: "",
			},
			func(t *testing.T, prop *props.Font) {
				assert.Equal(t, prop.Family, consts.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&props.Font{
				Style: "",
			},
			func(t *testing.T, prop *props.Font) {
				assert.Equal(t, prop.Style, consts.Bold)
			},
		},
		{
			"When size is zero, should define 10.0",
			&props.Font{
				Size: 0.0,
			},
			func(t *testing.T, prop *props.Font) {
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
		tableListProp *props.TableList
		assert        func(t *testing.T, m *props.TableList)
	}{
		{
			"When HeaderProp/ContentProp is not defined",
			&props.TableList{
				HeaderProp:  props.Font{},
				ContentProp: props.Font{},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Size, 10.0)
				assert.Equal(t, m.HeaderProp.Family, consts.Arial)
				assert.Equal(t, m.HeaderProp.Style, consts.Bold)
				assert.Equal(t, m.ContentProp.Size, 10.0)
				assert.Equal(t, m.ContentProp.Family, consts.Arial)
				assert.Equal(t, m.ContentProp.Style, consts.Normal)
			},
		},
		{
			"When HeaderProp.Size is 0.0",
			&props.TableList{
				HeaderProp: props.Font{
					Size: 0.0,
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Size, 10.0)
			},
		},
		{
			"When HeaderProp.Style is empty",
			&props.TableList{
				HeaderProp: props.Font{
					Style: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Style, consts.Bold)
			},
		},
		{
			"When HeaderProp.Family is empty",
			&props.TableList{
				HeaderProp: props.Font{
					Family: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Family, consts.Arial)
			},
		},
		{
			"When HeaderHeight is 0.0",
			&props.TableList{
				HeaderHeight: 0.0,
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderHeight, 7.0)
			},
		},
		{
			"When Align is empty",
			&props.TableList{
				Align: "",
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.Align, consts.Left)
			},
		},
		{
			"When ContentProp.Size is 0.0",
			&props.TableList{
				ContentProp: props.Font{
					Size: 0.0,
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.ContentProp.Size, 10.0)
			},
		},
		{
			"When ContentProp.Style is empty",
			&props.TableList{
				HeaderProp: props.Font{
					Style: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.ContentProp.Style, consts.Normal)
			},
		},
		{
			"When ContentProp.Family is empty",
			&props.TableList{
				HeaderProp: props.Font{
					Family: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.ContentProp.Family, consts.Arial)
			},
		},
		{
			"When ContentHeight is 0.0",
			&props.TableList{
				ContentHeight: 0.0,
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.ContentHeight, 5.0)
			},
		},
		{
			"When HeaderContentSpace is 0.0",
			&props.TableList{
				HeaderContentSpace: 0.0,
			},
			func(t *testing.T, m *props.TableList) {
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
