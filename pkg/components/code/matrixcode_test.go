// nolint: dupl
package code_test

import (
	"errors"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrix("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrix("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_custom_prop.json")
	})
}

func TestTestNewMatrixCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrixCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrixCol(12, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_col_custom_prop.json")
	})
}

func TestNewMatrixRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrixRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrixRow(10, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_row_custom_prop.json")
	})
}

func TestAutoNewMatrixRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewAutoMatrixRow("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_matrix_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewAutoMatrixRow("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_matrix_row_custom_prop.json")
	})
}

func TestMatrixCode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewMatrix(codeValue, prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddMatrixCode(codeValue, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddMatrixCode", 1)
	})
}

func TestMatrixCode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := code.NewMatrix("code")

		// Act
		sut.SetConfig(nil)
	})
}

func TestMatrixCode_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the matrix code, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByMatrixCode("code").Return(nil, errors.New("anyError2"))

		sut := code.NewMatrix("code")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the matrix code is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByMatrixCode("code").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)

		sut := code.NewMatrix("code")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}
