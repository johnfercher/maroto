package checkbox_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/checkbox"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := checkbox.New("Male")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
		}

		// Act
		sut := checkbox.New("Male", prop)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := checkbox.NewCol(12, "Male")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
		}

		// Act
		sut := checkbox.NewCol(12, "Male", prop)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := checkbox.NewRow(10, "Male")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
		}

		// Act
		sut := checkbox.NewRow(10, "Male", prop)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := checkbox.NewAutoRow("Male")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
		}

		// Act
		sut := checkbox.NewAutoRow("Male", prop)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_auto_row_custom_prop.json")
	})
}

func TestCheckbox_Render(t *testing.T) {
	t.Run("should call provider.AddCheckbox with correct arguments", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{Checked: true}
		sut := checkbox.New("Male", prop)
		provider := &mocks.Provider{}
		cell := &entity.Cell{}

		provider.On("AddCheckbox", "Male", cell, mock.Anything).Return()

		// Act
		sut.Render(provider, cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddCheckbox", 1)
	})
}

func TestCheckbox_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
		}
		sut := checkbox.New("Male", prop)

		// Act
		sut.SetConfig(nil)
	})
}

func TestCheckbox_GetHeight(t *testing.T) {
	t.Run("should return correct height based on provider font height", func(t *testing.T) {
		// Arrange
		prop := props.Checkbox{BoxSize: 5.0, Top: 1, Bottom: 1}
		sut := checkbox.New("Male", prop)
		provider := &mocks.Provider{}
		cell := &entity.Cell{}

		// Return a font height larger than box size to test max logic
		provider.On("GetFontHeight", mock.Anything).Return(10.0)

		// Act
		height := sut.GetHeight(provider, cell)

		// Assert
		assert.Equal(t, 12.0, height) // 10.0 + 1 + 1
	})
}
