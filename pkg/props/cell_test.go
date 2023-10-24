package props_test

import (
	"testing"

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
}
