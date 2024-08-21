package text_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.New("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := text.New("code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.NewCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := text.NewCol(12, "code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.NewRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := text.NewRow(10, "code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.NewAutoRow("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := text.NewAutoRow("code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_custom_prop.json")
	})
}

func TestText_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		value := "textValue"
		cell := fixture.CellEntity()
		prop := fixture.TextProp()
		sut := text.New(value, prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddText(value, &cell, &prop)
		sut.SetConfig(&entity.Config{})

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddText", 1)
	})
}

func TestText_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := text.New("textValue")
		fontProp := fixture.FontProp()
		cfg := &entity.Config{
			DefaultFont: &fontProp,
		}

		// Act
		sut.SetConfig(cfg)
	})
}

func TestText_GetHeight(t *testing.T) {
	t.Run("When top margin is sent, should increment row height with top margin", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{Top: 10}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 20.0, height)
	})

	t.Run("When vertical padding is sent, should increment row height with vertical padding", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{VerticalPadding: 5}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 30.0, height)
	})

	t.Run("When font has a height of 2, should return 10", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 10.0, height)
	})
}
