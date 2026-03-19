// Package richtext implements creation of text with mixed styles (bold, italic, color, etc.)
// within a single flowing paragraph.
package richtext

import (
	"fmt"
	"strings"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Chunk represents a piece of text with its own style.
type Chunk struct {
	Text  string
	Style props.Text
}

// NewChunk creates a new Chunk with the given text and optional style.
func NewChunk(text string, ps ...props.Text) Chunk {
	style := props.Text{}
	if len(ps) > 0 {
		style = ps[0]
	}
	return Chunk{
		Text:  text,
		Style: style,
	}
}

// RichText is a component that renders multiple chunks of text with different styles
// as a single flowing paragraph with automatic word wrapping.
type RichText struct {
	chunks []Chunk
	config *entity.Config
}

// New creates a new RichText component from the given chunks.
func New(chunks ...Chunk) core.Component {
	return &RichText{
		chunks: chunks,
	}
}

// NewCol creates a RichText wrapped in a Col.
func NewCol(size int, chunks ...Chunk) core.Col {
	rt := New(chunks...)
	return col.New(size).Add(rt)
}

// NewAutoRow creates a RichText wrapped in a Row with automatic height.
func NewAutoRow(chunks ...Chunk) core.Row {
	rt := New(chunks...)
	c := col.New().Add(rt)
	return row.New().Add(c)
}

// NewRow creates a RichText wrapped in a Row with the given height.
func NewRow(height float64, chunks ...Chunk) core.Row {
	rt := New(chunks...)
	c := col.New().Add(rt)
	return row.New(height).Add(c)
}

// SetConfig sets the configuration for the RichText component.
func (r *RichText) SetConfig(config *entity.Config) {
	r.config = config
	for i := range r.chunks {
		r.chunks[i].Style.MakeValid(r.config.DefaultFont)
	}
}

// GetStructure returns the tree structure of the component.
func (r *RichText) GetStructure() *node.Node[core.Structure] {
	details := make(map[string]interface{})
	for i, chunk := range r.chunks {
		details[chunkKey(i, "text")] = chunk.Text
		for k, v := range chunk.Style.ToMap() {
			details[chunkKey(i, k)] = v
		}
	}

	str := core.Structure{
		Type:    "richtext",
		Value:   len(r.chunks),
		Details: details,
	}
	return node.New(str)
}

func chunkKey(index int, key string) string {
	return fmt.Sprintf("chunk_%d_%s", index, key)
}

// word represents a single word that belongs to a specific chunk, carrying its style.
type word struct {
	text  string
	style *props.Text
}

// getWords splits all chunks into individual words, preserving style association.
func (r *RichText) getWords() []word {
	var words []word
	for i := range r.chunks {
		chunk := &r.chunks[i]
		parts := strings.Fields(chunk.Text)
		for _, p := range parts {
			words = append(words, word{text: p, style: &chunk.Style})
		}
	}
	return words
}

// fontPropFromText extracts the Font properties from a Text style.
func fontPropFromText(t *props.Text) *props.Font {
	return &props.Font{
		Family: t.Family,
		Style:  t.Style,
		Size:   t.Size,
		Color:  t.Color,
	}
}

// lineSegment represents a positioned piece of text on a line, ready for rendering.
type lineSegment struct {
	text  string
	style *props.Text
	x     float64
}

// line represents a single line of text composed of multiple segments.
type line struct {
	segments []lineSegment
	width    float64
}

// layoutLines performs the word-wrapping algorithm, breaking words into lines
// that fit within the given width. It returns the laid-out lines and the
// font height of the tallest chunk (used for line spacing).
func (r *RichText) layoutLines(provider core.Provider, colWidth float64) ([]line, float64) {
	words := r.getWords()
	if len(words) == 0 {
		return nil, 0
	}

	// Calculate the maximum font height across all chunks for uniform line spacing.
	maxFontHeight := 0.0
	for i := range r.chunks {
		fh := provider.GetFontHeight(fontPropFromText(&r.chunks[i].Style))
		if fh > maxFontHeight {
			maxFontHeight = fh
		}
	}

	var lines []line
	currentLine := line{}

	for _, w := range words {
		fp := fontPropFromText(w.style)
		wordWidth := provider.GetStringWidth(w.text, fp)

		if len(currentLine.segments) > 0 {
			// Use the larger space width between adjacent fonts so that
			// transitions from a big font to a small font still have
			// a visually adequate gap.
			prevFp := fontPropFromText(currentLine.segments[len(currentLine.segments)-1].style)
			spaceWidth := max(
				provider.GetStringWidth(" ", prevFp),
				provider.GetStringWidth(" ", fp),
			)

			if currentLine.width+spaceWidth+wordWidth > colWidth {
				// Wrap to a new line.
				lines = append(lines, currentLine)
				currentLine = line{}
			} else {
				// Stay on current line, advance past the space.
				currentLine.width += spaceWidth
			}
		}

		x := currentLine.width
		currentLine.segments = append(currentLine.segments, lineSegment{
			text:  w.text,
			style: w.style,
			x:     x,
		})
		currentLine.width = x + wordWidth
	}

	// Don't forget the last line.
	if len(currentLine.segments) > 0 {
		lines = append(lines, currentLine)
	}

	return lines, maxFontHeight
}

// GetHeight calculates the total height the RichText will occupy.
func (r *RichText) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	if len(r.chunks) == 0 {
		return 0
	}

	// Use first chunk's margins for the overall component margins.
	top := r.chunks[0].Style.Top
	bottom := r.chunks[0].Style.Bottom
	left := r.chunks[0].Style.Left
	right := r.chunks[0].Style.Right
	verticalPadding := r.chunks[0].Style.VerticalPadding

	colWidth := cell.Width - left - right
	if colWidth < 0 {
		colWidth = 0
	}

	lines, maxFontHeight := r.layoutLines(provider, colWidth)
	numLines := len(lines)
	if numLines == 0 {
		return top + bottom
	}

	textHeight := float64(numLines)*maxFontHeight + float64(numLines-1)*verticalPadding
	return textHeight + top + bottom
}

// Render draws the RichText into the PDF.
func (r *RichText) Render(provider core.Provider, cell *entity.Cell) {
	if len(r.chunks) == 0 {
		return
	}

	top := r.chunks[0].Style.Top
	left := r.chunks[0].Style.Left
	right := r.chunks[0].Style.Right
	verticalPadding := r.chunks[0].Style.VerticalPadding

	colWidth := cell.Width - left - right
	if colWidth < 0 {
		colWidth = 0
	}

	lines, maxFontHeight := r.layoutLines(provider, colWidth)

	baseX := cell.X + left
	baseY := cell.Y + top

	for i, ln := range lines {
		lineY := baseY + float64(i)*maxFontHeight + float64(i)*verticalPadding

		for _, seg := range ln.segments {
			fp := fontPropFromText(seg.style)
			segWidth := provider.GetStringWidth(seg.text, fp)
			segFontHeight := provider.GetFontHeight(fp)

			// Align baselines: AddText internally adds segFontHeight to Y.
			// We offset Y so that all segments reach the same baseline
			// at lineY + maxFontHeight.
			segCell := &entity.Cell{
				X:      baseX + seg.x,
				Y:      lineY + maxFontHeight - segFontHeight,
				Width:  segWidth,
				Height: maxFontHeight,
			}

			textProp := *seg.style
			// Reset margins since we've already calculated the position.
			textProp.Top = 0
			textProp.Bottom = 0
			textProp.Left = 0
			textProp.Right = 0

			provider.AddText(seg.text, segCell, &textProp)
		}
	}
}
