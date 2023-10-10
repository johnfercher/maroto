package image_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestNewFromBytes(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromBytes([]byte{1, 2, 3}, extension.Jpg)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromBytes([]byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_custom_prop.json")
	})
}

func TestNewFromBytesCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromBytesCol(12, []byte{1, 2, 3}, extension.Jpg)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromBytesCol(12, []byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Save("components/images/new_image_from_bytes_col_custom_prop.json")
	})
}

func TestNewFromBytesRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := image.NewFromBytesRow(10, []byte{1, 2, 3}, extension.Jpg)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := image.NewFromBytesRow(10, []byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_row_custom_prop.json")
	})
}

func TestBytesImage_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		bytes := []byte{1, 2, 3}
		ext := extension.Jpg
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := image.NewFromBytes(bytes, ext, prop)

		provider := &mocks.Provider{}
		provider.EXPECT().AddImageFromBytes(bytes, &cell, &prop, ext)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddImageFromBytes", 1)
	})
}

func TestBytesImage_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		bytes := []byte{1, 2, 3}
		ext := extension.Jpg
		prop := fixture.RectProp()
		sut := image.NewFromBytes(bytes, ext, prop)

		// Act
		sut.SetConfig(nil)
	})
}
