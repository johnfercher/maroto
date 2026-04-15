// Package richtext implements creation of text with mixed styles inside one flowing paragraph.
package richtext

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Chunk represents a styled fragment inside a rich text paragraph.
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

// RichText renders multiple chunks with different styles as a single flowing paragraph.
type RichText struct {
	chunks []Chunk
	config *entity.Config
}

// New creates a RichText component from the given chunks.
func New(chunks ...Chunk) core.Component {
	return &RichText{
		chunks: chunks,
	}
}

// NewCol creates a RichText wrapped in a column.
func NewCol(size int, chunks ...Chunk) core.Col {
	rt := New(chunks...)
	return col.New(size).Add(rt)
}

// NewAutoRow creates a RichText wrapped in an automatic-height row.
func NewAutoRow(chunks ...Chunk) core.Row {
	rt := New(chunks...)
	c := col.New().Add(rt)
	return row.New().Add(c)
}

// NewRow creates a RichText wrapped in a fixed-height row.
func NewRow(height float64, chunks ...Chunk) core.Row {
	rt := New(chunks...)
	c := col.New().Add(rt)
	return row.New(height).Add(c)
}

// SetConfig sets the component configuration.
func (r *RichText) SetConfig(config *entity.Config) {
	r.config = config
	for i := range r.chunks {
		r.chunks[i].Style.MakeValid(r.config.DefaultFont)
	}
}

// GetStructure returns the tree structure of the component.
func (r *RichText) GetStructure() *node.Node[core.Structure] {
	details := make(map[string]any)
	for i, chunk := range r.chunks {
		details[chunkKey(i, "text")] = chunk.Text
		for k, v := range chunk.Style.ToMap() {
			details[chunkKey(i, k)] = v
		}
	}

	return node.New(core.Structure{
		Type:    "richtext",
		Value:   len(r.chunks),
		Details: details,
	})
}

// GetHeight returns the height that the rich text will occupy in the PDF.
func (r *RichText) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	if len(r.chunks) == 0 {
		return 0
	}

	layoutStyle := r.layoutStyle()
	defaultLineHeight := provider.GetFontHeight(fontPropFromText(layoutStyle))
	lines := r.layoutLines(
		provider,
		max(cell.Width-layoutStyle.Left-layoutStyle.Right, 0),
		defaultLineHeight,
	)
	if len(lines) == 0 {
		return layoutStyle.Top + layoutStyle.Bottom
	}

	total := layoutStyle.Top + layoutStyle.Bottom
	for i, line := range lines {
		total += line.height
		if i < len(lines)-1 {
			total += layoutStyle.VerticalPadding
		}
	}

	return total
}

// Render renders the rich text into a PDF context.
func (r *RichText) Render(provider core.Provider, cell *entity.Cell) {
	if len(r.chunks) == 0 {
		return
	}

	layoutStyle := r.layoutStyle()
	colWidth := max(cell.Width-layoutStyle.Left-layoutStyle.Right, 0)
	defaultLineHeight := provider.GetFontHeight(fontPropFromText(layoutStyle))
	lines := r.layoutLines(provider, colWidth, defaultLineHeight)

	xBase := cell.X + layoutStyle.Left
	yBase := cell.Y + layoutStyle.Top
	currentTop := yBase

	for i, line := range lines {
		x := xBase + lineOffsetX(line.width, colWidth, layoutStyle.Align)
		extraSpace := 0.0
		if layoutStyle.Align == align.Justify && i < len(lines)-1 && line.spaceCount > 0 && line.width < colWidth {
			extraSpace = (colWidth - line.width) / float64(line.spaceCount)
		}

		for _, part := range line.parts {
			if part.kind == richTextSpacePart {
				x += part.width + extraSpace
				continue
			}

			prop := *part.style
			prop.Align = align.Left
			prop.Top = 0
			prop.Bottom = 0
			prop.Left = 0
			prop.Right = 0
			prop.VerticalPadding = 0

			segmentHeight := provider.GetFontHeight(fontPropFromText(&prop))
			segmentCell := &entity.Cell{
				X:      x,
				Y:      currentTop + line.height - segmentHeight,
				Width:  part.width,
				Height: line.height,
			}

			provider.AddText(part.text, segmentCell, &prop)
			x += part.width
		}

		currentTop += line.height
		if i < len(lines)-1 {
			currentTop += layoutStyle.VerticalPadding
		}
	}
}

type richTextTokenKind int

const (
	richTextTextToken richTextTokenKind = iota
	richTextSpaceToken
	richTextNewlineToken
)

type richTextToken struct {
	kind  richTextTokenKind
	text  string
	style *props.Text
}

type richTextElementKind int

const (
	richTextClusterElement richTextElementKind = iota
	richTextSpaceElement
	richTextNewlineElement
)

type richTextSegment struct {
	text  string
	style *props.Text
}

type richTextElement struct {
	kind     richTextElementKind
	segments []richTextSegment
	text     string
	style    *props.Text
}

type richTextLinePartKind int

const (
	richTextTextPart richTextLinePartKind = iota
	richTextSpacePart
)

type richTextLinePart struct {
	kind  richTextLinePartKind
	text  string
	style *props.Text
	width float64
}

type richTextLine struct {
	parts      []richTextLinePart
	width      float64
	height     float64
	spaceCount int
}

func (r *RichText) layoutStyle() *props.Text {
	return &r.chunks[0].Style
}

func chunkKey(index int, key string) string {
	return fmt.Sprintf("chunk_%d_%s", index, key)
}

func fontPropFromText(t *props.Text) *props.Font {
	return &props.Font{
		Family: t.Family,
		Style:  t.Style,
		Size:   t.Size,
		Color:  t.Color,
	}
}

func measureText(provider core.Provider, text string, style *props.Text) float64 {
	return provider.GetStringWidth(text, fontPropFromText(style))
}

func normalizeChunkText(text string, preserve bool) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	if preserve {
		return text
	}
	return strings.ReplaceAll(text, "\n", " ")
}

func tokenizeChunk(chunk *Chunk, preserve bool) []richTextToken {
	text := normalizeChunkText(chunk.Text, preserve)
	if text == "" {
		return nil
	}

	tokens := make([]richTextToken, 0)
	var current strings.Builder
	currentKind := richTextTextToken
	hasCurrent := false

	flush := func() {
		if !hasCurrent || current.Len() == 0 {
			return
		}
		tokens = append(tokens, richTextToken{
			kind:  currentKind,
			text:  current.String(),
			style: &chunk.Style,
		})
		current.Reset()
		hasCurrent = false
	}

	for _, r := range text {
		switch {
		case r == '\n':
			flush()
			tokens = append(tokens, richTextToken{
				kind:  richTextNewlineToken,
				text:  "\n",
				style: &chunk.Style,
			})
		case unicode.IsSpace(r):
			if !hasCurrent || currentKind != richTextSpaceToken {
				flush()
				currentKind = richTextSpaceToken
				hasCurrent = true
			}
			current.WriteRune(' ')
		default:
			if !hasCurrent || currentKind != richTextTextToken {
				flush()
				currentKind = richTextTextToken
				hasCurrent = true
			}
			current.WriteRune(r)
		}
	}

	flush()
	return tokens
}

func (r *RichText) elements() []richTextElement {
	if len(r.chunks) == 0 {
		return nil
	}

	preserve := r.layoutStyle().PreserveLineBreaks
	elements := make([]richTextElement, 0)
	cluster := make([]richTextSegment, 0)

	flushCluster := func() {
		if len(cluster) == 0 {
			return
		}
		copied := make([]richTextSegment, len(cluster))
		copy(copied, cluster)
		elements = append(elements, richTextElement{
			kind:     richTextClusterElement,
			segments: copied,
		})
		cluster = cluster[:0]
	}

	for i := range r.chunks {
		chunk := &r.chunks[i]
		for _, token := range tokenizeChunk(chunk, preserve) {
			switch token.kind {
			case richTextTextToken:
				cluster = append(cluster, richTextSegment{
					text:  token.text,
					style: token.style,
				})
			case richTextSpaceToken:
				flushCluster()
				elements = append(elements, richTextElement{
					kind:  richTextSpaceElement,
					text:  token.text,
					style: token.style,
				})
			case richTextNewlineToken:
				flushCluster()
				elements = append(elements, richTextElement{
					kind: richTextNewlineElement,
				})
			}
		}
	}

	flushCluster()
	return elements
}

func finalizeRichTextLine(line richTextLine, defaultHeight float64) richTextLine {
	for len(line.parts) > 0 && line.parts[len(line.parts)-1].kind == richTextSpacePart {
		line.width -= line.parts[len(line.parts)-1].width
		line.spaceCount--
		line.parts = line.parts[:len(line.parts)-1]
	}

	if line.height == 0 {
		line.height = defaultHeight
	}

	return line
}

func (r *RichText) layoutLines(provider core.Provider, colWidth, defaultHeight float64) []richTextLine {
	elements := r.elements()
	if len(elements) == 0 {
		return nil
	}

	lines := make([]richTextLine, 0)
	current := richTextLine{}

	flush := func(forceEmpty bool) {
		if len(current.parts) == 0 {
			if forceEmpty {
				lines = append(lines, richTextLine{height: defaultHeight})
			}
			return
		}

		lines = append(lines, finalizeRichTextLine(current, defaultHeight))
		current = richTextLine{}
	}

	for _, element := range elements {
		switch element.kind {
		case richTextNewlineElement:
			flush(true)
			current = richTextLine{}
		case richTextSpaceElement:
			if len(current.parts) == 0 {
				continue
			}

			width := measureText(provider, element.text, element.style)
			if colWidth > 0 && current.width+width > colWidth {
				flush(false)
				current = richTextLine{}
				continue
			}

			current.parts = append(current.parts, richTextLinePart{
				kind:  richTextSpacePart,
				text:  element.text,
				style: element.style,
				width: width,
			})
			current.width += width
			current.spaceCount++
		case richTextClusterElement:
			clusterParts := make([]richTextLinePart, 0, len(element.segments))
			clusterWidth := 0.0
			clusterHeight := 0.0

			for _, segment := range element.segments {
				width := measureText(provider, segment.text, segment.style)
				height := provider.GetFontHeight(fontPropFromText(segment.style))

				clusterParts = append(clusterParts, richTextLinePart{
					kind:  richTextTextPart,
					text:  segment.text,
					style: segment.style,
					width: width,
				})
				clusterWidth += width
				clusterHeight = max(clusterHeight, height)
			}

			if len(current.parts) > 0 && colWidth > 0 && current.width+clusterWidth > colWidth {
				flush(false)
				current = richTextLine{}
			}

			current.parts = append(current.parts, clusterParts...)
			current.width += clusterWidth
			current.height = max(current.height, clusterHeight)
		}
	}

	if len(current.parts) > 0 {
		lines = append(lines, finalizeRichTextLine(current, defaultHeight))
	}

	return lines
}

func lineOffsetX(lineWidth, colWidth float64, lineAlign align.Type) float64 {
	if colWidth <= lineWidth {
		return 0
	}

	switch lineAlign {
	case align.Center:
		return (colWidth - lineWidth) / 2
	case align.Right:
		return colWidth - lineWidth
	default:
		return 0
	}
}
