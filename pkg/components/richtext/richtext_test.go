package richtext_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func defaultFont() *props.Font {
	return &props.Font{
		Family: "arial",
		Style:  fontstyle.Normal,
		Size:   10,
	}
}

func boldFont() *props.Font {
	return &props.Font{
		Family: "arial",
		Style:  fontstyle.Bold,
		Size:   10,
	}
}

func TestNew(t *testing.T) {
	t.Run("when no chunks are provided, should create component", func(t *testing.T) {
		sut := richtext.New()
		assert.NotNil(t, sut)
	})

	t.Run("when chunks are provided, should create component with chunks", func(t *testing.T) {
		sut := richtext.New(
			richtext.NewChunk("hello "),
			richtext.NewChunk("world", props.Text{Style: fontstyle.Bold}),
		)
		assert.NotNil(t, sut)
	})
}

func TestNewChunk(t *testing.T) {
	t.Run("when style is not sent, should use default", func(t *testing.T) {
		chunk := richtext.NewChunk("text")
		assert.Equal(t, "text", chunk.Text)
	})

	t.Run("when style is sent, should use the provided", func(t *testing.T) {
		chunk := richtext.NewChunk("text", props.Text{Style: fontstyle.Bold, Size: 14})
		assert.Equal(t, "text", chunk.Text)
		assert.Equal(t, fontstyle.Bold, chunk.Style.Style)
		assert.Equal(t, 14.0, chunk.Style.Size)
	})
}

func TestRichText_SetConfig(t *testing.T) {
	t.Run("when config is set, should make chunk styles valid", func(t *testing.T) {
		sut := richtext.New(
			richtext.NewChunk("hello"),
			richtext.NewChunk("world"),
		)

		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		cfg := &entity.Config{DefaultFont: &font}

		sut.SetConfig(cfg)
		// No panic means success — styles have been validated.
	})
}

func TestRichText_GetStructure(t *testing.T) {
	t.Run("when called, should return structure with type richtext", func(t *testing.T) {
		sut := richtext.New(richtext.NewChunk("hello"))
		n := sut.GetStructure()
		assert.Equal(t, "richtext", n.GetData().Type)
	})
}

func TestRichText_GetHeight(t *testing.T) {
	t.Run("when chunks are empty, should return 0", func(t *testing.T) {
		sut := richtext.New()
		provider := mocks.NewProvider(t)

		cell := &entity.Cell{Width: 100, Height: 50}
		height := sut.GetHeight(provider, cell)
		assert.Equal(t, 0.0, height)
	})

	t.Run("when text fits in one line, should return single line height", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		cell := &entity.Cell{Width: 100, Height: 50}
		height := sut.GetHeight(provider, cell)
		// 1 line * 4.0 fontHeight + 0 padding + 0 top + 0 bottom = 4.0
		assert.Equal(t, 4.0, height)
	})

	t.Run("when text exceeds column width, should wrap to two lines", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		// Width 30 forces "world" to wrap (20 + 3 + 20 = 43 > 30)
		cell := &entity.Cell{Width: 30, Height: 50}
		height := sut.GetHeight(provider, cell)
		// 2 lines * 4.0 fontHeight + 1 * 0 padding = 8.0
		assert.Equal(t, 8.0, height)
	})

	t.Run("when top and bottom margins are set, should include them in height", func(t *testing.T) {
		chunk := richtext.NewChunk("hello", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
			Top:    5,
			Bottom: 3,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)

		cell := &entity.Cell{Width: 100, Height: 50}
		height := sut.GetHeight(provider, cell)
		// 1 line * 4.0 + top(5) + bottom(3) = 12.0
		assert.Equal(t, 12.0, height)
	})

	t.Run("when vertical padding is set, should include padding between lines", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family:          "arial",
			Style:           fontstyle.Normal,
			Size:            10,
			VerticalPadding: 2,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		// Force wrap
		cell := &entity.Cell{Width: 30, Height: 50}
		height := sut.GetHeight(provider, cell)
		// 2 lines * 4.0 + 1 * 2 padding = 10.0
		assert.Equal(t, 10.0, height)
	})

	t.Run("when chunks have mixed styles, should use max font height", func(t *testing.T) {
		c1 := richtext.NewChunk("hello", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		c2 := richtext.NewChunk("world", props.Text{
			Family: "arial",
			Style:  fontstyle.Bold,
			Size:   10,
		})
		sut := richtext.New(c1, c2)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fpNormal := defaultFont()
		fpBold := boldFont()

		provider.EXPECT().GetFontHeight(fpNormal).Return(4.0)
		provider.EXPECT().GetFontHeight(fpBold).Return(5.0)
		provider.EXPECT().GetStringWidth("hello", fpNormal).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fpBold).Return(20.0)
		// Space uses max of prev (normal) and curr (bold) font widths
		provider.EXPECT().GetStringWidth(" ", fpNormal).Return(2.5)
		provider.EXPECT().GetStringWidth(" ", fpBold).Return(3.0)

		cell := &entity.Cell{Width: 100, Height: 50}
		height := sut.GetHeight(provider, cell)
		// 1 line * 5.0 (max font height) = 5.0
		assert.Equal(t, 5.0, height)
	})
}

func TestRichText_Render(t *testing.T) {
	t.Run("when chunks are empty, should not call provider", func(t *testing.T) {
		sut := richtext.New()
		provider := mocks.NewProvider(t)
		cell := &entity.Cell{Width: 100, Height: 50}

		sut.Render(provider, cell)
		// No calls expected
	})

	t.Run("when rendering two words, should call AddText for each", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		provider.EXPECT().AddText(mock.AnythingOfType("string"), mock.AnythingOfType("*entity.Cell"), mock.AnythingOfType("*props.Text")).Times(2)

		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		sut.Render(provider, cell)

		provider.AssertNumberOfCalls(t, "AddText", 2)
	})

	t.Run("when rendering on one line, should position words with correct X and Y", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		var calls []struct {
			text string
			cell *entity.Cell
		}

		provider.EXPECT().AddText(
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*entity.Cell"),
			mock.AnythingOfType("*props.Text"),
		).Run(func(text string, cell *entity.Cell, prop *props.Text) {
			calls = append(calls, struct {
				text string
				cell *entity.Cell
			}{text, cell})
		}).Times(2)

		cell := &entity.Cell{X: 5, Y: 10, Width: 100, Height: 50}
		sut.Render(provider, cell)

		assert.Equal(t, 2, len(calls))
		// "hello" at X=5 (cell.X + 0 left margin + 0 word offset)
		assert.Equal(t, "hello", calls[0].text)
		assert.Equal(t, 5.0, calls[0].cell.X)
		assert.Equal(t, 10.0, calls[0].cell.Y)

		// "world" at X=5 + 20 (hello width) + 3 (space) = 28
		assert.Equal(t, "world", calls[1].text)
		assert.Equal(t, 28.0, calls[1].cell.X)
		assert.Equal(t, 10.0, calls[1].cell.Y)
	})

	t.Run("when text exceeds width, should wrap words to next line", func(t *testing.T) {
		chunk := richtext.NewChunk("hello world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(chunk)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fp := defaultFont()

		provider.EXPECT().GetFontHeight(fp).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fp).Return(20.0)
		provider.EXPECT().GetStringWidth("world", fp).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fp).Return(3.0)

		var calls []struct {
			text string
			cell *entity.Cell
		}

		provider.EXPECT().AddText(
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*entity.Cell"),
			mock.AnythingOfType("*props.Text"),
		).Run(func(text string, cell *entity.Cell, prop *props.Text) {
			calls = append(calls, struct {
				text string
				cell *entity.Cell
			}{text, cell})
		}).Times(2)

		// Force wrap: width 30, "hello"=20, " "=3, "world"=20 => 43 > 30
		cell := &entity.Cell{X: 0, Y: 0, Width: 30, Height: 50}
		sut.Render(provider, cell)

		assert.Equal(t, 2, len(calls))
		// "hello" on line 0, Y = 0
		assert.Equal(t, "hello", calls[0].text)
		assert.Equal(t, 0.0, calls[0].cell.X)
		assert.Equal(t, 0.0, calls[0].cell.Y)

		// "world" on line 1, Y = 0 + 4.0 (fontHeight)
		assert.Equal(t, "world", calls[1].text)
		assert.Equal(t, 0.0, calls[1].cell.X)
		assert.Equal(t, 4.0, calls[1].cell.Y)
	})

	t.Run("when chunks have mixed styles, should preserve each chunk style", func(t *testing.T) {
		c1 := richtext.NewChunk("hello ", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		c2 := richtext.NewChunk("bold", props.Text{
			Family: "arial",
			Style:  fontstyle.Bold,
			Size:   10,
		})
		c3 := richtext.NewChunk(" world", props.Text{
			Family: "arial",
			Style:  fontstyle.Normal,
			Size:   10,
		})
		sut := richtext.New(c1, c2, c3)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)
		fpNormal := defaultFont()
		fpBold := boldFont()

		provider.EXPECT().GetFontHeight(fpNormal).Return(4.0)
		provider.EXPECT().GetFontHeight(fpBold).Return(4.0)
		provider.EXPECT().GetStringWidth("hello", fpNormal).Return(20.0)
		provider.EXPECT().GetStringWidth("bold", fpBold).Return(18.0)
		provider.EXPECT().GetStringWidth("world", fpNormal).Return(20.0)
		provider.EXPECT().GetStringWidth(" ", fpBold).Return(3.0)
		provider.EXPECT().GetStringWidth(" ", fpNormal).Return(3.0)

		var calls []struct {
			text string
			cell *entity.Cell
			prop *props.Text
		}

		provider.EXPECT().AddText(
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*entity.Cell"),
			mock.AnythingOfType("*props.Text"),
		).Run(func(text string, cell *entity.Cell, prop *props.Text) {
			calls = append(calls, struct {
				text string
				cell *entity.Cell
				prop *props.Text
			}{text, cell, prop})
		}).Times(3)

		cell := &entity.Cell{X: 0, Y: 0, Width: 200, Height: 50}
		sut.Render(provider, cell)

		assert.Equal(t, 3, len(calls))

		// "hello" normal
		assert.Equal(t, "hello", calls[0].text)
		assert.Equal(t, fontstyle.Normal, calls[0].prop.Style)

		// "bold" bold
		assert.Equal(t, "bold", calls[1].text)
		assert.Equal(t, fontstyle.Bold, calls[1].prop.Style)

		// "world" normal
		assert.Equal(t, "world", calls[2].text)
		assert.Equal(t, fontstyle.Normal, calls[2].prop.Style)
	})
}

func TestRichText_Render_BaselineAlignment(t *testing.T) {
	t.Run("when chunks have mixed font sizes, should align baselines", func(t *testing.T) {
		smallFont := &props.Font{Family: "arial", Style: fontstyle.Normal, Size: 8}
		largeFont := &props.Font{Family: "arial", Style: fontstyle.Bold, Size: 16}

		c1 := richtext.NewChunk("small", props.Text{Family: "arial", Style: fontstyle.Normal, Size: 8})
		c2 := richtext.NewChunk("LARGE", props.Text{Family: "arial", Style: fontstyle.Bold, Size: 16})
		sut := richtext.New(c1, c2)
		font := props.Font{Family: "arial", Style: fontstyle.Normal, Size: 8}
		sut.SetConfig(&entity.Config{DefaultFont: &font})

		provider := mocks.NewProvider(t)

		// maxFontHeight will be 6.0 (from large)
		provider.EXPECT().GetFontHeight(smallFont).Return(3.0)
		provider.EXPECT().GetFontHeight(largeFont).Return(6.0)
		provider.EXPECT().GetStringWidth("small", smallFont).Return(15.0)
		provider.EXPECT().GetStringWidth("LARGE", largeFont).Return(30.0)
		// Space uses max of prev (small) and curr (large) font widths
		provider.EXPECT().GetStringWidth(" ", smallFont).Return(2.0)
		provider.EXPECT().GetStringWidth(" ", largeFont).Return(4.0)

		var calls []struct {
			text string
			cell *entity.Cell
		}

		provider.EXPECT().AddText(
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*entity.Cell"),
			mock.AnythingOfType("*props.Text"),
		).Run(func(text string, cell *entity.Cell, prop *props.Text) {
			calls = append(calls, struct {
				text string
				cell *entity.Cell
			}{text, cell})
		}).Times(2)

		cell := &entity.Cell{X: 0, Y: 0, Width: 200, Height: 50}
		sut.Render(provider, cell)

		assert.Equal(t, 2, len(calls))

		// "small" (fontHeight=3): Y = lineY + maxFontHeight - segFontHeight = 0 + 6 - 3 = 3
		assert.Equal(t, "small", calls[0].text)
		assert.Equal(t, 3.0, calls[0].cell.Y)

		// "LARGE" (fontHeight=6): Y = lineY + maxFontHeight - segFontHeight = 0 + 6 - 6 = 0
		assert.Equal(t, "LARGE", calls[1].text)
		assert.Equal(t, 0.0, calls[1].cell.Y)

		// After AddText adds each segment's fontHeight internally:
		// "small" baseline: 3 + 3 = 6
		// "LARGE" baseline: 0 + 6 = 6
		// Baselines match!
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when called, should create richtext wrapped in col", func(t *testing.T) {
		sut := richtext.NewCol(12, richtext.NewChunk("hello"))
		assert.NotNil(t, sut)
		n := sut.GetStructure()
		assert.Equal(t, "col", n.GetData().Type)
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when called, should create richtext wrapped in row", func(t *testing.T) {
		sut := richtext.NewRow(10, richtext.NewChunk("hello"))
		assert.NotNil(t, sut)
		n := sut.GetStructure()
		assert.Equal(t, "row", n.GetData().Type)
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when called, should create richtext wrapped in auto row", func(t *testing.T) {
		sut := richtext.NewAutoRow(richtext.NewChunk("hello"))
		assert.NotNil(t, sut)
		n := sut.GetStructure()
		assert.Equal(t, "row", n.GetData().Type)
	})
}
