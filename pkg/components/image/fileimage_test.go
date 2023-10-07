package image_test

import (
	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
)

func TestNewFromFile(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFile("path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFile("path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_custom_prop.json")
	})
}

func TestNewFromFileCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFileCol(12, "path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFileCol(12, "path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_col_custom_prop.json")
	})
}

func TestNewFromFileRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFileRow(10, "path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFileRow(12, "path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_image_from_file_row_custom_prop.json")
	})
}

func TestFileImage_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		path := "path"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := image.NewFromFile(path, prop)

		provider := &mocks.Provider{}
		provider.EXPECT().AddImageFromFile(path, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddImageFromFile", 1)
	})
}

func TestFileImageSetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		path := "path"
		prop := fixture.RectProp()
		sut := image.NewFromFile(path, prop)

		// Act
		sut.SetConfig(nil)
	})
}
