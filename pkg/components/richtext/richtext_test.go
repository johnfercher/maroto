package richtext_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func defaultConfig() *entity.Config {
	return &entity.Config{DefaultFont: &props.Font{Family: fontfamily.Arial, Style: fontstyle.Normal, Size: 10}}
}

func widths(values map[string]float64) func(string, *props.Font) float64 {
	return func(text string, _ *props.Font) float64 {
		return values[text]
	}
}

func halfOfFontSize(font *props.Font) float64 {
	return font.Size / 2
}

func threeChunks() []richtext.Chunk {
	return []richtext.Chunk{
		richtext.NewChunk("Total: ", props.Text{Top: 2}),
		richtext.NewChunk("$10", props.Text{Style: fontstyle.Bold}),
		richtext.NewChunk(" due."),
	}
}

type renderedText struct {
	text string
	x    float64
	y    float64
}

func captureRendered(rendered *[]renderedText) func(string, *entity.Cell, *props.Text) {
	return func(text string, cell *entity.Cell, _ *props.Text) {
		*rendered = append(*rendered, renderedText{text: text, x: cell.X, y: cell.Y})
	}
}

func TestNewChunk(t *testing.T) {
	t.Parallel()
	t.Run("when props are not sent, should use an empty style", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewChunk("text")

		// Assert
		assert.Equal(t, "text", sut.Text)
		assert.Equal(t, props.Text{}, sut.Style)
	})
	t.Run("when props are sent, should use the first one", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewChunk("text", props.Text{Style: fontstyle.Bold}, props.Text{Style: fontstyle.Italic})

		// Assert
		assert.Equal(t, fontstyle.Bold, sut.Style.Style)
	})
}

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("when chunks are nil, should create a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.New([]richtext.Chunk(nil)...)

		// Assert
		assert.Equal(t, "*richtext.RichText", fmt.Sprintf("%T", sut))
		assert.Equal(t, 0, sut.GetStructure().GetData().Value)
	})
	t.Run("when chunks are empty, should create a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.New([]richtext.Chunk{}...)

		// Assert
		assert.Equal(t, "*richtext.RichText", fmt.Sprintf("%T", sut))
		assert.Equal(t, 0, sut.GetStructure().GetData().Value)
	})
	t.Run("when three chunks are sent, should keep all of them", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.New(threeChunks()...)

		// Assert
		assert.Equal(t, "*richtext.RichText", fmt.Sprintf("%T", sut))
		assert.Equal(t, 3, sut.GetStructure().GetData().Value)
	})
}

func TestNewCol(t *testing.T) {
	t.Parallel()
	t.Run("when chunks are nil, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewCol(12, []richtext.Chunk(nil)...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetData().Value)
	})
	t.Run("when chunks are empty, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewCol(12, []richtext.Chunk{}...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetData().Value)
	})
	t.Run("when three chunks are sent, should wrap them in a column of the given size", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewCol(12, threeChunks()...)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/richtexts/new_richtext_col.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Parallel()
	t.Run("when chunks are nil, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewRow(10, []richtext.Chunk(nil)...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetNexts()[0].GetData().Value)
	})
	t.Run("when chunks are empty, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewRow(10, []richtext.Chunk{}...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetNexts()[0].GetData().Value)
	})
	t.Run("when three chunks are sent, should wrap them in a fixed-height row", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewRow(10, threeChunks()...)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/richtexts/new_richtext_row.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Parallel()
	t.Run("when chunks are nil, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewAutoRow([]richtext.Chunk(nil)...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetNexts()[0].GetData().Value)
	})
	t.Run("when chunks are empty, should wrap a rich text without chunks", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewAutoRow([]richtext.Chunk{}...)

		// Assert
		assert.Equal(t, 0, sut.GetStructure().GetNexts()[0].GetNexts()[0].GetData().Value)
	})
	t.Run("when three chunks are sent, should wrap them in an auto row", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := richtext.NewAutoRow(threeChunks()...)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/richtexts/new_richtext_auto_row.json")
	})
}

func TestRichText_SetConfig(t *testing.T) {
	t.Parallel()
	t.Run("when chunk props are empty, should take the default font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("text"))

		// Act
		sut.SetConfig(defaultConfig())

		// Assert
		details := sut.GetStructure().GetData().Details
		assert.Equal(t, fontfamily.Arial, details["chunk_0_prop_font_family"])
		assert.Equal(t, 10.0, details["chunk_0_prop_font_size"])
	})
}

func TestRichText_GetStructure(t *testing.T) {
	t.Parallel()
	t.Run("when chunks are sent, should expose each chunk text and style", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(
			richtext.NewChunk("Hello ", props.Text{Top: 2, Left: 1}),
			richtext.NewChunk("World", props.Text{Style: fontstyle.Bold}),
		)

		// Act
		data := sut.GetStructure().GetData()

		// Assert
		assert.Equal(t, "richtext", data.Type)
		assert.Equal(t, 2, data.Value)
		assert.Equal(t, "Hello ", data.Details["chunk_0_text"])
		assert.Equal(t, "World", data.Details["chunk_1_text"])
		assert.Equal(t, fontstyle.Bold, data.Details["chunk_1_prop_font_style"])
	})
	t.Run("when there are no chunks, should return an empty structure", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New()

		// Act
		data := sut.GetStructure().GetData()

		// Assert
		assert.Equal(t, 0, data.Value)
		assert.Empty(t, data.Details)
	})
}

func TestRichText_GetHeight(t *testing.T) {
	t.Parallel()
	t.Run("when there are no chunks, should return zero", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New()
		sut.SetConfig(defaultConfig())

		// Act & Assert
		assert.Equal(t, 0.0, sut.GetHeight(mocks.NewProvider(t), &entity.Cell{Width: 100}))
	})
	t.Run("when chunks have no text, should return only the vertical offsets", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("", props.Text{Top: 2, Bottom: 3}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)

		// Act
		height := sut.GetHeight(provider, &entity.Cell{Width: 100})

		// Assert
		assert.Equal(t, 5.0, height)
	})
	t.Run("when line breaks are preserved, should add vertical padding between lines", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("first\nsecond", props.Text{
			Top:                2,
			Bottom:             3,
			VerticalPadding:    1,
			PreserveLineBreaks: true,
		}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"first": 10, "second": 12}))

		// Act
		height := sut.GetHeight(provider, &entity.Cell{Width: 100})

		// Assert
		assert.Equal(t, 16.0, height)
	})
	t.Run("when words exceed the column width, should wrap at spaces", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("aaa bbb", props.Text{Left: 2, Right: 3}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"aaa": 10, " ": 2, "bbb": 10}))

		// Act
		height := sut.GetHeight(provider, &entity.Cell{Width: 20})

		// Assert
		assert.Equal(t, 10.0, height)
	})
	t.Run("when a line mixes font sizes, should use the tallest chunk", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(
			richtext.NewChunk("small "),
			richtext.NewChunk("big", props.Text{Size: 20}),
		)
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).RunAndReturn(halfOfFontSize)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"small": 10, " ": 2, "big": 10}))

		// Act
		height := sut.GetHeight(provider, &entity.Cell{Width: 100})

		// Assert
		assert.Equal(t, 10.0, height)
	})
}

func TestRichText_Render(t *testing.T) {
	t.Parallel()
	t.Run("when there are no chunks, should not render anything", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New()
		sut.SetConfig(defaultConfig())

		// Act
		sut.Render(mocks.NewProvider(t), &entity.Cell{Width: 100})
	})
	t.Run("when chunks touch without spaces, should keep them next to each other", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(
			richtext.NewChunk("red"),
			richtext.NewChunk(", "),
			richtext.NewChunk("blue", props.Text{Style: fontstyle.Bold}),
		)
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"red": 10, ",": 2, " ": 3, "blue": 12}))
		var rendered []renderedText
		provider.EXPECT().AddText(mock.Anything, mock.Anything, mock.Anything).Run(captureRendered(&rendered)).Return()

		// Act
		sut.Render(provider, &entity.Cell{Width: 100})

		// Assert
		assert.Equal(t, []renderedText{{"red", 0, 0}, {",", 10, 0}, {"blue", 15, 0}}, rendered)
	})
	t.Run("when align is center, should center the line", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(
			richtext.NewChunk("ab", props.Text{Align: align.Center}),
			richtext.NewChunk("cd"),
		)
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"ab": 10, "cd": 10}))
		var rendered []renderedText
		provider.EXPECT().AddText(mock.Anything, mock.Anything, mock.Anything).Run(captureRendered(&rendered)).Return()

		// Act
		sut.Render(provider, &entity.Cell{Width: 40})

		// Assert
		assert.Equal(t, []renderedText{{"ab", 10, 0}, {"cd", 20, 0}}, rendered)
	})
	t.Run("when align is right, should push the line to the right edge", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("ab", props.Text{Align: align.Right}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"ab": 10}))
		var rendered []renderedText
		provider.EXPECT().AddText(mock.Anything, mock.Anything, mock.Anything).Run(captureRendered(&rendered)).Return()

		// Act
		sut.Render(provider, &entity.Cell{X: 5, Width: 40})

		// Assert
		assert.Equal(t, []renderedText{{"ab", 35, 0}}, rendered)
	})
	t.Run("when align is justify, should stretch the spaces of every line except the last", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("aa bb cc", props.Text{Align: align.Justify}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"aa": 10, " ": 2, "bb": 10, "cc": 10}))
		var rendered []renderedText
		provider.EXPECT().AddText(mock.Anything, mock.Anything, mock.Anything).Run(captureRendered(&rendered)).Return()

		// Act
		sut.Render(provider, &entity.Cell{Width: 25})

		// Assert
		assert.Equal(t, []renderedText{{"aa", 0, 0}, {"bb", 15, 0}, {"cc", 0, 5}}, rendered)
	})
	t.Run("when chunks use different font sizes, should align them on the same baseline", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(
			richtext.NewChunk("small "),
			richtext.NewChunk("big", props.Text{Size: 20}),
		)
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).RunAndReturn(halfOfFontSize)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"small": 10, " ": 2, "big": 10}))
		var rendered []renderedText
		provider.EXPECT().AddText(mock.Anything, mock.Anything, mock.Anything).Run(captureRendered(&rendered)).Return()

		// Act
		sut.Render(provider, &entity.Cell{Width: 100})

		// Assert
		assert.Equal(t, []renderedText{{"small", 0, 5}, {"big", 12, 0}}, rendered)
	})
	t.Run("when a chunk is rendered, should keep its style and drop the paragraph offsets", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := richtext.New(richtext.NewChunk("ab", props.Text{
			Style: fontstyle.Bold, Align: align.Center, Top: 2, Bottom: 3, Left: 4, Right: 5, VerticalPadding: 1,
		}))
		sut.SetConfig(defaultConfig())

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(mock.Anything).Return(5.0)
		provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(widths(map[string]float64{"ab": 10}))
		var prop *props.Text
		var cell *entity.Cell
		provider.EXPECT().AddText("ab", mock.Anything, mock.Anything).Run(func(_ string, c *entity.Cell, p *props.Text) {
			cell, prop = c, p
		}).Return()

		// Act
		sut.Render(provider, &entity.Cell{Width: 49})

		// Assert
		assert.Equal(t, fontstyle.Bold, prop.Style)
		assert.Equal(t, align.Left, prop.Align)
		assert.Equal(t, 0.0, prop.Top+prop.Left+prop.Right+prop.Bottom+prop.VerticalPadding)
		assert.Equal(t, 4.0+15.0, cell.X)
		assert.Equal(t, 2.0, cell.Y)
	})
}
