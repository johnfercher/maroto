package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

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
				assert.Equal(t, m.Style, linestyle.Solid)
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
