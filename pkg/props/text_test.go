package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

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
				assert.Equal(t, prop.Family, fontfamily.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&props.Text{
				Style: "",
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Style, fontstyle.Normal)
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
				assert.Equal(t, prop.Align, align.Left)
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
			"When Left is less than 0, it should become 0",
			&props.Text{
				Left: -5.0,
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Left, 0.0)
			},
		},
		{
			"When Right is less than 0, it should become 0",
			&props.Text{
				Right: -5.0,
			},
			func(t *testing.T, prop *props.Text) {
				assert.Equal(t, prop.Right, 0.0)
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
		c.fontProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})
		c.assert(t, c.fontProp)
	}
}
