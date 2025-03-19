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
		cell := fixture.CellEntity()
		prop := fixture.TextProp()
		sut := text.New("textValue", prop)

		provider := mocks.NewProvider(t)
		provider.EXPECT().AddCustomText(&cell, &prop, &entity.SubText{Value: "textValue", Props: props.NewSubText(&prop)})
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

func TestText_GetHeight(t *testing.T) {
	t.Run("When top and bottom margin are sent, should increment row height with margin", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{Top: 10, Bottom: 10}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetTextHeight(&textProp, 100.0, &entity.SubText{Value: "text", Props: props.NewSubText(&textProp)}).Return(20.0)

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 40.0, height)
	})
}

func TestAddSubText(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := text.New("subText 1")
		sut.AddSubText("subText 2")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/add_subtext_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		ps := fixture.TextProp()

		// Act
		sut := text.New("subText 1", ps)
		sut.AddSubText("subText 2", props.NewSubText(&ps))
		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/add_subtext_custom_prop.json")
	})
}
