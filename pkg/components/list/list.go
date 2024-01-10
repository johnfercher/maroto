// Package implements creation of lists (old tablelist).
package list

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

// Listable is the main abstraction of a listable item in a TableList.
// A collection of objects that implements this interface may be added
// in a list.
type Listable interface {
	GetHeader() core.Row
	GetContent(i int) core.Row
}

// BuildFromPointer is responsible to receive a collection of objects that implements
// Listable and build the rows of TableList. This method should be used in case of a collection
// of pointers.
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

// Build is responsible to receive a collection of objects that implements
// Listable and build the rows of TableList. This method should be used in case of a collection
// of values.
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
