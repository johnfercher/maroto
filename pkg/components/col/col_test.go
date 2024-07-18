package col_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestNew(t *testing.T) {
	t.Run("when size is not defined, should use is as max", func(t *testing.T) {
		// Act
		c := col.New()

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_zero_size.json")
	})
	t.Run("when size is defined, should not use max", func(t *testing.T) {
		// Act
		c := col.New(12)

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_defined_size.json")
	})
	t.Run("when has component, should retrieve components", func(t *testing.T) {
		// Act
		c := col.New(12).Add(code.NewQr("code"))

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_with_components.json")
	})
	t.Run("when has component, should retrieve components", func(t *testing.T) {
		// Act
		prop := fixture.CellProp()
		c := col.New(12).WithStyle(&prop)

		// Assert
		test.New(t).Assert(c.GetStructure()).Equals("components/cols/new_with_props.json")
	})
}

func TestCol_GetSize(t *testing.T) {
	t.Run("when size defined in creation, should use it", func(t *testing.T) {
		// Arrange
		c := col.New(12)

		// Act
		size := c.GetSize()

		// Assert
		assert.Equal(t, 12, size)
	})
	t.Run("when size not defined in creation, should use config max grid size", func(t *testing.T) {
		// Arrange
		c := col.New()
		c.SetConfig(&entity.Config{MaxGridSize: 14})

		// Act
		size := c.GetSize()

		// Assert
		assert.Equal(t, 14, size)
	})
}

func TestCol_Render(t *testing.T) {
	t.Run("when not createCell, should call provider correctly", func(t *testing.T) {
		// Arrange
		cfg := &entity.Config{}
		cell := fixture.CellEntity()
		style := &props.Cell{}

		provider := mocks.NewProvider(t)

		component := mocks.NewComponent(t)
		component.EXPECT().Render(provider, &cell)
		component.EXPECT().SetConfig(cfg)

		sut := col.New(12).Add(component)
		sut.WithStyle(style)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell, false)

		// Assert
		component.AssertNumberOfCalls(t, "Render", 1)
		component.AssertNumberOfCalls(t, "SetConfig", 1)
	})
	t.Run("when createCell, should call provider correctly", func(t *testing.T) {
		// Arrange
		cfg := &entity.Config{}
		cell := fixture.CellEntity()
		style := &props.Cell{}

		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateCol(cell.Width, cell.Height, cfg, style)

		component := mocks.NewComponent(t)
		component.EXPECT().Render(provider, &cell)
		component.EXPECT().SetConfig(cfg)

		sut := col.New(12).Add(component)
		sut.WithStyle(style)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell, true)

		// Assert
		provider.AssertNumberOfCalls(t, "CreateCol", 1)
		component.AssertNumberOfCalls(t, "Render", 1)
		component.AssertNumberOfCalls(t, "SetConfig", 1)
	})
}

func TestCol_GetHeight(t *testing.T) {
	t.Run("when column has two components, should return the largest", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		cfg := &entity.Config{MaxGridSize: 12}

		provider := mocks.NewProvider(t)

		component := mocks.NewComponent(t)
		component.EXPECT().GetHeight(provider, &cell).Return(10.0)
		component.EXPECT().SetConfig(cfg)

		component2 := mocks.NewComponent(t)
		component2.EXPECT().GetHeight(provider, &cell).Return(15.0)
		component2.EXPECT().SetConfig(cfg)

		sut := col.New(12).Add(component, component2)
		sut.SetConfig(cfg)
		// Act
		height := sut.GetHeight(provider, &cell)

		// Assert

		component.AssertNumberOfCalls(t, "GetHeight", 1)
		assert.Equal(t, height, 15.0)
	})
}
