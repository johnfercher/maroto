package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/stretchr/testify/assert"
)

func TestWhiteColor(t *testing.T) {
	// Act
	white := props.WhiteColor

	// Assert
	assert.Equal(t, 255, white.Red)
	assert.Equal(t, 255, white.Green)
	assert.Equal(t, 255, white.Blue)
}

func TestBlackColor(t *testing.T) {
	// Act
	black := props.BlackColor

	// Assert
	assert.Equal(t, 0, black.Red)
	assert.Equal(t, 0, black.Green)
	assert.Equal(t, 0, black.Blue)
}

/*func TestColor_IsWhite(t *testing.T) {
	// Act
	white := props.NewWhite()

	// Assert
	assert.True(t, white.IsWhite())
}*/
