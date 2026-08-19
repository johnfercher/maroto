package text_test

import (
	"math"
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
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.New("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.New("code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewCol(12, "code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewRow(10, "code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Parallel()
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewAutoRow("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := text.NewAutoRow("code", fixture.TextProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_custom_prop.json")
	})
}

func TestText_Render(t *testing.T) {
	t.Parallel()
	t.Run("should call provider correctly", func(t *testing.T) {
		t.Parallel()
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
	t.Parallel()
	t.Run("should call correctly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("code unexpectedly panicked: %v", r)
			}
		}()

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
	t.Parallel()
	t.Run("When top margin is sent, should increment row height with top margin", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
		t.Parallel()
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

	t.Run("When rotation is 90, should return rotated bounding box height (= actual string width)", func(t *testing.T) {
		t.Parallel()
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{Rotation: 90}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(1.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)
		provider.EXPECT().GetStringWidth("text", &textProp).Return(40.0)

		// stringWidth=40, h=2: sin(90)=1, cos(90)=0 → 40*1 + 2*0 = 40
		height := sut.GetHeight(provider, &cell)
		assert.InDelta(t, 40.0, height, 0.0001)
	})

	t.Run("When rotation is 45, should return (stringWidth + h) / sqrt(2)", func(t *testing.T) {
		t.Parallel()
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{Rotation: 45}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(1.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)
		provider.EXPECT().GetStringWidth("text", &textProp).Return(40.0)

		// stringWidth=40, h=2: sin(45)=cos(45)=1/sqrt(2) → (40 + 2) / sqrt(2)
		height := sut.GetHeight(provider, &cell)
		expected := (40.0 + 2.0) / math.Sqrt2
		assert.InDelta(t, expected, height, 0.0001)
	})

	t.Run("When string width exceeds content width, should clamp to content width", func(t *testing.T) {
		t.Parallel()
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := props.Text{Rotation: 90}
		textProp.MakeValid(&font)

		sut := text.New("text", textProp)

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(1.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)
		// stringWidth (200) > contentWidth (100) → clamps to 100
		provider.EXPECT().GetStringWidth("text", &textProp).Return(200.0)

		height := sut.GetHeight(provider, &cell)
		assert.InDelta(t, 100.0, height, 0.0001)
	})
}
