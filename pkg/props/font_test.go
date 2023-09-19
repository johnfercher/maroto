package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

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
				assert.Equal(t, prop.Family, fontfamily.Arial)
			},
		},
		{
			"When style is not defined, should define normal",
			&props.Font{
				Style: "",
			},
			func(t *testing.T, prop *props.Font) {
				assert.Equal(t, prop.Style, fontstyle.Bold)
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
		c.signatureProp.MakeValid(fontfamily.Arial)

		// Assert
		c.assert(t, c.signatureProp)
	}
}
