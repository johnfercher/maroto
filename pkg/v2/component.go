package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type Component interface {
	Render(fpdf fpdf.Fpdf, ctx context.Context)
	GetType() string
	Add(components ...Component) Component
	GetStructure() *tree.Node[Structure]
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
