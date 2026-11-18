package gopdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gopdf"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestNewFont(t *testing.T) {
	t.Parallel()
	// Arrange
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewGopdf(t)
	fpdf.EXPECT().SetFont(family, style, size)

	// Act
	font := gopdf.NewFont(fpdf, size, family, style)

	// Assert
	assert.NotNil(t, font)
	assert.Equal(t, "*gopdf.Font", fmt.Sprintf("%T", font))
	assert.Equal(t, family, font.GetFamily())
	assert.Equal(t, style, font.GetStyle())
	assert.Equal(t, size, font.GetSize())
	assert.Equal(t, &props.Color{Red: 0, Green: 0, Blue: 0}, font.GetColor())
}

func TestFont_GetHeight(t *testing.T) {
	t.Parallel()
	// Arrange
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewGopdf(t)
	fpdf.EXPECT().SetFont(family, style, size)
	font := gopdf.NewFont(fpdf, size, family, style)

	// Act
	height := font.GetHeight(family, style, size)

	// Assert
	assert.Equal(t, 3.527777777777778, height)
}

func TestFont_SetFamily(t *testing.T) {
	t.Parallel()
	// Arrange
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewGopdf(t)
	fpdf.EXPECT().SetFont(family, style, size)
	fpdf.EXPECT().SetFont(fontfamily.Helvetica, style, size)
	font := gopdf.NewFont(fpdf, size, family, style)

	// Act
	font.SetFamily(fontfamily.Helvetica)

	// Assert
	assert.Equal(t, fontfamily.Helvetica, font.GetFamily())
}

func TestFont_SetStyle(t *testing.T) {
	t.Parallel()
	// Arrange
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewGopdf(t)
	fpdf.EXPECT().SetFont(family, style, size)
	fpdf.EXPECT().SetFont(family, fontstyle.BoldItalic, size)
	font := gopdf.NewFont(fpdf, size, family, style)

	// Act
	font.SetStyle(fontstyle.BoldItalic)

	// Assert
	assert.Equal(t, fontstyle.BoldItalic, font.GetStyle())
}

func TestFont_SetSize(t *testing.T) {
	t.Parallel()
	// Arrange
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewGopdf(t)
	fpdf.EXPECT().SetFont(family, style, size)
	fpdf.EXPECT().SetFont(family, style, 14.0)
	font := gopdf.NewFont(fpdf, size, family, style)

	// Act
	font.SetSize(14.0)

	// Assert
	assert.Equal(t, 14.0, font.GetSize())
}

func TestFont_SetColor(t *testing.T) {
	t.Parallel()
	t.Run("when color is invalid, should not apply color", func(t *testing.T) {
		t.Parallel()
		// Arrange
		size := 10.0
		family := fontfamily.Arial
		style := fontstyle.Bold

		fpdf := mocks.NewGopdf(t)
		fpdf.EXPECT().SetFont(family, style, size)
		font := gopdf.NewFont(fpdf, size, family, style)
		color := &props.Color{Red: 0, Green: 0, Blue: 0}

		// Act
		font.SetColor(nil)

		// Assert
		assert.Equal(t, color, font.GetColor())
	})
	t.Run("when color is valid, should apply color", func(t *testing.T) {
		t.Parallel()
		// Arrange
		size := 10.0
		family := fontfamily.Arial
		style := fontstyle.Bold

		fpdf := mocks.NewGopdf(t)
		fpdf.EXPECT().SetFont(family, style, size)
		fpdf.EXPECT().SetTextColor(200, 200, 200)
		font := gopdf.NewFont(fpdf, size, family, style)
		color := &props.Color{Red: 200, Green: 200, Blue: 200}

		// Act
		font.SetColor(color)

		// Assert
		assert.Equal(t, color, font.GetColor())
	})
}
