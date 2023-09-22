package list

import (
	"errors"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type Listable interface {
	GetHeader() core.Row
	GetContent(i int) core.Row
}

func BuildFromPointer[T Listable](arr []*T) ([]core.Row, error) {
	if len(arr) == 0 {
		return nil, errors.New("empty array")
	}

	var list []T
	for _, pointer := range arr {
		if pointer == nil {
			return nil, errors.New("nil element in array")
		}
		list = append(list, *pointer)
	}

	return Build(list)
}

func Build[T Listable](arr []T) ([]core.Row, error) {
	if len(arr) == 0 {
		return nil, errors.New("empty array")
	}

	var rows []core.Row

	rows = append(rows, arr[0].GetHeader())

	for i, element := range arr {
		rows = append(rows, element.GetContent(i))
	}

	return rows, nil
}
