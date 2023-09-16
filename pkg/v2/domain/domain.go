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
	RegisterHeader(rows ...Row) error
	GetStructure() *tree.Node[Structure]
}

type Node interface {
	Render(provider Provider, cell internal.Cell, config *config.Maroto)
	GetStructure() *tree.Node[Structure]
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
}

type Col interface {
	Node
	Add(nodes ...Node) Col
	AddInner(rows ...Row) Col
	GetSize() int
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
