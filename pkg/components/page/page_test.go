package page_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := page.New()

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := page.New(fixture.PageProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop.json")
	})
	t.Run("when prop is sent and there is rows, should use the provided", func(t *testing.T) {
		// Act
		sut := page.New(fixture.PageProp())

		row := image.NewFromFileRow(10, "path")
		sut.Add(row)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop_and_with_rows.json")
	})
}

func TestPage_Render(t *testing.T) {
	t.Run("when there is no background image and there is no page pattern, should call row render correctly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		prop.Pattern = ""
		cfg := &entity.Config{}

		provider := mocks.NewProvider(t)
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)

		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell)

		// Assert
		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})

	t.Run("when there is background image and there is no page pattern, should call row render and provider correctly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		prop.Pattern = ""
		cfg := &entity.Config{
			BackgroundImage: &entity.Image{
				Bytes:     []byte{1, 2, 3},
				Extension: extension.Jpg,
			},
		}

		rectProp := &props.Rect{}
		rectProp.MakeValid()

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddBackgroundImageFromBytes(cfg.BackgroundImage.Bytes, &cell, rectProp, cfg.BackgroundImage.Extension)
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)

		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddBackgroundImageFromBytes", 1)
		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})
	t.Run("when there is background image and there is page pattern, should call row render and provider correctly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		cfg := &entity.Config{
			BackgroundImage: &entity.Image{
				Bytes:     []byte{1, 2, 3},
				Extension: extension.Jpg,
			},
		}

		rectProp := &props.Rect{}
		rectProp.MakeValid()

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddBackgroundImageFromBytes(cfg.BackgroundImage.Bytes, &cell, rectProp, cfg.BackgroundImage.Extension)
		provider.EXPECT().AddPageNumber(0, 0, &prop, &cell)
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)

		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		// Act
		sut.Render(provider, cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddBackgroundImageFromBytes", 1)
		provider.AssertNumberOfCalls(t, "AddPageNumber", 1)
		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})
}

func TestPage_SetNumber(t *testing.T) {
	t.Run("when called set number, should set correctly", func(t *testing.T) {
		// Arrange
		sut := page.New()

		// Act
		sut.SetNumber(1, 2)

		// Assert
		assert.Equal(t, 1, sut.GetNumber())
	})
}

func TestPage_GetRows(t *testing.T) {
	t.Run("when called get rows, should return rows correctly", func(t *testing.T) {
		// Arrange
		row := mocks.NewRow(t)

		sut := page.New()
		sut.Add(row)

		// Act
		rows := sut.GetRows()

		// Assert
		assert.Equal(t, []core.Row{row}, rows)
	})
}
