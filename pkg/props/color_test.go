package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"

	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/stretchr/testify/assert"
)

func TestWhiteColor(t *testing.T) {
	// Act
	sut := props.WhiteColor

	// Assert
	assert.Equal(t, 255, sut.Red)
	assert.Equal(t, 255, sut.Green)
	assert.Equal(t, 255, sut.Blue)
}

func TestBlackColor(t *testing.T) {
	// Act
	sut := props.BlackColor

	// Assert
	assert.Equal(t, 0, sut.Red)
	assert.Equal(t, 0, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestRedColor(t *testing.T) {
	// Act
	sut := props.RedColor

	// Assert
	assert.Equal(t, 255, sut.Red)
	assert.Equal(t, 0, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestGreenColor(t *testing.T) {
	// Act
	sut := props.GreenColor

	// Assert
	assert.Equal(t, 0, sut.Red)
	assert.Equal(t, 255, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestBlueColor(t *testing.T) {
	// Act
	blue := props.BlueColor

	// Assert
	assert.Equal(t, 0, blue.Red)
	assert.Equal(t, 0, blue.Green)
	assert.Equal(t, 255, blue.Blue)
}

func TestColor_ToString(t *testing.T) {
	t.Run("when prop is nil, should return empty", func(t *testing.T) {
		// Arrange
		var prop *props.Color

		// Act
		s := prop.ToString()

		// Assert
		assert.Equal(t, "", s)
	})
	t.Run("when prop is filled, should return correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.ColorProp()

		// Act
		s := prop.ToString()

		// Assert
		assert.Equal(t, "RGB(100, 50, 200)", s)
	})
}
