package row

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/col"
	"github.com/johnfercher/maroto/v2/pkg/processor/provider"
)

type Row struct {
	Cols []col.Col
}

func NewRow(cols ...col.Col) *Row {
	return &Row{
		Cols: cols,
	}
}

func (r *Row) Generate(provider provider.Provider) core.Row {
	cols := make([]core.Col, len(r.Cols))

	for i, col := range r.Cols {
		cols[i] = col.Generate(provider)
	}
	return provider.CreateRow(cols...)
}
