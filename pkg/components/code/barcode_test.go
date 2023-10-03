package code_test

import (
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
)

func TestNewBar(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBar("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).EqualsToJsonFile("new_bar_prop_default.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBar("code", barcodePropFixture())

		// Assert
		test.New(t).Assert(sut.GetStructure()).EqualsToJsonFile("new_bar_custom_prop.json")
	})
}

/*func TestNewBarCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBarCol(12, "code")

		// Assert
		test.New(t).AssertStructure(sut.GetStructure()).EqualsToJsonFile(`{"value":12,"type":"col","details":{"components_size":1,"is_max":false},"nodes":[{"value":"code","type":"barcode","details":{"prop_center":false,"prop_left":0,"prop_percent":100,"prop_proportion_height":0.2,"prop_proportion_width":1,"prop_top":0}}]}`)
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBarCol(12, "code", barcodePropFixture())

		// Assert
		test.New(t).AssertStructure(sut.GetStructure()).EqualsToJsonFile(`{"value":12,"type":"col","details":{"components_size":1,"is_max":false},"nodes":[{"value":"code","type":"barcode","details":{"prop_center":false,"prop_left":10,"prop_percent":98,"prop_proportion_height":3.2,"prop_proportion_width":16,"prop_top":10}}]}`)
	})
}

/func TestNewBarRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewBarRow(10, "code")

		// Assert
		test.New(t).AssertStructure(sut.GetStructure()).EqualsToJsonFile(`{"value":10,"type":"row","details":{"cols_size":1},"nodes":[{"value":0,"type":"col","details":{"components_size":1,"is_max":true},"nodes":[{"value":"code","type":"barcode","details":{"prop_center":false,"prop_left":0,"prop_percent":100,"prop_proportion_height":0.2,"prop_proportion_width":1,"prop_top":0}}]}]}`)
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewBarRow(10, "code", barcodePropFixture())

		// Assert
		test.New(t).AssertStructure(sut.GetStructure()).EqualsToJsonFile(`{"value":10,"type":"row","details":{"cols_size":1},"nodes":[{"value":0,"type":"col","details":{"components_size":1,"is_max":true},"nodes":[{"value":"code","type":"barcode","details":{"prop_center":false,"prop_left":10,"prop_percent":98,"prop_proportion_height":3.2,"prop_proportion_width":16,"prop_top":10}}]}]}`)
	})
}*/

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
