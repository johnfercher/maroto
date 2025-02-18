package chart

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type HeatMap struct {
	name   string
	heat   [][]int
	prop   props.HeatMap
	config *entity.Config
}

func NewHeatMap(name string, heat [][]int, ps ...props.HeatMap) core.Component {
	prop := props.HeatMap{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &HeatMap{
		name: name,
		prop: prop,
		heat: heat,
	}
}

func NewHeatMapCol(size int, name string, heat [][]int, ps ...props.HeatMap) core.Col {
	heatMap := NewHeatMap(name, heat, ps...)
	return col.New(size).Add(heatMap)
}

func NewHeatMapRow(height float64, name string, heat [][]int, ps ...props.HeatMap) core.Row {
	heatMap := NewHeatMap(name, heat, ps...)
	c := col.New().Add(heatMap)
	return row.New(height).Add(c)
}

func (b *HeatMap) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddHeatMap(b.heat, cell, &b.prop)
}

func (b *HeatMap) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "heatmap",
		Value:   b.name,
		Details: b.prop.ToMap(),
	}

	return node.New(str)
}

func (b *HeatMap) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	return float64(len(b.heat))
}

func (b *HeatMap) SetConfig(config *entity.Config) {
	b.config = config
}
