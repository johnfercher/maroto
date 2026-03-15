package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestText_MakeValid(t *testing.T) {
	t.Parallel()
	t.Run("when family is not defined, should define arial", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Family: "",
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, fontfamily.Arial, prop.Family)
	})
	t.Run("when style is not defined, should define normal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Style: "",
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, fontstyle.Normal, prop.Style)
	})
	t.Run("when size is zero, should define 10.0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Size: 0.0,
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, 10.0, prop.Size)
	})
	t.Run("when align is not defined, should define left", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Align: "",
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, align.Left, prop.Align)
	})
	t.Run("when top is less than 0, should become 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Top: -5.0,
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, 0.0, prop.Top)
	})
	t.Run("when left is less than 0, should become 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Left: -5.0,
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, 0.0, prop.Left)
	})
	t.Run("when right is less than 0, should become 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			Right: -5.0,
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, 0.0, prop.Right)
	})
	t.Run("when vertical padding is less than 0, should become 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Text{
			VerticalPadding: -5.0,
		}

		// Act
		prop.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		// Assert
		assert.Equal(t, 0.0, prop.VerticalPadding)
	})
}
