package checkbox_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/checkbox"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.New("label")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.New("label", fixture.CheckboxProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewCol(12, "label")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewCol(12, "label", fixture.CheckboxProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewRow(10, "label")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewRow(10, "label", fixture.CheckboxProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewAutoRow("label")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := checkbox.NewAutoRow("label", fixture.CheckboxProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/checkboxes/new_checkbox_auto_row_custom_prop.json")
	})
}

func TestCheckbox_Render(t *testing.T) {
	t.Parallel()
	t.Run("should call provider correctly when unchecked", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.CheckboxProp()
		prop.Checked = false
		sut := checkbox.New("label", prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddCheckbox("label", &cell, &prop)
		sut.SetConfig(&entity.Config{})

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddCheckbox", 1)
	})
	t.Run("should call provider correctly when checked", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.CheckboxProp()
		sut := checkbox.New("label", prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddCheckbox("label", &cell, &prop)
		sut.SetConfig(&entity.Config{})

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddCheckbox", 1)
	})
}

func TestCheckbox_SetConfig(t *testing.T) {
	t.Parallel()
	t.Run("should call correctly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("code unexpectedly panicked: %v", r)
			}
		}()

		sut := checkbox.New("label")

		// Act
		sut.SetConfig(&entity.Config{})
	})
}

func TestCheckbox_GetHeight(t *testing.T) {
	t.Parallel()
	t.Run("when default prop is used, should return size", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		sut := checkbox.New("label")
		provider := mocks.NewProvider(t)

		// Act
		height := sut.GetHeight(provider, &cell)

		// Assert
		assert.Equal(t, 5.0, height)
	})
	t.Run("when top is set, should add top to size", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		sut := checkbox.New("label", fixture.CheckboxProp())
		provider := mocks.NewProvider(t)

		// Act
		height := sut.GetHeight(provider, &cell)

		// Assert
		assert.Equal(t, 15.0, height) // Size(10) + Top(5)
	})
}
