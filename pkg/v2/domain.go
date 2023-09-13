package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type Maroto interface {
	Generate() error
	ForceAddPage(pages ...Page)
	Add(rows ...Row)
	GetStructure() *tree.Node[Structure]
}

type Node interface {
	Render(fpdf fpdf.Fpdf, ctx context.Context)
	GetStructure() *tree.Node[Structure]
}

type Component interface {
	Node
	GetType() string // Just to differentiate from Node
}

type Page interface {
	Node
	Add(rows ...Row) Page
}

type Row interface {
	Node
	Add(cols ...Col)
}

type Col interface {
	Node
	Add(component ...Component) Col
	AddInner(rows ...Row) Col
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
