package domain

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
)

type Maroto interface {
	Generate() error
	ForceAddPage(pages ...Page)
	Add(rows ...Row)
	GetStructure() *tree.Node[Structure]
}

type MarotoMetrified interface {
	Maroto
	GenerateWithReport() (*Report, error)
}

type Node interface {
	Render(fpdf Provider, ctx internal.Cell)
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

type Image struct {
	Value     string
	Extension consts.Extension
}
