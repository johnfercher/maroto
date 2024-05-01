package line_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/test"

	"github.com/johnfercher/maroto/v2/pkg/components/line"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := line.New()

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := line.New(fixture.LineProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := line.NewCol(12)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := line.NewCol(12, fixture.LineProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := line.NewRow(10)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := line.NewRow(10, fixture.LineProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_line_row_custom_prop.json")
	})
}

func TestLine_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.LineProp()
		sut := line.New(prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddLine(&cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddLine", 1)
	})
}

func TestLine_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.LineProp()
		sut := line.New(prop)

		// Act
		sut.SetConfig(nil)
	})
}
