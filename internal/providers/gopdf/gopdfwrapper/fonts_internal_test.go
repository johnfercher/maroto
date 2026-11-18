package gopdfwrapper

import (
	"slices"
	"sync"
	"testing"

	"github.com/go-fonts/liberation/liberationsansregular"
	"github.com/go-fonts/liberation/liberationserifbold"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

func TestFonts_Lookup(t *testing.T) {
	t.Parallel()
	t.Run("when the same font is used by several documents, should parse it once", func(t *testing.T) {
		t.Parallel()
		// Arrange
		first := NewFonts([]entity.CustomFont{
			fixture.TestFont{Family: "cache-once", Style: fontstyle.Bold, Bytes: slices.Clone(liberationserifbold.TTF)},
		})
		second := NewFonts([]entity.CustomFont{
			fixture.TestFont{Family: "Cache-Once", Style: fontstyle.Bold, Bytes: slices.Clone(liberationserifbold.TTF)},
		})

		// Act
		firstFont, firstErr := first.lookup("cache-once", fontstyle.Bold)
		secondFont, secondErr := second.lookup("CACHE-ONCE", fontstyle.Bold)

		// Assert
		require.NoError(t, firstErr)
		require.NoError(t, secondErr)
		assert.Same(t, firstFont, secondFont)
		assert.Same(t, firstFont.container, secondFont.container)
	})
	t.Run("when the caller changes its bytes, should keep the parsed font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := slices.Clone(liberationserifbold.TTF)
		sut := NewFonts([]entity.CustomFont{fixture.TestFont{Family: "cache-copy", Style: fontstyle.Normal, Bytes: data}})
		font, err := sut.lookup("cache-copy", fontstyle.Normal)
		require.NoError(t, err)

		// Act
		clear(data)

		// Assert
		assert.Equal(t, liberationserifbold.TTF, font.data)
	})
	t.Run("when the same file is registered for several styles, should keep its bytes once", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := slices.Clone(liberationserifbold.TTF)
		sut := NewFonts([]entity.CustomFont{
			fixture.TestFont{Family: "cache-styles", Style: fontstyle.Normal, Bytes: data},
			fixture.TestFont{Family: "cache-styles", Style: fontstyle.Bold, Bytes: data},
		})

		// Act
		normal, normalErr := sut.lookup("cache-styles", fontstyle.Normal)
		bold, boldErr := sut.lookup("cache-styles", fontstyle.Bold)

		// Assert
		require.NoError(t, normalErr)
		require.NoError(t, boldErr)
		assert.NotSame(t, normal, bold)
		assert.Same(t, &normal.data[0], &bold.data[0])
	})
	t.Run("when two fonts have the same family and different content, should parse both", func(t *testing.T) {
		t.Parallel()
		// Arrange
		first := NewFonts([]entity.CustomFont{
			fixture.TestFont{Family: "cache-content", Style: fontstyle.Normal, Bytes: liberationserifbold.TTF},
		})
		second := NewFonts([]entity.CustomFont{
			fixture.TestFont{Family: "cache-content", Style: fontstyle.Normal, Bytes: liberationsansregular.TTF},
		})

		// Act
		firstFont, firstErr := first.lookup("cache-content", fontstyle.Normal)
		secondFont, secondErr := second.lookup("cache-content", fontstyle.Normal)

		// Assert
		require.NoError(t, firstErr)
		require.NoError(t, secondErr)
		assert.NotSame(t, firstFont, secondFont)
	})
	t.Run("when a standard family is used, should share the default font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		first := NewFonts(nil)
		second := NewFonts(nil)

		// Act
		firstFont, _ := first.lookup(fontfamily.Helvetica, fontstyle.Italic)
		secondFont, _ := second.lookup(fontfamily.Helvetica, fontstyle.Italic)

		// Assert
		assert.Same(t, firstFont, secondFont)
	})
	t.Run("when the family is unknown, should return an error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := NewFonts(nil)

		// Act
		font, err := sut.lookup("unknown", fontstyle.Normal)

		// Assert
		assert.Nil(t, font)
		assert.ErrorIs(t, err, ErrUndefinedFont)
	})
	t.Run("when documents look fonts up concurrently, should not race", func(t *testing.T) {
		t.Parallel()
		// Arrange
		const documents = 8
		fonts := make([]*parsedFont, documents)
		var wg sync.WaitGroup

		// Act
		for i := range documents {
			wg.Go(func() {
				sut := NewFonts([]entity.CustomFont{
					fixture.TestFont{Family: "cache-race", Style: fontstyle.Normal, Bytes: liberationserifbold.TTF},
				})
				fonts[i], _ = sut.lookup("cache-race", fontstyle.Normal)
				fonts[i].widths.width("concurrent text")
			})
		}
		wg.Wait()

		// Assert
		for _, font := range fonts {
			assert.Same(t, fonts[0], font)
		}
	})
}

func TestGlyphWidths_Width(t *testing.T) {
	t.Parallel()
	t.Run("when glyph advance is fractional, should round it like gofpdf", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut, err := NewFonts(nil).lookup(fontfamily.Arial, fontstyle.Normal)
		require.NoError(t, err)

		// Act & Assert
		// Liberation Sans space is 569/2048 em, 277.83 thousandths, Helvetica has 278.
		assert.Equal(t, 278, sut.widths.width(" "))
	})
	t.Run("when text has several glyphs, should sum their widths", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut, err := NewFonts(nil).lookup(fontfamily.Arial, fontstyle.Normal)
		require.NoError(t, err)

		// Act & Assert
		assert.Equal(t, 556+278+611, sut.widths.width("a T"))
	})
	t.Run("when glyph is a control character, should measure a space", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut, err := NewFonts(nil).lookup(fontfamily.Arial, fontstyle.Normal)
		require.NoError(t, err)

		// Act & Assert
		assert.Equal(t, sut.widths.width(" "), sut.widths.width("\t"))
	})
	t.Run("when glyph is missing, should measure the notdef glyph", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut, err := NewFonts(nil).lookup(fontfamily.Arial, fontstyle.Normal)
		require.NoError(t, err)

		// Act & Assert
		assert.Equal(t, 750, sut.widths.width("\U0001F600"))
	})
}

func TestSplitStyle(t *testing.T) {
	t.Parallel()
	cases := []struct {
		style         fontstyle.Type
		base          fontstyle.Type
		underline     bool
		strikethrough bool
	}{
		{style: fontstyle.Normal, base: fontstyle.Normal},
		{style: fontstyle.Bold, base: fontstyle.Bold},
		{style: fontstyle.Italic, base: fontstyle.Italic},
		{style: "ib", base: fontstyle.BoldItalic},
		{style: fontstyle.Underline, base: fontstyle.Normal, underline: true},
		{style: "BSIU", base: fontstyle.BoldItalic, underline: true, strikethrough: true},
	}

	for _, c := range cases {
		t.Run("when style is "+string(c.style)+", should split it", func(t *testing.T) {
			t.Parallel()
			// Act
			base, underline, strikethrough := splitStyle(c.style)

			// Assert
			assert.Equal(t, c.base, base)
			assert.Equal(t, c.underline, underline)
			assert.Equal(t, c.strikethrough, strikethrough)
		})
	}
}

func TestSubstituteMissingGlyph(t *testing.T) {
	t.Parallel()
	t.Run("when rune is a line break, should return a space", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.Equal(t, ' ', substituteMissingGlyph('\n'))
	})
	t.Run("when rune is a control character, should return a space", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.Equal(t, ' ', substituteMissingGlyph('\u0007'))
	})
	t.Run("when rune is printable, should keep it to draw the notdef glyph", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.Equal(t, 'ሰ', substituteMissingGlyph('ሰ'))
	})
}

func TestIsSolid(t *testing.T) {
	t.Parallel()
	t.Run("when pattern is empty, should be solid", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.True(t, isSolid(nil))
	})
	t.Run("when gaps are zero, should be solid", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.True(t, isSolid([]float64{1, 0}))
	})
	t.Run("when a gap is not zero, should not be solid", func(t *testing.T) {
		t.Parallel()
		// Act & Assert
		assert.False(t, isSolid([]float64{1, 0, 2, 1}))
	})
}
