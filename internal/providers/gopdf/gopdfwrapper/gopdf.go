// nolint:interfacebloat // the interface mirrors the drawing operations maroto needs from gopdf
package gopdfwrapper

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/signintech/gopdf"

	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

const (
	// maxBMPRune is the last character of the Unicode Basic Multilingual Plane.
	maxBMPRune = 0xFFFF
	// defaultLineWidth is the line width every page starts with, in millimeters.
	defaultLineWidth = 0.2
	// conversionUnitMM is the number of points in a millimeter.
	conversionUnitMM = 72.0 / 25.4
)

var (
	// ErrProtectionAfterContent is returned when protection is changed after content was drawn.
	ErrProtectionAfterContent = errors.New("pdf protection must be set before any content is added")
	// ErrNoFontSelected is returned when text is measured or drawn before any font is selected.
	ErrNoFontSelected = errors.New("no font selected")
)

// Gopdf is the set of drawing operations maroto needs from the PDF backend.
// All coordinates and sizes are in millimeters, with the origin in the top-left corner of the page.
//
// Besides drawing, it keeps a cursor (like the one gofpdf had): Ln moves it to the next row,
// CellFormat moves it to the right of the cell it draws and adds a new page when the cell does
// not fit above the bottom margin and the automatic page break is enabled.
type Gopdf interface {
	AddPage()
	PageCount() int
	GetMargins() (left, top, right, bottom float64)
	GetXY() (float64, float64)
	SetHomeXY()
	Ln(h float64)
	CellFormat(w, h float64, bd border.Type, fill bool)

	SetFont(family string, style fontstyle.Type, size float64)
	SetTextColor(r, g, b int)
	GetStringWidth(s string) float64
	Text(x, y float64, txt string)
	LinkString(x, y, w, h float64, url string)

	SetDrawColor(r, g, b int)
	SetFillColor(r, g, b int)
	SetLineWidth(width float64)
	SetDashPattern(dashArray []float64, dashPhase float64)
	Line(x1, y1, x2, y2 float64)
	Rect(x, y, w, h float64, style string)
	Image(img []byte, x, y, w, h float64, flow bool) error

	SetCompression(compress bool)
	SetProtection(protection *entity.Protection)
	SetMetadata(metadata *entity.Metadata)
	Error() error
	Output(w io.Writer) error
}

// Options are the settings a Gopdf is created with.
type Options struct {
	Width                float64
	Height               float64
	Margins              entity.Margins
	DisableAutoPageBreak bool
	Protection           *entity.Protection
	Fonts                *Fonts
}

type color struct {
	r, g, b int
}

type fontState struct {
	family        string
	style         fontstyle.Type
	size          float64
	font          *parsedFont
	underline     bool
	strikethrough bool
}

type document struct {
	pdf     *gopdf.GoPdf
	opts    Options
	fonts   *Fonts
	err     error
	drawn   bool
	pages   int
	x, y    float64
	trigger float64

	registered map[fontKey]bool
	font       *fontState
	textColor  color
	drawColor  color
	fillColor  color
	lineWidth  float64
	dash       []float64
	dashPhase  float64
	info       *gopdf.PdfInfo
	keywords   string
	compress   bool
	passwords  *passwords
}

// New creates a Gopdf with the first page already added.
func New(opts Options) Gopdf {
	opts.Protection = copyProtection(opts.Protection)
	d := &document{
		opts:      opts,
		fonts:     opts.Fonts,
		lineWidth: defaultLineWidth,
	}
	if d.fonts == nil {
		d.fonts = NewFonts(nil)
	}

	d.start()
	d.AddPage()
	d.drawn = false

	return d
}

// AddPage adds a new page, moves the cursor to the top-left margin and restores the
// graphics state (line width, dash pattern and colors) on the new page.
func (d *document) AddPage() {
	d.pdf.AddPage()
	d.pages++
	d.drawn = true
	d.SetHomeXY()

	d.pdf.SetLineWidth(d.lineWidth)
	d.applyDash()
	if d.drawColor != (color{}) {
		d.pdf.SetStrokeColor(toUint8(d.drawColor.r), toUint8(d.drawColor.g), toUint8(d.drawColor.b))
	}
	if d.fillColor != (color{}) {
		d.pdf.SetFillColor(toUint8(d.fillColor.r), toUint8(d.fillColor.g), toUint8(d.fillColor.b))
	}
}

// PageCount returns the number of pages added so far.
func (d *document) PageCount() int {
	return d.pages
}

// GetMargins returns the page margins.
func (d *document) GetMargins() (float64, float64, float64, float64) {
	m := d.opts.Margins
	return m.Left, m.Top, m.Right, m.Bottom
}

// GetXY returns the cursor position.
func (d *document) GetXY() (float64, float64) {
	return d.x, d.y
}

// SetHomeXY moves the cursor to the top-left margin of the current page.
func (d *document) SetHomeXY() {
	d.x = d.opts.Margins.Left
	d.y = d.opts.Margins.Top
}

// Ln moves the cursor to the left margin, h millimeters below.
func (d *document) Ln(h float64) {
	d.x = d.opts.Margins.Left
	d.y += h
}

// CellFormat draws an empty cell of w x h at the cursor, with the given borders and
// filled with the fill color when fill is true, then moves the cursor to its right.
func (d *document) CellFormat(w, h float64, bd border.Type, fill bool) {
	if d.shouldBreakPage(h) {
		x := d.x
		d.AddPage()
		d.x = x
	}

	d.drawn = true
	x, y := d.x, d.y

	if fill {
		d.pdf.RectFromUpperLeftWithStyle(x, y, w, h, "F")
	}

	if bd.HasLeft() {
		d.pdf.Line(x, y, x, y+h)
	}
	if bd.HasTop() {
		d.pdf.Line(x, y, x+w, y)
	}
	if bd.HasRight() {
		d.pdf.Line(x+w, y, x+w, y+h)
	}
	if bd.HasBottom() {
		d.pdf.Line(x, y+h, x+w, y+h)
	}

	d.x += w
}

// SetFont selects the font used to measure and draw text. The font is registered in the
// document the first time it is selected. Underline ("U") and strikethrough ("S") are drawn
// over the text of the regular, bold or italic font file.
func (d *document) SetFont(family string, style fontstyle.Type, size float64) {
	family = normalizeFamily(family)
	if d.font != nil && d.font.family == family && d.font.style == style {
		if d.font.size != size {
			_ = d.pdf.SetFontSize(size)
			d.font.size = size
		}
		return
	}

	base, underline, strikethrough := splitStyle(style)
	key := fontKey{family: family, style: base}

	font, err := d.fonts.lookup(family, base)
	if err != nil {
		d.setError(err)
		return
	}

	name := pdfFontName(family, base)
	if !d.registered[key] {
		err = d.pdf.AddTTFFontFromFontContainer(name, font.container)
		if err != nil {
			d.setError(fmt.Errorf("could not add font %s %s: %w", family, style, err))
			return
		}
		d.registered[key] = true
	}

	err = d.pdf.SetFont(name, "", size)
	if err != nil {
		d.setError(fmt.Errorf("%w: %s %s", err, family, style))
		return
	}

	d.font = &fontState{
		family:        family,
		style:         style,
		size:          size,
		font:          font,
		underline:     underline,
		strikethrough: strikethrough,
	}
}

// SetTextColor defines the color used to draw text.
func (d *document) SetTextColor(r, g, b int) {
	d.textColor = color{r, g, b}
}

// GetStringWidth returns the width of s with the current font.
func (d *document) GetStringWidth(s string) float64 {
	if s == "" {
		return 0
	}
	if d.font == nil {
		d.setError(ErrNoFontSelected)
		return 0
	}

	return float64(d.font.font.widths.width(s)) / thousandths * d.font.size / conversionUnitMM
}

// Text draws txt with its baseline starting at x, y.
func (d *document) Text(x, y float64, txt string) {
	if d.font == nil {
		d.setError(ErrNoFontSelected)
		return
	}

	d.drawn = true
	// gopdf writes the text color as the fill color without restoring it,
	// the graphics state is saved so the fill color of the next shapes is preserved.
	d.pdf.SaveGraphicsState()
	d.pdf.SetTextColor(toUint8(d.textColor.r), toUint8(d.textColor.g), toUint8(d.textColor.b))
	d.pdf.SetXY(x, y)
	err := d.pdf.Text(txt)
	if err != nil {
		d.setError(fmt.Errorf("could not add text: %w", err))
	}

	if d.font.underline || d.font.strikethrough {
		d.decorate(x, y, txt)
	}
	d.pdf.RestoreGraphicsState()
}

// LinkString adds a link to url in the rectangle with top-left corner x, y.
func (d *document) LinkString(x, y, w, h float64, url string) {
	d.drawn = true
	d.pdf.AddExternalLink(url, x, y, w, h)
}

// SetDrawColor defines the color used to stroke lines and borders.
func (d *document) SetDrawColor(r, g, b int) {
	d.drawColor = color{r, g, b}
	d.pdf.SetStrokeColor(toUint8(r), toUint8(g), toUint8(b))
}

// SetFillColor defines the color used to fill cells and rectangles.
func (d *document) SetFillColor(r, g, b int) {
	d.fillColor = color{r, g, b}
	d.pdf.SetFillColor(toUint8(r), toUint8(g), toUint8(b))
}

// SetLineWidth defines the width of lines and borders.
func (d *document) SetLineWidth(width float64) {
	d.lineWidth = width
	d.pdf.SetLineWidth(width)
}

// SetDashPattern defines the dash pattern of lines and borders, a pattern without gaps is a solid line.
func (d *document) SetDashPattern(dashArray []float64, dashPhase float64) {
	d.dash = slices.Clone(dashArray)
	d.dashPhase = dashPhase
	d.applyDash()
}

// Line draws a line from x1, y1 to x2, y2.
func (d *document) Line(x1, y1, x2, y2 float64) {
	d.drawn = true
	d.pdf.Line(x1, y1, x2, y2)
}

// Rect draws a rectangle, style is "D" (draw), "F" (fill) or "FD"/"DF" (both).
func (d *document) Rect(x, y, w, h float64, style string) {
	d.drawn = true
	d.pdf.RectFromUpperLeftWithStyle(x, y, w, h, style)
}

// Image draws img (PNG or JPEG bytes) in the rectangle with top-left corner x, y.
// When flow is true, x and y are ignored: the image is placed at the cursor, a page is added
// first if it does not fit, and the cursor moves below the image.
func (d *document) Image(img []byte, x, y, w, h float64, flow bool) error {
	holder, err := gopdf.ImageHolderByBytes(img)
	if err != nil {
		return fmt.Errorf("could not read image: %w", err)
	}

	if flow {
		if d.shouldBreakPage(h) {
			cx := d.x
			d.AddPage()
			d.x = cx
		}
		y = d.y
		d.y += h
	}

	d.drawn = true
	err = d.pdf.ImageByHolder(holder, x, y, &gopdf.Rect{W: w, H: h})
	if err != nil {
		return fmt.Errorf("could not add image: %w", err)
	}

	return nil
}

// SetCompression enables or disables the compression of the content streams.
func (d *document) SetCompression(compress bool) {
	d.compress = compress
	if compress {
		d.pdf.SetCompressLevel(zlib.DefaultCompression)
		return
	}

	d.pdf.SetNoCompression()
}

// SetProtection encrypts the document. As gopdf encrypts every object when it is created,
// protection can only change before any content is added.
func (d *document) SetProtection(protection *entity.Protection) {
	if protection == nil || equalProtection(protection, d.opts.Protection) {
		return
	}

	if d.drawn {
		d.setError(ErrProtectionAfterContent)
		return
	}

	d.opts.Protection = copyProtection(protection)
	d.start()
	d.AddPage()
	d.drawn = false
}

// SetMetadata defines the document information dictionary. Fields that are nil keep their
// current value.
func (d *document) SetMetadata(metadata *entity.Metadata) {
	if metadata == nil {
		return
	}

	info := gopdf.PdfInfo{}
	if d.info != nil {
		info = *d.info
	}
	if metadata.Author != nil {
		info.Author = infoText(metadata.Author)
	}
	if metadata.Creator != nil {
		info.Creator = infoText(metadata.Creator)
	}
	if metadata.Subject != nil {
		info.Subject = infoText(metadata.Subject)
	}
	if metadata.Title != nil {
		info.Title = infoText(metadata.Title)
	}
	if metadata.CreationDate != nil {
		info.CreationDate = *metadata.CreationDate
	}
	if metadata.KeywordsStr != nil {
		d.keywords = toUTF8(metadata.KeywordsStr)
	}

	d.info = &info
	d.pdf.SetInfo(info)
}

// Error returns the first error that happened while drawing.
func (d *document) Error() error {
	return d.err
}

// Output writes the document to w. It fails with the first error that happened while drawing.
func (d *document) Output(w io.Writer) error {
	if d.err != nil {
		return d.err
	}

	// pdfcpu rejects blank keywords, gofpdf wrote them as they were: they are left out.
	if strings.TrimSpace(d.keywords) == "" {
		_, err := d.pdf.WriteTo(w)
		if err != nil {
			return fmt.Errorf("could not write pdf: %w", err)
		}

		return nil
	}

	var buf bytes.Buffer
	_, err := d.pdf.WriteTo(&buf)
	if err != nil {
		return fmt.Errorf("could not write pdf: %w", err)
	}

	return addKeywords(buf.Bytes(), d.keywords, d.passwords, w)
}

func (d *document) start() {
	d.pdf = &gopdf.GoPdf{}
	d.pdf.Start(gopdf.Config{
		Unit:       gopdf.UnitMM,
		PageSize:   gopdf.Rect{W: d.opts.Width, H: d.opts.Height},
		Protection: d.protectionConfig(),
	})
	d.pdf.SetMargins(d.opts.Margins.Left, d.opts.Margins.Top, d.opts.Margins.Right, d.opts.Margins.Bottom)
	d.trigger = d.opts.Height - d.opts.Margins.Bottom
	d.registered = map[fontKey]bool{}
	d.font = nil
	d.pages = 0
	d.SetCompression(d.compress)
	if d.info != nil {
		d.pdf.SetInfo(*d.info)
	}
}

func (d *document) applyDash() {
	if isSolid(d.dash) {
		d.pdf.SetLineType("solid")
		return
	}

	// gopdf converts the values in place.
	d.pdf.SetCustomLineType(slices.Clone(d.dash), d.dashPhase)
}

// decorate draws underline and strikethrough like gofpdf did: rectangles filled with the
// text color, placed from the underline metrics of the font.
func (d *document) decorate(x, y float64, txt string) {
	const strikethroughFactor = 4

	width := d.GetStringWidth(txt)
	fontSize := d.font.size / conversionUnitMM
	position := d.font.font.underlinePosition / thousandths * fontSize
	thickness := d.font.font.underlineThickness / thousandths * fontSize

	d.pdf.SetFillColor(toUint8(d.textColor.r), toUint8(d.textColor.g), toUint8(d.textColor.b))
	if d.font.underline {
		d.pdf.RectFromUpperLeftWithStyle(x, y-position, width, thickness, "F")
	}
	if d.font.strikethrough {
		d.pdf.RectFromUpperLeftWithStyle(x, y+strikethroughFactor*position, width, thickness, "F")
	}
}

func (d *document) shouldBreakPage(h float64) bool {
	return !d.opts.DisableAutoPageBreak && d.y+h > d.trigger
}

func (d *document) setError(err error) {
	if d.err == nil {
		d.err = err
	}
}

// infoText returns text as gopdf can write it in the information dictionary. gopdf writes each
// character as a single UTF-16 code unit, so characters outside the Basic Multilingual Plane
// (e.g. emoji) are replaced with U+FFFD: as they are, they make an invalid string and, in an
// encrypted document, an unterminated information dictionary.
func infoText(text *entity.Utf8Text) string {
	return strings.Map(func(r rune) rune {
		if r > maxBMPRune {
			return utf8.RuneError
		}

		return r
	}, toUTF8(text))
}

// toUTF8 returns text in UTF-8. Text not marked as UTF-8 is ISO-8859-1, as gofpdf took it.
func toUTF8(text *entity.Utf8Text) string {
	if text.UTF8 {
		return text.Text
	}

	runes := make([]rune, len(text.Text))
	for i := range len(text.Text) {
		runes[i] = rune(text.Text[i])
	}

	return string(runes)
}

func isSolid(dash []float64) bool {
	if len(dash) == 0 {
		return true
	}

	for i := 1; i < len(dash); i += 2 {
		if dash[i] != 0 {
			return false
		}
	}

	return true
}

// protectionConfig converts the protection of the options. When no owner password is given,
// a random one is generated (as gofpdf did), it is kept to be able to post-process the document.
func (d *document) protectionConfig() gopdf.PDFProtectionConfig {
	protection := d.opts.Protection
	if protection == nil {
		d.passwords = nil
		return gopdf.PDFProtectionConfig{}
	}

	owner := protection.OwnerPassword
	if owner == "" {
		owner = randomPassword()
	}
	d.passwords = &passwords{user: protection.UserPassword, owner: owner}

	return gopdf.PDFProtectionConfig{
		UseProtection: true,
		Permissions:   int(protection.Type),
		// gopdf appends padding to the passwords, new slices keep the strings untouched.
		UserPass:  []byte(protection.UserPassword),
		OwnerPass: []byte(owner),
	}
}

func randomPassword() string {
	const size = 24
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// copyProtection returns a copy of protection, so later changes to the caller's value are
// compared with the settings in use instead of being missed.
func copyProtection(protection *entity.Protection) *entity.Protection {
	if protection == nil {
		return nil
	}

	c := *protection

	return &c
}

func equalProtection(a, b *entity.Protection) bool {
	if a == nil || b == nil {
		return a == b
	}

	return *a == *b
}

func toUint8(v int) uint8 {
	return uint8(min(max(v, 0), 255)) //nolint:gosec // clamped to the uint8 range
}
