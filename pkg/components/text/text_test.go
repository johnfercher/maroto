package text_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestNewCustomText(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.NewCustomText([]*entity.SubText{entity.NewSubText("code")})

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_default_prop.json")
	})

	t.Run("when sub text is not sent, should use an empty string", func(t *testing.T) {
		// Act
		sut := text.NewCustomText([]*entity.SubText{})

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_empty.json")
	})
}

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

func TestText_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		prop := fixture.TextProp()
		subs := []*entity.SubText{entity.NewSubText("textValue", props.NewSubText(&prop))}
		sut := text.NewCustomText(subs, prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddCustomText(subs, &cell, &prop)
		sut.SetConfig(&entity.Config{})

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddCustomText", 1)
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
