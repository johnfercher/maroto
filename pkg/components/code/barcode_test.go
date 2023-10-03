package code_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"
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
		sut := code.NewBar("code", barcodePropFixture())

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
		sut := code.NewBarCol(12, "code", barcodePropFixture())

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
		sut := code.NewBarRow(10, "code", barcodePropFixture())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("new_bar_row_custom_prop.json")
	})
}

func barcodePropFixture() props.Barcode {
	return props.Barcode{
		Top:     10,
		Left:    10,
		Percent: 98,
		Proportion: props.Proportion{
			Width:  16,
			Height: 9,
		},
		Center: false,
	}
}
