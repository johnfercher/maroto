// Package list implements creation of lists (old tablelist).
package list

import (
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type List struct {
	Header  core.Row
	Content []core.Row
}

// Add is responsible for adding a row to the table content.
func (l *List) Add(rows ...core.Row) *List {
	l.Content = append(l.Content, rows...)
	return l
}

// GetRows is responsible for returning all rows in the table
func (l *List) GetRows() []core.Row {
	return append([]core.Row{l.Header}, l.Content...)
}

// New is responsible to create an instance of a List.
func New(header core.Row) *List {
	if header == nil {
		header = row.New(0)
	}

	return &List{
		Header:  header,
		Content: make([]core.Row, 0),
	}
}
