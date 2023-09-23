package line

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type line struct {
	config *config.Config
	prop   props.Line
}

func New(ps ...props.Line) core.Component {
	lineProp := props.Line{}
	if len(ps) > 0 {
		lineProp = ps[0]
	}
	lineProp.MakeValid()

	return &line{
		prop: lineProp,
	}
}

func NewCol(size int, ps ...props.Line) core.Col {
	r := New(ps...)
	return col.New(size).Add(r)
}

func NewRow(height float64, ps ...props.Line) core.Row {
	r := New(ps...)
	c := col.New().Add(r)
	return row.New(height).Add(c)
}

func (l *line) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type: "linestyle",
	}

	return tree.NewNode(str)
}

func (l *line) SetConfig(config *config.Config) {
	l.config = config
}

func (l *line) Render(provider core.Provider, cell core.Cell) {
	provider.AddLine(cell, l.prop)
}