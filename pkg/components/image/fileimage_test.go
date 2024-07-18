package image_test

import (
	"errors"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
)

func TestNewFromFile(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFile("path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFile("path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_custom_prop.json")
	})
}

func TestNewFromFileCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFileCol(12, "path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFileCol(12, "path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_col_custom_prop.json")
	})
}

func TestNewFromFileRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromFileRow(10, "path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromFileRow(12, "path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_row_custom_prop.json")
	})
}

func TestNewAutoFromFileRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewAutoFromFileRow("path")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewAutoFromFileRow("path", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_auto_row_custom_prop.json")
	})
}

func TestFileImage_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		path := "path"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := image.NewFromFile(path, prop)

		provider := mocks.NewProvider(t)
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

func TestFileImage_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the file image, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImage("path").Return(nil, errors.New("anyError2"))

		sut := image.NewFromFile("path")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the file image is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImage("path").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)

		sut := image.NewFromFile("path")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}
