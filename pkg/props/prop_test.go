package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
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
				assert.Equal(t, prop.Proportion.Height, 0.20)
			},
		},
		{
			"When height is smaller than 10% of width",
			props.Barcode{
				Proportion: props.Proportion{
					Width:  11,
					Height: 1,
				},
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Proportion.Height, prop.Proportion.Width*0.10)
			},
		},
		{
			"When height is grather than 20% of width",
			props.Barcode{
				Proportion: props.Proportion{
					Width:  5,
					Height: 5,
				},
			},
			func(t *testing.T, prop props.Barcode) {
				assert.Equal(t, prop.Proportion.Height, prop.Proportion.Width*0.20)
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
		{
			"When vertical padding is less than 0",
			&props.Text{
				VerticalPadding: -5.0,
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.VerticalPadding, 0.0)
			},
		},
	}

	for _, c := range cases {
		c.fontProp.MakeValid(consts.Arial)
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
		c.signatureProp.MakeValid(consts.Arial)

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
				HeaderProp:  props.TableListContent{},
				ContentProp: props.TableListContent{},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Size, 10.0)
				assert.Equal(t, m.HeaderProp.Family, consts.Arial)
				assert.Equal(t, m.HeaderProp.Style, consts.Bold)
				assert.Equal(t, len(m.HeaderProp.GridSizes), 3)
				assert.Equal(t, m.ContentProp.Size, 10.0)
				assert.Equal(t, m.ContentProp.Family, consts.Arial)
				assert.Equal(t, m.ContentProp.Style, consts.Normal)
				assert.Equal(t, len(m.ContentProp.GridSizes), 3)
			},
		},
		{
			"When HeaderProp.Size is 0.0",
			&props.TableList{
				HeaderProp: props.TableListContent{
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
				HeaderProp: props.TableListContent{
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
				HeaderProp: props.TableListContent{
					Family: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.HeaderProp.Family, consts.Arial)
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
				ContentProp: props.TableListContent{
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
				HeaderProp: props.TableListContent{
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
				HeaderProp: props.TableListContent{
					Family: "",
				},
			},
			func(t *testing.T, m *props.TableList) {
				assert.Equal(t, m.ContentProp.Family, consts.Arial)
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
		c.tableListProp.MakeValid([]string{"a", "b", "c"}, consts.Arial)

		// Assert
		c.assert(t, c.tableListProp)
	}
}

func TestLine_MakeValid(t *testing.T) {
	cases := []struct {
		name        string
		prop        *props.Line
		spaceHeight float64
		assert      func(t *testing.T, m *props.Line)
	}{
		{
			"When style not defined must use solid",
			&props.Line{},
			1.0,
			func(t *testing.T, m *props.Line) {
				assert.Equal(t, m.Style, consts.Solid)
			},
		},
		{
			"When width not defined must use 0.1",
			&props.Line{},
			1.0,
			func(t *testing.T, m *props.Line) {
				assert.Equal(t, m.Width, 0.1)
			},
		},
		{
			"When width greater than space height",
			&props.Line{
				Width: 5.0,
			},
			1.0,
			func(t *testing.T, m *props.Line) {
				assert.Equal(t, m.Width, 1.0)
			},
		},
	}

	for _, c := range cases {
		// Act
		c.prop.MakeValid(c.spaceHeight)

		// Assert
		c.assert(t, c.prop)
	}
}
