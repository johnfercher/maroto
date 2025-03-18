package chart

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type TimeSeries struct {
	timeSeriesList []entity.TimeSeries
	config         *entity.Config
	prop           props.Chart
}

func NewTimeSeries(timeSeriesList []entity.TimeSeries, ps ...props.Chart) core.Component {
	prop := props.Chart{}
	if len(ps) > 0 {
		prop = ps[0]
	}

	return &TimeSeries{
		timeSeriesList: timeSeriesList,
		prop:           prop,
	}
}

func NewTimeSeriesCol(size int, timeSeriesList []entity.TimeSeries, ps ...props.Chart) core.Col {
	timeSeries := NewTimeSeries(timeSeriesList, ps...)
	return col.New(size).Add(timeSeries)
}

func NewTimeSeriesRow(height float64, timeSeriesList []entity.TimeSeries, ps ...props.Chart) core.Row {
	timeSeries := NewTimeSeries(timeSeriesList, ps...)
	c := col.New().Add(timeSeries)
	return row.New(height).Add(c)
}

func (b *TimeSeries) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddTimeSeries(b.timeSeriesList, cell, b.prop)
}

func (b *TimeSeries) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type: "time_series",
	}

	return node.New(str)
}

func (b *TimeSeries) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	return 1
}

func (b *TimeSeries) SetConfig(config *entity.Config) {
	b.config = config
}
