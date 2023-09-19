package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

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
