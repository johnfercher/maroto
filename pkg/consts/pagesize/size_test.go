package pagesize_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/stretchr/testify/assert"
)

func TestGetDimensions(t *testing.T) {
	t.Run("when pageSize is invalid, should return a4", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.Type("invalid")

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 210.0, w)
		assert.Equal(t, 297.0, h)
	})
	t.Run("when pageSize is a1, should return a1", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A1

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 594.0, w)
		assert.Equal(t, 841.0, h)
	})
	t.Run("when pageSize is a2, should return a2", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A2

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 419.9, w)
		assert.Equal(t, 594.0, h)
	})
	t.Run("when pageSize is a3, should return a3", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A3

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 297.0, w)
		assert.Equal(t, 419.9, h)
	})
	t.Run("when pageSize is a4, should return a4", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A4

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 210.0, w)
		assert.Equal(t, 297.0, h)
	})
	t.Run("when pageSize is a5, should return a5", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A5

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 148.4, w)
		assert.Equal(t, 210.0, h)
	})
	t.Run("when pageSize is a6, should return a6", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.A6

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 105.0, w)
		assert.Equal(t, 148.5, h)
	})
	t.Run("when pageSize is letter, should return letter", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.Letter

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 215.9, w)
		assert.Equal(t, 279.4, h)
	})
	t.Run("when pageSize is legal, should return legal", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.Legal

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 215.9, w)
		assert.Equal(t, 355.6, h)
	})
	t.Run("when pageSize is tabloid, should return tabloid", func(t *testing.T) {
		// Arrange
		pageSize := pagesize.Tabloid

		// Act
		w, h := pagesize.GetDimensions(pageSize)

		// Arrange
		assert.Equal(t, 279.4, w)
		assert.Equal(t, 431.8, h)
	})
}
