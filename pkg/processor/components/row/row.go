package row

import "github.com/johnfercher/maroto/v2/pkg/processor/components/col"

type Row struct {
	Cols []col.Col
}

func NewRow(cols ...col.Col) *Row {
	return &Row{
		Cols: cols,
	}
}
