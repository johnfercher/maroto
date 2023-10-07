package line

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type line struct {
	config *entity.Config
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

func (l *line) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "lineStyle",
		Details: l.prop.ToMap(),
	}

	return node.New(str)
}

func (l *line) SetConfig(config *entity.Config) {
	l.config = config
}

func (l *line) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddLine(cell, &l.prop)
}
