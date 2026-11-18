package gopdfwrapper

import (
	"bytes"
	"errors"
	"fmt"
	"hash/maphash"
	"math"
	"slices"
	"strings"
	"sync"
	"unicode"

	"github.com/go-fonts/liberation/liberationmonobold"
	"github.com/go-fonts/liberation/liberationmonobolditalic"
	"github.com/go-fonts/liberation/liberationmonoitalic"
	"github.com/go-fonts/liberation/liberationmonoregular"
	"github.com/go-fonts/liberation/liberationsansbold"
	"github.com/go-fonts/liberation/liberationsansbolditalic"
	"github.com/go-fonts/liberation/liberationsansitalic"
	"github.com/go-fonts/liberation/liberationsansregular"
	"github.com/go-fonts/liberation/liberationserifbold"
	"github.com/go-fonts/liberation/liberationserifbolditalic"
	"github.com/go-fonts/liberation/liberationserifitalic"
	"github.com/go-fonts/liberation/liberationserifregular"
	"github.com/signintech/gopdf"
	xfont "golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// ErrUndefinedFont is returned when a font family and style is neither a custom font nor a default one.
var ErrUndefinedFont = errors.New("undefined font")

const (
	// thousandths converts font units into thousandths of the font size.
	thousandths = 1000.0
	// fixedPointScale converts a 26.6 fixed point value into a float.
	fixedPointScale = 64.0
)

type fontKey struct {
	family string
	style  fontstyle.Type
}

// parsedFont is a TTF parsed once and shared, read-only, by every document that uses it.
type parsedFont struct {
	once      sync.Once
	data      []byte
	container *gopdf.FontContainer
	widths    *glyphWidths
	// underlinePosition and underlineThickness are in thousandths of the font size.
	underlinePosition  float64
	underlineThickness float64
	err                error
}

// parse parses the font the first time it is called, with the bytes returned by data.
func (p *parsedFont) parse(family string, style fontstyle.Type, data func() []byte) {
	p.once.Do(func() {
		p.data = data()
		p.container = &gopdf.FontContainer{}
		p.err = p.container.AddTTFFontDataWithOption(pdfFontName(family, style), p.data, gopdf.TtfOption{
			OnGlyphNotFoundSubstitute: substituteMissingGlyph,
		})
		if p.err != nil {
			p.err = fmt.Errorf("could not parse font %s %s: %w", family, style, p.err)
			return
		}

		font, err := sfnt.Parse(p.data)
		if err != nil {
			p.err = fmt.Errorf("could not parse font %s %s: %w", family, style, err)
			return
		}

		p.widths = newGlyphWidths(font)
		p.underlinePosition, p.underlineThickness = underlineMetrics(font)
	})
}

// pdfFontName is the name a font is registered with in gopdf: each style is registered
// as a regular font of its own (named like gofpdf did, e.g. "arialB"), so every style
// has a distinct BaseFont in the document.
func pdfFontName(family string, style fontstyle.Type) string {
	return family + string(style)
}

// substituteMissingGlyph draws whitespace and control characters (e.g. a line break) that
// the font has no glyph for as a space, any other missing glyph as the .notdef glyph of
// the font (usually a box), so missing characters stay visible, as they were with gofpdf.
func substituteMissingGlyph(r rune) rune {
	if unicode.IsSpace(r) || unicode.IsControl(r) {
		return ' '
	}

	return r
}

// contentKey identifies the bytes of a font file.
type contentKey struct {
	size int
	hash uint64
}

// cacheKey identifies a custom font: the same file may be registered for several families or styles,
// each one is a different font in gopdf.
type cacheKey struct {
	family  string
	style   fontstyle.Type
	content contentKey
}

// fontContent is a copy of the bytes of a font file, shared by every font parsed from it.
type fontContent struct {
	once sync.Once
	data []byte
}

var (
	hashSeed = maphash.MakeSeed()
	// customFonts keeps every custom font parsed in the process, so each font is parsed only once,
	// no matter how many documents use it.
	customFonts sync.Map // cacheKey -> *parsedFont
	// fontContents keeps one copy of the bytes of each font file used by the custom fonts.
	fontContents sync.Map // contentKey -> *fontContent

	defaultFontsOnce sync.Once
	defaultFonts     map[fontKey]*parsedFont
)

// Fonts resolves the fonts available to a document: the custom fonts of its configuration and,
// for the standard PDF families gopdf cannot reference without embedding, Liberation fonts
// (SIL Open Font License), which have the same metrics: Liberation Sans for arial and helvetica,
// Liberation Mono for courier, Liberation Serif for times. symbol and zapfdingbats have no free equivalent and use Liberation Sans.
type Fonts struct {
	custom map[fontKey]entity.CustomFont
	mu     sync.Mutex
	parsed map[fontKey]*parsedFont
}

// NewFonts creates the fonts of a document with its custom fonts.
func NewFonts(customFonts []entity.CustomFont) *Fonts {
	custom := make(map[fontKey]entity.CustomFont, len(customFonts))
	for _, font := range customFonts {
		style, _, _ := splitStyle(font.GetStyle())
		custom[fontKey{family: normalizeFamily(font.GetFamily()), style: style}] = font
	}

	return &Fonts{custom: custom, parsed: map[fontKey]*parsedFont{}}
}

func (f *Fonts) lookup(family string, style fontstyle.Type) (*parsedFont, error) {
	key := fontKey{family: normalizeFamily(family), style: style}

	f.mu.Lock()
	defer f.mu.Unlock()

	if font, ok := f.parsed[key]; ok {
		return font, font.err
	}

	font, err := f.resolve(key)
	if err != nil {
		return nil, err
	}

	f.parsed[key] = font
	return font, font.err
}

func (f *Fonts) resolve(key fontKey) (*parsedFont, error) {
	if custom, ok := f.custom[key]; ok {
		return getCustomFont(key, custom.GetBytes()), nil
	}

	defaultFontsOnce.Do(loadDefaultFonts)
	font, ok := defaultFonts[key]
	if !ok {
		return nil, fmt.Errorf("%w: %s %s", ErrUndefinedFont, key.family, key.style)
	}

	font.parse(key.family, key.style, func() []byte { return font.data })
	return font, nil
}

// getCustomFont returns the parsed font of data, parsing it only the first time the same bytes are
// used for this family and style.
func getCustomFont(key fontKey, data []byte) *parsedFont {
	content := contentKey{size: len(data), hash: maphash.Bytes(hashSeed, data)}
	cKey := cacheKey{family: key.family, style: key.style, content: content}

	cached, ok := customFonts.Load(cKey)
	if !ok {
		cached, _ = customFonts.LoadOrStore(cKey, &parsedFont{})
	}

	font, _ := cached.(*parsedFont)
	font.parse(key.family, key.style, func() []byte { return getFontContent(content, data) })
	if bytes.Equal(font.data, data) {
		return font
	}

	// Hash collision: the font is parsed without being cached.
	font = &parsedFont{}
	font.parse(key.family, key.style, func() []byte { return slices.Clone(data) })
	return font
}

// getFontContent returns the shared copy of data. The bytes are copied once, so the parsed fonts
// do not change if the caller reuses its slice, and a file registered for several styles is kept once.
func getFontContent(key contentKey, data []byte) []byte {
	cached, ok := fontContents.Load(key)
	if !ok {
		cached, _ = fontContents.LoadOrStore(key, &fontContent{})
	}

	content, _ := cached.(*fontContent)
	content.once.Do(func() {
		content.data = slices.Clone(data)
	})

	if !bytes.Equal(content.data, data) {
		return slices.Clone(data)
	}

	return content.data
}

func loadDefaultFonts() {
	sans := map[fontstyle.Type][]byte{
		fontstyle.Normal:     liberationsansregular.TTF,
		fontstyle.Bold:       liberationsansbold.TTF,
		fontstyle.Italic:     liberationsansitalic.TTF,
		fontstyle.BoldItalic: liberationsansbolditalic.TTF,
	}
	mono := map[fontstyle.Type][]byte{
		fontstyle.Normal:     liberationmonoregular.TTF,
		fontstyle.Bold:       liberationmonobold.TTF,
		fontstyle.Italic:     liberationmonoitalic.TTF,
		fontstyle.BoldItalic: liberationmonobolditalic.TTF,
	}
	serif := map[fontstyle.Type][]byte{
		fontstyle.Normal:     liberationserifregular.TTF,
		fontstyle.Bold:       liberationserifbold.TTF,
		fontstyle.Italic:     liberationserifitalic.TTF,
		fontstyle.BoldItalic: liberationserifbolditalic.TTF,
	}

	families := map[string]map[fontstyle.Type][]byte{
		fontfamily.Arial:     sans,
		fontfamily.Helvetica: sans,
		fontfamily.Symbol:    sans,
		fontfamily.ZapBats:   sans,
		fontfamily.Courier:   mono,
		fontfamily.Times:     serif,
	}

	defaultFonts = map[fontKey]*parsedFont{}
	for family, styles := range families {
		for style, data := range styles {
			defaultFonts[fontKey{family: family, style: style}] = &parsedFont{data: data}
		}
	}
}

// splitStyle splits a style into the style of the font file (bold and/or italic)
// and the decorations drawn over the text (underline and/or strikethrough).
func splitStyle(style fontstyle.Type) (fontstyle.Type, bool, bool) {
	s := strings.ToUpper(string(style))
	underline := strings.Contains(s, string(fontstyle.Underline))
	strikethrough := strings.Contains(s, string(fontstyle.Strikethrough))

	bold := strings.Contains(s, string(fontstyle.Bold))
	italic := strings.Contains(s, string(fontstyle.Italic))

	switch {
	case bold && italic:
		return fontstyle.BoldItalic, underline, strikethrough
	case bold:
		return fontstyle.Bold, underline, strikethrough
	case italic:
		return fontstyle.Italic, underline, strikethrough
	default:
		return fontstyle.Normal, underline, strikethrough
	}
}

func normalizeFamily(family string) string {
	return strings.ToLower(family)
}

// underlineMetrics reads the underline position and thickness from the post table,
// falling back to the Helvetica values.
func underlineMetrics(font *sfnt.Font) (float64, float64) {
	const (
		helveticaUnderlinePosition  = -100
		helveticaUnderlineThickness = 50
	)

	post := font.PostTable()
	unitsPerEm := float64(font.UnitsPerEm())
	if post == nil || unitsPerEm == 0 {
		return helveticaUnderlinePosition, helveticaUnderlineThickness
	}

	return float64(post.UnderlinePosition) * thousandths / unitsPerEm,
		float64(post.UnderlineThickness) * thousandths / unitsPerEm
}

// glyphWidths measures text with the advance of each glyph rounded to thousandths of the
// font size, as gofpdf did. gopdf truncates them instead, which makes every text slightly
// narrower than it was and moves line breaks.
type glyphWidths struct {
	font    *sfnt.Font
	buffers sync.Pool
	widths  sync.Map // rune -> int
}

func newGlyphWidths(font *sfnt.Font) *glyphWidths {
	return &glyphWidths{
		font:    font,
		buffers: sync.Pool{New: func() any { return &sfnt.Buffer{} }},
	}
}

// width returns the width of text in thousandths of the font size.
func (g *glyphWidths) width(text string) int {
	total := 0
	for _, r := range text {
		total += g.runeWidth(r)
	}

	return total
}

func (g *glyphWidths) runeWidth(r rune) int {
	if width, ok := g.widths.Load(r); ok {
		w, _ := width.(int)
		return w
	}

	buf, _ := g.buffers.Get().(*sfnt.Buffer)
	defer g.buffers.Put(buf)

	index, err := g.font.GlyphIndex(buf, r)
	if (err != nil || index == 0) && substituteMissingGlyph(r) != r {
		index, err = g.font.GlyphIndex(buf, substituteMissingGlyph(r))
	}
	if err != nil {
		index = 0
	}

	unitsPerEm := g.font.UnitsPerEm()
	advance, err := g.font.GlyphAdvance(buf, index, fixed.I(int(unitsPerEm)), xfont.HintingNone)
	width := 0
	if err == nil && unitsPerEm != 0 {
		width = int(math.Round(float64(advance) / fixedPointScale * thousandths / float64(unitsPerEm)))
	}

	g.widths.Store(r, width)
	return width
}
