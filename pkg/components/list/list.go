package list

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type Listable interface {
	GetHeader() core.Row
	GetContent(i int) core.Row
}

func Build[T Listable](list ...T) []core.Row {
	if len(list) == 0 {
		return nil
	}

	var rows []core.Row

	rows = append(rows, list[0].GetHeader())

	for i, element := range list {
		rows = append(rows, element.GetContent(i))
	}

	return rows
}
