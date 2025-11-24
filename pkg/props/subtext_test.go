package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestSubText_MakeValid(t *testing.T) {
	cases := []struct {
		name     string
		fontProp *props.SubText
		assert   func(t *testing.T, prop *props.SubText)
	}{
		{
			"When family is not defined, should define arial",
			&props.SubText{
				Family: "",
			},
			func(t *testing.T, prop *props.SubText) {
				assert.Equal(t, prop.Family, fontfamily.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&props.SubText{
				Style: "",
			},
			func(t *testing.T, prop *props.SubText) {
				assert.Equal(t, prop.Style, fontstyle.Normal)
			},
		},
		{
			"When size is zero, should define 10.0",
			&props.SubText{
				Size: 0.0,
			},
			func(t *testing.T, prop *props.SubText) {
				assert.Equal(t, prop.Size, 10.0)
			},
		},
	}

	for _, c := range cases {
		c.fontProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})
		c.assert(t, c.fontProp)
	}
}
