package providers

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
)

type Div struct {
	_type      string
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

func (d Div) Copy(_type string) Div {
	return Div{
		_type: _type,
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
}

func NewHTML() *html {
	div := Div{
		_type: "body",
		dimensions: Dimensions{
			Width:  300,
			Height: 400,
		},
		margins: margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		},
	}

	return &html{
		div: div,
		cursor: Cursor{
			X: 0,
			Y: 0,
		},
		currentRow: 0,
		currentCol: 0,
	}
}

func (h *html) CreateRow(height float64) {
	/*div := h.div.Copy("div")
	div.dimensions.Height = height

	row := tree.NewNode(div)
	h.rows = append(h.rows, row)
	h.*/
	h.currentRow++
}

func (h *html) CreateCol(width, _ float64) {
	var row *tree.Node[Div]

	rowsLength := len(h.rows)
	if rowsLength > h.currentRow {
		row = h.rows[rowsLength-1]
	} else {
		div := h.div.Copy("div")
		row = tree.NewNode(div)
		h.rows = append(h.rows, row)
	}

	colDiv := row.GetData()
	colDiv.dimensions.Width = width
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

func (h *html) AddText(text string, _ internal.Cell, _ props.Text) {
	col := h.getLastCol()

	textDiv := col.GetData()
	textDiv._type = "span"
	textDiv.content = text
	textNode := tree.NewNode(textDiv)

	col.AddNext(textNode)
}

func (h *html) Generate(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var content string
	for _, row := range h.rows {
		div := row.GetData()
		content += fmt.Sprintf("<%s>%s</%s>\n", div._type, div.content, div._type)
	}

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (h *html) getLastCol() *tree.Node[Div] {
	return h.cols[len(h.cols)-1]
}
