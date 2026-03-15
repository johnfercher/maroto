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

func TestNewQr(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQr("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQr("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_custom_prop.json")
	})
}

func TestNewQrCol(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQrCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQrCol(12, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_custom_prop.json")
	})
}

func TestNewQrRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQrRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewQrRow(10, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_custom_prop.json")
	})
}

func TestNewAutoQrRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewAutoQrRow("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_qr_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := code.NewAutoQrRow("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_qr_row_custom_prop.json")
	})
}

func TestQrCode_Render(t *testing.T) {
	t.Parallel()
	t.Run("should call provider correctly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewQr(codeValue, prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddQrCode(codeValue, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddQrCode", 1)
	})
}

func TestQrCode_SetConfig(t *testing.T) {
	t.Parallel()
	t.Run("should call correctly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("code unexpectedly panicked: %v", r)
			}
		}()
		sut := code.NewQr("code")

		// Act
		sut.SetConfig(nil)
	})
}

func TestQrCode_GetHeight(t *testing.T) {
	t.Parallel()
	t.Run("When it is not possible to know the dimensions of the qrcode, should return height 0", func(t *testing.T) {
		t.Parallel()
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByQrCode("code").Return(nil, errors.New("anyError2"))

		sut := code.NewQr("code")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 0.0, height)
	})

	t.Run("When the height of the qr code is half the width, should return half the width of the cell", func(t *testing.T) {
		t.Parallel()
		cell := fixture.CellEntity()

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByQrCode("code").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)

		sut := code.NewQr("code")

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}
