package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	text := gofpdf.NewText(mocks.NewFpdf(t), mocks.NewMath(t), mocks.NewFont(t))

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*gofpdf.text")
}

func TestGetLinesheight(t *testing.T) {
	t.Run("when the text lines have height 10, it should return 10", func(t *testing.T) {
		textProp := &props.Text{}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})
		font := mocks.NewFont(t)
		pdf := mocks.NewFpdf(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)
		font.EXPECT().GetHeight(textProp.Family, textProp.Style, textProp.Size).Return(10)
		pdf.EXPECT().GetStringWidth("text ").Return(5)
		pdf.EXPECT().GetStringWidth("text text ").Return(5)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)
		subText := entity.SubText{Value: "text text", Props: props.NewSubText(textProp)}

		height := text.GetTextHeight([]*entity.SubText{&subText}, textProp, 11)

		assert.Equal(t, 10.0, height)
	})

	t.Run("when 3 subtexts have height 10, it should return 30", func(t *testing.T) {
		textProp := &props.Text{}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})
		font := mocks.NewFont(t)
		pdf := mocks.NewFpdf(t)

		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)
		font.EXPECT().GetHeight(textProp.Family, textProp.Style, textProp.Size).Return(10)
		pdf.EXPECT().GetStringWidth("text ").Return(5)
		pdf.EXPECT().GetStringWidth("text text ").Return(5)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)
		subText := entity.SubText{Value: "text text text", Props: props.NewSubText(textProp)}

		height := text.GetTextHeight([]*entity.SubText{&subText, &subText}, textProp, 11)
		assert.Equal(t, 30.0, height)
	})

	t.Run("When vertical padding is sent, should increment row height with vertical padding", func(t *testing.T) {
		textProp := props.Text{VerticalPadding: 5}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})
		font := mocks.NewFont(t)
		pdf := mocks.NewFpdf(t)

		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)
		font.EXPECT().GetHeight(textProp.Family, textProp.Style, textProp.Size).Return(10)
		pdf.EXPECT().GetStringWidth("text ").Return(5)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)
		height := text.GetTextHeight([]*entity.SubText{{Value: "text", Props: props.NewSubText(&textProp)}}, &textProp, 10.0)

		// Act
		assert.Equal(t, 15.0, height)
	})
}
