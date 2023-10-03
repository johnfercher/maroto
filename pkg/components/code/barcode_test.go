package code_test

import (
	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/test"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
)

func TestNewBar(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBar("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBar("code", fixture.BarcodeProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_custom_prop.json")
	})
}

func TestNewBarCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBarCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBarCol(12, "code", fixture.BarcodeProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_col_custom_prop.json")
	})
}

func TestNewBarRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBarRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBarRow(10, "code", fixture.BarcodeProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_row_custom_prop.json")
	})
}

func TestBarcode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.BarcodeProp()
		sut := code.NewBar(codeValue, prop)

		provider := &mocks.Provider{}
		provider.EXPECT().AddBarCode(codeValue, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddBarCode", 1)
	})
}

func TestBarcode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := code.NewBar("code")

		// Act
		sut.SetConfig(nil)
	})
}
