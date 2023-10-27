package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestCell_ToMap(t *testing.T) {
	t.Run("when cell is nil, should return nil", func(t *testing.T) {
		// Arrange
		var sut *props.Cell

		// Act
		m := sut.ToMap()

		// Assert
		assert.Nil(t, m)
	})
	t.Run("when cell is filled, should return map filled correctly", func(t *testing.T) {
		// Arrange
		sut := fixture.CellProp()

		// Act
		m := sut.ToMap()

		// Assert
		assert.Equal(t, border.Left, m["prop_border_type"])
		assert.Equal(t, 0.6, m["prop_border_thickness"])
		assert.Equal(t, linestyle.Dashed, m["prop_border_line_style"])
		assert.Equal(t, "RGB(255, 100, 50)", m["prop_background_color"])
		assert.Equal(t, "RGB(200, 80, 60)", m["prop_border_color"])
	})
}
