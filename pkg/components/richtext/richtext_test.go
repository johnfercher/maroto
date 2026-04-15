package richtext_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/richtext"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRichText_GetStructure(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := richtext.New(
		richtext.NewChunk("Hello ", props.Text{Top: 2, Left: 1}),
		richtext.NewChunk("World", props.Text{Style: fontstyle.Bold}),
	)

	// Act
	tree := sut.GetStructure()
	data := tree.GetData()

	// Assert
	assert.Equal(t, "richtext", data.Type)
	assert.Equal(t, 2, data.Value)
	assert.Equal(t, "Hello ", data.Details["chunk_0_text"])
	assert.Equal(t, "World", data.Details["chunk_1_text"])
	assert.Equal(t, fontstyle.Bold, data.Details["chunk_1_prop_font_style"])
}

func TestRichText_GetHeight(t *testing.T) {
	t.Parallel()
	// Arrange
	font := &props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
	sut := richtext.New(
		richtext.NewChunk("first\nsecond", props.Text{
			Top:                2,
			Bottom:             3,
			VerticalPadding:    1,
			PreserveLineBreaks: true,
		}),
	)
	sut.SetConfig(&entity.Config{DefaultFont: font})

	cell := &entity.Cell{Width: 100}

	provider := mocks.NewProvider(t)
	provider.EXPECT().GetFontHeight(mock.Anything).RunAndReturn(func(_ *props.Font) float64 {
		return 5.0
	})
	provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(func(text string, _ *props.Font) float64 {
		switch text {
		case "first":
			return 10.0
		case "second":
			return 12.0
		default:
			return 0.0
		}
	})

	// Act
	height := sut.GetHeight(provider, cell)

	// Assert
	assert.Equal(t, 16.0, height)
}

func TestRichText_RenderPreservesChunkSpacing(t *testing.T) {
	t.Parallel()
	// Arrange
	font := &props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
	sut := richtext.New(
		richtext.NewChunk("red"),
		richtext.NewChunk(", "),
		richtext.NewChunk("blue", props.Text{Style: fontstyle.Bold}),
	)
	sut.SetConfig(&entity.Config{DefaultFont: font})

	cell := &entity.Cell{X: 0, Y: 0, Width: 100}

	provider := mocks.NewProvider(t)
	provider.EXPECT().GetFontHeight(mock.Anything).RunAndReturn(func(_ *props.Font) float64 {
		return 5.0
	})
	provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(func(text string, _ *props.Font) float64 {
		switch text {
		case "red":
			return 10.0
		case ",":
			return 2.0
		case " ":
			return 3.0
		case "blue":
			return 12.0
		default:
			return 0.0
		}
	})

	var rendered []struct {
		text string
		x    float64
		y    float64
	}

	provider.EXPECT().AddText(
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(text string, cell *entity.Cell, prop *props.Text) {
		rendered = append(rendered, struct {
			text string
			x    float64
			y    float64
		}{
			text: text,
			x:    cell.X,
			y:    cell.Y,
		})
	}).Return().Times(3)

	// Act
	sut.Render(provider, cell)

	// Assert
	assert.Equal(t, []struct {
		text string
		x    float64
		y    float64
	}{
		{text: "red", x: 0, y: 0},
		{text: ",", x: 10, y: 0},
		{text: "blue", x: 15, y: 0},
	}, rendered)
}

func TestRichText_RenderRespectsCenterAlign(t *testing.T) {
	t.Parallel()
	// Arrange
	font := &props.Font{Family: "arial", Style: fontstyle.Normal, Size: 10}
	sut := richtext.New(
		richtext.NewChunk("ab", props.Text{Align: align.Center}),
		richtext.NewChunk("cd"),
	)
	sut.SetConfig(&entity.Config{DefaultFont: font})

	cell := &entity.Cell{X: 0, Y: 0, Width: 40}

	provider := mocks.NewProvider(t)
	provider.EXPECT().GetFontHeight(mock.Anything).RunAndReturn(func(_ *props.Font) float64 {
		return 5.0
	})
	provider.EXPECT().GetStringWidth(mock.Anything, mock.Anything).RunAndReturn(func(text string, _ *props.Font) float64 {
		switch text {
		case "ab", "cd":
			return 10.0
		default:
			return 0.0
		}
	})

	var xs []float64
	provider.EXPECT().AddText(
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(_ string, cell *entity.Cell, _ *props.Text) {
		xs = append(xs, cell.X)
	}).Return().Times(2)

	// Act
	sut.Render(provider, cell)

	// Assert
	assert.Equal(t, []float64{10, 20}, xs)
}
