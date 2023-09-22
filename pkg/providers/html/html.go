package html

import (
	"bytes"
	"fmt"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/cache"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/providers"
	"github.com/yosssi/gohtml"
)

const (
	BodyTag = "body"
	DivTag  = "div"
	SpanTag = "span"
	RowTag  = "row"
	ColTag  = "col"
)

type Div struct {
	_type      string
	label      string
	content    string
	dimensions Dimensions
	margins    margins
}

type Dimensions struct {
	Width  float64
	Height float64
}

type Cursor struct {
	X float64
	Y float64
}

func (d Div) Copy(_type string, label string) Div {
	return Div{
		_type: _type,
		label: label,
		dimensions: Dimensions{
			Width:  d.dimensions.Width,
			Height: d.dimensions.Height,
		},
		margins: margins{
			Left:   d.margins.Left,
			Right:  d.margins.Right,
			Top:    d.margins.Top,
			Bottom: d.margins.Bottom,
		},
	}
}

type margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

type html struct {
	div        Div
	cursor     Cursor
	rows       []*tree.Node[Div]
	currentRow int
	cols       []*tree.Node[Div]
	currentCol int
	imageCache cache.Cache
}

func New(maroto *config.Config, options ...providers.ProviderOption) *html {
	div := Div{
		_type: BodyTag,
		dimensions: Dimensions{
			Width:  maroto.Dimensions.Width,
			Height: maroto.Dimensions.Height,
		},
		margins: margins{
			Left:   pagesize.MinLeftMargin,
			Right:  pagesize.MinRightMargin,
			Top:    pagesize.MinTopMargin,
			Bottom: pagesize.MinBottomMargin,
		},
	}

	provider := &html{
		div: div,
		cursor: Cursor{
			X: 0,
			Y: 0,
		},
		currentRow: 0,
		currentCol: 0,
	}

	for _, option := range options {
		option(provider)
	}

	return provider
}

func (h *html) CreateRow(_ float64) {
	h.currentRow++
}

func (h *html) SetCache(cache cache.Cache) {
	h.imageCache = cache
}

func (g *html) SetProtection(_ *config.Protection) {

}

func (h *html) CreateCol(width, height float64, config *config.Config, prop *props.Cell) {
	var row *tree.Node[Div]

	rowsLength := len(h.rows)
	if rowsLength > h.currentRow {
		row = h.rows[rowsLength-1]
	} else {
		div := h.div.Copy(DivTag, RowTag)
		row = tree.NewNode(div)
		h.rows = append(h.rows, row)
	}

	colDiv := row.GetData().Copy(DivTag, ColTag)
	colDiv.dimensions.Width = width
	colDiv.dimensions.Height = height
	colNode := tree.NewNode(colDiv)

	h.cols = append(h.cols, colNode)
	row.AddNext(colNode)
}

func (h *html) GetDimensions() (width float64, height float64) {
	return h.div.dimensions.Width, h.div.dimensions.Height
}

func (h *html) GetMargins() (left float64, top float64, right float64, bottom float64) {
	return h.div.margins.Left, h.div.margins.Top, h.div.margins.Right, h.div.margins.Bottom
}

func (h *html) AddText(text string, _ core.Cell, _ props.Text) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag
	textDiv.content = text
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) AddSignature(text string, _ core.Cell, _ props.Text) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag
	textDiv.content = text
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) AddMatrixCode(text string, _ core.Cell, _ props.Rect) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag
	textDiv.content = text
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) AddQrCode(code string, _ core.Cell, _ props.Rect) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag
	textDiv.content = code
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) AddBarCode(code string, _ core.Cell, _ props.Barcode) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag
	textDiv.content = code
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) AddImage(value string, _ core.Cell, _ props.Rect, extension extension.Type) {
	minSize := 20
	if len(value) < minSize {
		minSize = len(value)
	}
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = SpanTag

	image, err := h.imageCache.Load(value, extension)
	if err != nil {
		textDiv.content = "Failed to load image from file"
		textNode := tree.NewNode(textDiv)
		col.AddNext(textNode)
		return
	}

	textDiv.content = image.Value[:minSize]
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) GenerateFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	htmlTemplate := htmlTemplate()
	content := h.getRows()
	html := fmt.Sprintf(htmlTemplate, content)

	_, err = f.WriteString(gohtml.Format(html))
	if err != nil {
		return err
	}

	return nil
}

func (h *html) GenerateBytes() ([]byte, error) {
	htmlTemplate := htmlTemplate()
	content := h.getRows()
	html := fmt.Sprintf(htmlTemplate, content)

	buf := *bytes.NewBufferString(gohtml.Format(html))
	return buf.Bytes(), nil
}

func (h *html) getLastCol() *tree.Node[Div] {
	return h.cols[len(h.cols)-1]
}

func (h *html) getRows() string {
	var content string
	for _, row := range h.rows {
		colContent := h.getCols(row)
		div := row.GetData()
		content += fmt.Sprintf("<%s title=\"%s\" style=\"position:relative; "+
			"width: %0.fmm;\">%s</%s>", div._type, div.label, div.dimensions.Width, colContent, div._type)
	}
	return content
}

func (h *html) getCols(row *tree.Node[Div]) string {
	var content string
	colNodes := row.GetNexts()
	for _, colNode := range colNodes {
		col := colNode.GetData()
		componentContent := h.getComponent(colNode)
		content += fmt.Sprintf("<%s title=\"%s\" style=\"float: left; position:relative; "+
			"border: solid black 1px; width: %0.fmm; height: %0.fmm;\">%s</%s>",
			col._type, col.label, col.dimensions.Width, col.dimensions.Height, componentContent, col._type)
	}

	return content
}

func (h *html) getComponent(col *tree.Node[Div]) string {
	var content string
	componentNodes := col.GetNexts()
	for _, componentNode := range componentNodes {
		component := componentNode.GetData()
		if component._type == "span" {
			content += fmt.Sprintf("<span>%s</span>", component.content)
		}
	}

	return content
}

func htmlTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta name="description" content="Description">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0">
		<link rel="stylesheet" href="css/theme.css">
	</head>
	<body>
		%s
	</body>
</html>
`
}
