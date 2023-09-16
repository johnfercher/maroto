package domain

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

type Maroto interface {
	Generate() (Document, error)
	ForceAddPage(pages ...Page)
	Add(rows ...Row)
	AddCols(rowHeight float64, cols ...Col)
	RegisterHeader(rows ...Row) error
	GetStructure() *tree.Node[Structure]
}

type Component interface {
	SetConfig(config *config.Maroto)
	Render(provider Provider, cell internal.Cell)
	GetStructure() *tree.Node[Structure]
}

type Col interface {
	Component
	Add(components ...Component) Col
	AddInner(rows ...Row) Col
	GetSize() (int, bool)
}

type Row interface {
	Component
	Add(cols ...Col) Row
	GetHeight() float64
}

type Page interface {
	Component
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int)
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
