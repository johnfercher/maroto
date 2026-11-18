package gopdfwrapper_test

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"strings"
	"testing"
	"time"

	"github.com/go-fonts/liberation/liberationserifregular"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/providers/gopdf/gopdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

func options() gopdfwrapper.Options {
	return gopdfwrapper.Options{
		Width:   210,
		Height:  297,
		Margins: entity.Margins{Left: 10, Top: 15, Right: 10, Bottom: 20},
	}
}

func newSut(t *testing.T, opts gopdfwrapper.Options) gopdfwrapper.Gopdf {
	t.Helper()
	sut := gopdfwrapper.New(opts)
	sut.SetCompression(false)
	return sut
}

func output(t *testing.T, sut gopdfwrapper.Gopdf) string {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, sut.Output(&buf))
	return buf.String()
}

func pngBytes(t *testing.T) []byte {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, png.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 10, 10))))
	return buf.Bytes()
}

func TestNew(t *testing.T) {
	t.Parallel()
	// Act
	sut := gopdfwrapper.New(options())

	// Assert
	x, y := sut.GetXY()
	assert.Equal(t, "*gopdfwrapper.document", fmt.Sprintf("%T", sut))
	assert.Equal(t, 1, sut.PageCount())
	assert.Equal(t, 10.0, x)
	assert.Equal(t, 15.0, y)
	assert.NoError(t, sut.Error())
}

func TestDocument_AddPage(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())
	sut.Ln(50)

	// Act
	sut.AddPage()

	// Assert
	x, y := sut.GetXY()
	assert.Equal(t, 2, sut.PageCount())
	assert.Equal(t, 10.0, x)
	assert.Equal(t, 15.0, y)
}

func TestDocument_GetMargins(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	left, top, right, bottom := sut.GetMargins()

	// Assert
	assert.Equal(t, 10.0, left)
	assert.Equal(t, 15.0, top)
	assert.Equal(t, 10.0, right)
	assert.Equal(t, 20.0, bottom)
}

func TestDocument_Ln(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())
	sut.CellFormat(30, 10, border.None, false)

	// Act
	sut.Ln(10)

	// Assert
	x, y := sut.GetXY()
	assert.Equal(t, 10.0, x)
	assert.Equal(t, 25.0, y)
}

func TestDocument_SetHomeXY(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())
	sut.CellFormat(30, 10, border.None, false)
	sut.Ln(10)

	// Act
	sut.SetHomeXY()

	// Assert
	x, y := sut.GetXY()
	assert.Equal(t, 10.0, x)
	assert.Equal(t, 15.0, y)
}

func TestDocument_CellFormat(t *testing.T) {
	t.Parallel()
	t.Run("when cell fits in the page, should move the cursor to its right", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.CellFormat(30, 10, border.None, false)

		// Assert
		x, y := sut.GetXY()
		assert.Equal(t, 40.0, x)
		assert.Equal(t, 15.0, y)
		assert.Equal(t, 1, sut.PageCount())
	})
	t.Run("when cell does not fit in the page, should add a page and keep x", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.CellFormat(30, 10, border.None, false)
		sut.Ln(262)

		// Act
		sut.CellFormat(30, 10, border.None, false)

		// Assert
		x, y := sut.GetXY()
		assert.Equal(t, 2, sut.PageCount())
		assert.Equal(t, 40.0, x)
		assert.Equal(t, 15.0, y)
	})
	t.Run("when cell does not fit and auto page break is disabled, should keep the page", func(t *testing.T) {
		t.Parallel()
		// Arrange
		opts := options()
		opts.DisableAutoPageBreak = true
		sut := newSut(t, opts)
		sut.Ln(262)

		// Act
		sut.CellFormat(30, 10, border.None, false)

		// Assert
		assert.Equal(t, 1, sut.PageCount())
	})
	t.Run("when fill is true, should draw a filled rectangle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.CellFormat(30, 10, border.None, true)

		// Assert
		assert.Contains(t, output(t, sut), " re f")
	})
	t.Run("when border is full, should draw the four sides", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.CellFormat(30, 10, border.Full, false)

		// Assert
		assert.Equal(t, 4, strings.Count(output(t, sut), " l S"))
	})
	t.Run("when border is left and bottom, should draw two sides", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.CellFormat(30, 10, border.Left|border.Bottom, false)

		// Assert
		assert.Equal(t, 2, strings.Count(output(t, sut), " l S"))
	})
	t.Run("when border is none, should not draw lines", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.CellFormat(30, 10, border.None, false)

		// Assert
		assert.Equal(t, 0, strings.Count(output(t, sut), " l S"))
	})
}

func TestDocument_SetFont(t *testing.T) {
	t.Parallel()
	t.Run("when family is a standard font, should select a liberation font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetFont(fontfamily.Courier, fontstyle.BoldItalic, 10)

		// Assert
		assert.NoError(t, sut.Error())
		assert.InDelta(t, 2.116666, sut.GetStringWidth("a"), 1e-6)
	})
	t.Run("when family is times, should select liberation serif", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetFont(fontfamily.Times, fontstyle.Normal, 10)

		// Assert
		assert.NoError(t, sut.Error())
		assert.InDelta(t, 1.566333, sut.GetStringWidth("a"), 1e-6)
	})
	t.Run("when custom font style is registered in another letter case or order, should find it", func(t *testing.T) {
		t.Parallel()
		// Arrange
		fonts := []entity.CustomFont{
			fixture.TestFont{Family: "serif", Style: fontstyle.Type("ib"), Bytes: liberationserifregular.TTF},
		}
		opts := options()
		opts.Fonts = gopdfwrapper.NewFonts(fonts)
		sut := newSut(t, opts)

		// Act
		sut.SetFont("serif", fontstyle.Type("BIU"), 10)

		// Assert
		assert.NoError(t, sut.Error())
		assert.InDelta(t, 1.566333, sut.GetStringWidth("a"), 1e-6)
	})
	t.Run("when family is unknown, should keep an undefined font error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetFont("unknown", fontstyle.Normal, 10)

		// Assert
		require.ErrorIs(t, sut.Error(), gopdfwrapper.ErrUndefinedFont)
		assert.ErrorIs(t, sut.Output(&bytes.Buffer{}), gopdfwrapper.ErrUndefinedFont)
	})
	t.Run("when family is a custom font, should use the custom font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		fonts := []entity.CustomFont{
			fixture.TestFont{Family: "Serif", Style: fontstyle.Normal, Bytes: liberationserifregular.TTF},
		}
		opts := options()
		opts.Fonts = gopdfwrapper.NewFonts(fonts)
		sut := newSut(t, opts)

		// Act
		sut.SetFont("serif", fontstyle.Normal, 10)

		// Assert
		assert.NoError(t, sut.Error())
		assert.InDelta(t, 1.566333, sut.GetStringWidth("a"), 1e-6)
		assert.Contains(t, output(t, sut), "serif")
	})
	t.Run("when custom font overrides a standard family, should use the custom font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		fonts := []entity.CustomFont{
			fixture.TestFont{Family: fontfamily.Arial, Style: fontstyle.Normal, Bytes: liberationserifregular.TTF},
		}
		opts := options()
		opts.Fonts = gopdfwrapper.NewFonts(fonts)
		sut := newSut(t, opts)

		// Act
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 10)

		// Assert
		assert.InDelta(t, 1.566333, sut.GetStringWidth("a"), 1e-6)
	})
	t.Run("when custom font is not valid, should keep the parse error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		fonts := []entity.CustomFont{
			fixture.TestFont{Family: "broken", Style: fontstyle.Normal, Bytes: []byte("not a font")},
		}
		opts := options()
		opts.Fonts = gopdfwrapper.NewFonts(fonts)
		sut := newSut(t, opts)

		// Act
		sut.SetFont("broken", fontstyle.Normal, 10)

		// Assert
		assert.Error(t, sut.Error())
	})
	t.Run("when only the size changes, should keep the font", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 10)

		// Act
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 20)

		// Assert
		assert.InDelta(t, 3.922888, sut.GetStringWidth("a"), 1e-6)
	})
}

func TestDocument_GetStringWidth(t *testing.T) {
	t.Parallel()
	t.Run("when text is empty, should return zero", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act & Assert
		assert.Zero(t, sut.GetStringWidth(""))
	})
	t.Run("when no font is selected, should return zero and keep an error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		width := sut.GetStringWidth("text")

		// Assert
		assert.Zero(t, width)
		assert.ErrorIs(t, sut.Error(), gopdfwrapper.ErrNoFontSelected)
	})
	t.Run("when font is arial, should measure with the helvetica metrics", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 10)

		// Act
		width := sut.GetStringWidth("This is a longer sentence")

		// Assert
		assert.InDelta(t, 39.214777, width, 1e-6)
	})
	t.Run("when text has a line break, should measure it as a space", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 10)

		// Act & Assert
		assert.Equal(t, sut.GetStringWidth("a b"), sut.GetStringWidth("a\nb"))
	})
}

func TestDocument_Text(t *testing.T) {
	t.Parallel()
	t.Run("when no font is selected, should keep an error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.Text(10, 10, "text")

		// Assert
		assert.ErrorIs(t, sut.Error(), gopdfwrapper.ErrNoFontSelected)
	})
	t.Run("when font is selected, should draw the text with its color", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetFont(fontfamily.Arial, fontstyle.Normal, 10)
		sut.SetTextColor(255, 0, 0)

		// Act
		sut.Text(10, 10, "text")

		// Assert
		pdf := output(t, sut)
		assert.Contains(t, pdf, "BT")
		assert.Contains(t, pdf, "1.000 0.000 0.000 rg")
	})
	t.Run("when style is underline and strikethrough, should draw two decorations", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetFont(fontfamily.Arial, fontstyle.Bold+fontstyle.Underline+fontstyle.Strikethrough, 10)

		// Act
		sut.Text(10, 10, "text")

		// Assert
		assert.NoError(t, sut.Error())
		assert.Equal(t, 2, strings.Count(output(t, sut), " re f"))
	})
}

func TestDocument_LinkString(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.LinkString(10, 10, 20, 5, "https://github.com/johnfercher/maroto")

	// Assert
	assert.Contains(t, output(t, sut), "/URI (https://github.com/johnfercher/maroto)")
}

func TestDocument_SetDrawColor(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.SetDrawColor(255, 0, 0)
	sut.AddPage()

	// Assert
	assert.Equal(t, 2, strings.Count(output(t, sut), "1.000 0.000 0.000 RG"))
}

func TestDocument_SetFillColor(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.SetFillColor(0, 0, 255)
	sut.AddPage()

	// Assert
	assert.Equal(t, 2, strings.Count(output(t, sut), "0.000 0.000 1.000 rg"))
}

func TestDocument_SetLineWidth(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.SetLineWidth(1)

	// Assert
	assert.Contains(t, output(t, sut), "2.83 w")
}

func TestDocument_SetDashPattern(t *testing.T) {
	t.Parallel()
	t.Run("when pattern has gaps, should draw dashed lines", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		dash := []float64{1, 1}

		// Act
		sut.SetDashPattern(dash, 0)

		// Assert
		assert.Contains(t, output(t, sut), "[2.83 2.83] 0.00 d")
		assert.Equal(t, []float64{1, 1}, dash)
	})
	t.Run("when pattern has no gaps, should draw solid lines", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetDashPattern([]float64{1, 0}, 0)

		// Assert
		assert.NotContains(t, output(t, sut), "[2.83")
	})
}

func TestDocument_Line(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.Line(10, 10, 20, 20)

	// Assert
	assert.Equal(t, 1, strings.Count(output(t, sut), " l S"))
}

func TestDocument_Rect(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := newSut(t, options())

	// Act
	sut.Rect(10, 10, 20, 20, "D")

	// Assert
	assert.Contains(t, output(t, sut), " re S")
}

func TestDocument_Image(t *testing.T) {
	t.Parallel()
	t.Run("when bytes are not an image, should return error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		err := sut.Image([]byte{1, 2, 3}, 10, 10, 10, 10, false)

		// Assert
		assert.Error(t, err)
	})
	t.Run("when image is valid, should draw it without moving the cursor", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		err := sut.Image(pngBytes(t), 10, 10, 10, 10, false)

		// Assert
		require.NoError(t, err)
		_, y := sut.GetXY()
		assert.Equal(t, 15.0, y)
		assert.Contains(t, output(t, sut), "/Subtype /Image")
	})
	t.Run("when flow is true, should place the image at the cursor and move it down", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		err := sut.Image(pngBytes(t), 10, 100, 10, 30, true)

		// Assert
		require.NoError(t, err)
		_, y := sut.GetXY()
		assert.Equal(t, 45.0, y)
	})
	t.Run("when flow is true and image does not fit, should add a page", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Ln(250)

		// Act
		err := sut.Image(pngBytes(t), 10, 10, 10, 30, true)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, 2, sut.PageCount())
	})
}

func TestDocument_SetCompression(t *testing.T) {
	t.Parallel()
	t.Run("when compression is enabled, should compress the content streams", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Line(10, 10, 20, 20)

		// Act
		sut.SetCompression(true)

		// Assert
		assert.NotContains(t, output(t, sut), " l S")
	})
	t.Run("when compression is disabled, should not compress the content streams", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Line(10, 10, 20, 20)

		// Act
		sut.SetCompression(false)

		// Assert
		assert.Contains(t, output(t, sut), " l S")
	})
}

func TestDocument_SetProtection(t *testing.T) {
	t.Parallel()
	t.Run("when protection is nil, should not encrypt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetProtection(nil)

		// Assert
		assert.NotContains(t, output(t, sut), "/Encrypt")
	})
	t.Run("when protection is set before content, should encrypt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetProtection(&entity.Protection{Type: protection.Print, UserPassword: "user"})

		// Assert
		assert.Contains(t, output(t, sut), "/Encrypt")
	})
	t.Run("when protection is the one of the options, should not keep an error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		opts := options()
		opts.Protection = &entity.Protection{Type: protection.Print, UserPassword: "user"}
		sut := newSut(t, opts)
		sut.Line(10, 10, 20, 20)

		// Act
		sut.SetProtection(&entity.Protection{Type: protection.Print, UserPassword: "user"})

		// Assert
		assert.NoError(t, sut.Error())
		assert.Contains(t, output(t, sut), "/Encrypt")
	})
	t.Run("when protection changes after content, should keep an error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Line(10, 10, 20, 20)

		// Act
		sut.SetProtection(&entity.Protection{Type: protection.Print, UserPassword: "user"})

		// Assert
		assert.ErrorIs(t, sut.Error(), gopdfwrapper.ErrProtectionAfterContent)
	})
}

func TestDocument_SetMetadata(t *testing.T) {
	t.Parallel()
	t.Run("when metadata is nil, should not write information", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetMetadata(nil)

		// Assert
		assert.NotContains(t, output(t, sut), "/Title")
	})
	t.Run("when metadata is filled, should write every field", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

		// Act
		sut.SetMetadata(&entity.Metadata{
			Author:       &entity.Utf8Text{Text: "author", UTF8: true},
			Creator:      &entity.Utf8Text{Text: "creator", UTF8: true},
			Subject:      &entity.Utf8Text{Text: "subject", UTF8: true},
			Title:        &entity.Utf8Text{Text: "título", UTF8: true},
			KeywordsStr:  &entity.Utf8Text{Text: "keyword", UTF8: true},
			CreationDate: &now,
		})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), nil)
		assert.Equal(t, "author", info.Author)
		assert.Equal(t, "creator", info.Creator)
		assert.Equal(t, "subject", info.Subject)
		assert.Equal(t, "título", info.Title)
		assert.Equal(t, []string{"keyword"}, info.Keywords)
		assert.Contains(t, info.CreationDate, "2024")
	})
	t.Run("when metadata is set again with some fields, should keep the other fields", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.SetMetadata(&entity.Metadata{Author: &entity.Utf8Text{Text: "author", UTF8: true}})

		// Act
		sut.SetMetadata(&entity.Metadata{Title: &entity.Utf8Text{Text: "title", UTF8: true}})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), nil)
		assert.Equal(t, "author", info.Author)
		assert.Equal(t, "title", info.Title)
	})
	t.Run("when text is not marked as utf8, should read it as latin-1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetMetadata(&entity.Metadata{
			Title:       &entity.Utf8Text{Text: "caf\xe9", UTF8: false},
			KeywordsStr: &entity.Utf8Text{Text: "na\xefve", UTF8: false},
		})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), nil)
		assert.Equal(t, "café", info.Title)
		assert.Equal(t, []string{"naïve"}, info.Keywords)
	})
	t.Run("when text has characters outside the basic plane, should write a valid protected pdf", func(t *testing.T) {
		t.Parallel()
		// Arrange
		opts := options()
		opts.Protection = &entity.Protection{Type: protection.Print, UserPassword: "user"}
		sut := newSut(t, opts)

		// Act
		sut.SetMetadata(&entity.Metadata{Title: &entity.Utf8Text{Text: "Report 😀", UTF8: true}})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), &entity.Protection{UserPassword: "user"})
		assert.Equal(t, "Report \uFFFD", info.Title)
	})
	t.Run("when keywords are blank, should write the pdf without them", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		sut.SetMetadata(&entity.Metadata{KeywordsStr: &entity.Utf8Text{Text: "   ", UTF8: true}})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), nil)
		assert.Empty(t, info.Keywords)
	})
	t.Run("when keywords are set on a protected document, should keep it readable with the password", func(t *testing.T) {
		t.Parallel()
		// Arrange
		opts := options()
		opts.Protection = &entity.Protection{Type: protection.Print, UserPassword: "user"}
		sut := newSut(t, opts)

		// Act
		sut.SetMetadata(&entity.Metadata{KeywordsStr: &entity.Utf8Text{Text: "keyword", UTF8: true}})

		// Assert
		info := readInfo(t, []byte(output(t, sut)), &entity.Protection{UserPassword: "user"})
		assert.Equal(t, []string{"keyword"}, info.Keywords)
		assert.True(t, info.Encrypted)
	})
}

func TestDocument_Error(t *testing.T) {
	t.Parallel()
	t.Run("when nothing failed, should return nil", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act & Assert
		assert.NoError(t, sut.Error())
	})
	t.Run("when several operations failed, should return the first error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Text(10, 10, "text")

		// Act
		sut.SetFont("unknown", fontstyle.Normal, 10)

		// Assert
		assert.ErrorIs(t, sut.Error(), gopdfwrapper.ErrNoFontSelected)
	})
}

func TestDocument_Output(t *testing.T) {
	t.Parallel()
	t.Run("when an operation failed, should return its error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())
		sut.Text(10, 10, "text")

		// Act
		err := sut.Output(&bytes.Buffer{})

		// Assert
		assert.ErrorIs(t, err, gopdfwrapper.ErrNoFontSelected)
	})
	t.Run("when nothing failed, should write a pdf", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := newSut(t, options())

		// Act
		pdf := output(t, sut)

		// Assert
		assert.True(t, strings.HasPrefix(pdf, "%PDF-"))
	})
}

func readInfo(t *testing.T, pdf []byte, protection *entity.Protection) *pdfInfo {
	t.Helper()
	conf := model.NewDefaultConfiguration()
	if protection != nil {
		conf.UserPW = protection.UserPassword
	}

	info, err := api.PDFInfo(context.Background(), bytes.NewReader(pdf), "test.pdf", nil, false, conf)
	require.NoError(t, err)

	return &pdfInfo{
		Author:       info.Author,
		Creator:      info.Creator,
		Subject:      info.Subject,
		Title:        info.Title,
		Keywords:     info.Keywords,
		CreationDate: info.CreationDate,
		Encrypted:    info.Encrypted,
	}
}

type pdfInfo struct {
	Author       string
	Creator      string
	Subject      string
	Title        string
	Keywords     []string
	CreationDate string
	Encrypted    bool
}
