package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewText(t *testing.T) {
	text := gofpdf.NewText(mocks.NewFpdf(t), mocks.NewMath(t), mocks.NewFont(t))

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*gofpdf.text")
}

func TestAdd(t *testing.T) {
	t.Run("when usePageMargin is false, should set the page margin to 0", func(t *testing.T) {
		color := fixture.ColorProp()

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().SetMargins(0.0, 0.0, 0.0)
		pdf.EXPECT().SetMargins(10.0, 10.0, 10.0)
		pdf.EXPECT().GetMargins().Return(10.0, 10.0, 10.0, 10.0)
		pdf.EXPECT().GetStringWidth(mock.Anything).Return(4)
		pdf.EXPECT().Text(mock.Anything, mock.Anything, mock.Anything)

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(mock.Anything, mock.Anything, mock.Anything)
		font.EXPECT().GetHeight(mock.Anything, mock.Anything, mock.Anything).Return(10.0)
		font.EXPECT().GetColor().Return(&color)

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)
		text.Add("test", &entity.Cell{X: 0, Y: 0, Width: 100, Height: 290}, &props.Text{Align: align.Center}, false)

		pdf.AssertNumberOfCalls(t, "SetMargins", 2)
		pdf.AssertNumberOfCalls(t, "GetMargins", 2)
	})
}

func TestGetLinesheight(t *testing.T) {
	t.Run("when a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		textProp := &props.Text{}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("text ").Return(5)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("text text text text", textProp, 11)

		assert.Equal(t, 2, height)
	})

	t.Run("When a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		textProp := &props.Text{BreakLineStrategy: breakline.DashStrategy}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("t").Return(1)
		pdf.EXPECT().GetStringWidth(" ").Return(1)
		pdf.EXPECT().GetStringWidth(" - ").Return(1)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("tttt tttt tttt tttt", textProp, 11)

		assert.Equal(t, 2, height)
	})
}
