package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"
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
