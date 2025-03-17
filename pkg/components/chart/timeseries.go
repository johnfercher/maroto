package chart

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type TimeSeries struct {
	timeSeriesList []entity.TimeSeries
	config         *entity.Config
}

func NewTimeSeries(timeSeriesList []entity.TimeSeries) core.Component {
	return &TimeSeries{
		timeSeriesList: timeSeriesList,
	}
}

func NewTimeSeriesCol(size int, timeSeriesList []entity.TimeSeries) core.Col {
	timeSeries := NewTimeSeries(timeSeriesList)
	return col.New(size).Add(timeSeries)
}

func NewTimeSeriesRow(height float64, timeSeriesList []entity.TimeSeries) core.Row {
	timeSeries := NewTimeSeries(timeSeriesList)
	c := col.New().Add(timeSeries)
	return row.New(height).Add(c)
}

func (b *TimeSeries) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddTimeSeries(b.timeSeriesList, cell)
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
