package col_test

import (
	"github.com/johnfercher/maroto/v2/pkg/test"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
)

func TestNew(t *testing.T) {
	t.Run("when size is not defined, should use is as max", func(t *testing.T) {
		// Act
		c := col.New()

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_zero_size.json")
	})
	t.Run("when size is not defined, should use is as max", func(t *testing.T) {
		// Act
		c := col.New(12)

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_defined_size.json")
	})
}

/*func TestCol_GetSize(t *testing.T) {
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
		c.SetConfig(&entity.Config{MaxGridSize: 14})

		// Act
		size := c.GetSize()

		// Assert
		assert.Equal(t, 14, size)
	})
}
*/
