package col_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/col"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	c := col.New()

	// Assert
	assert.NotNil(t, c)
	assert.Equal(t, "*col.col", fmt.Sprintf("%T", c))
}

func TestCol_GetSize(t *testing.T) {
	t.Run("when size defined in creation, should use it", func(t *testing.T) {
		// Arrange
		c := col.New(12)

		// Act
		size := c.GetSize()

		// Assert
		assert.Equal(t, 12, size)
	})
	t.Run("when size not defined in creation, should use config max grid size", func(t *testing.T) {
		// Arrange
		c := col.New()
		c.SetConfig(&config.Maroto{MaxGridSize: 14})

		// Act
		size := c.GetSize()

		// Assert
		assert.Equal(t, 14, size)
	})
}
