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
