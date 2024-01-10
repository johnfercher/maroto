// nolint: dupl
package code_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestNewQr(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewQr("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewQr("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_custom_prop.json")
	})
}

func TestNewQrCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewQrCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewQrCol(12, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_custom_prop.json")
	})
}

func TestNewQrRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewQrRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewQrRow(10, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_custom_prop.json")
	})
}

func TestQrCode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewQr(codeValue, prop)

		provider := &mocks.Provider{}
		provider.EXPECT().AddQrCode(codeValue, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddQrCode", 1)
	})
}

func TestQrCode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := code.NewQr("code")

		// Act
		sut.SetConfig(nil)
	})
}
