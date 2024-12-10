// Package list implements creation of lists (old tablelist).
package list

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type List struct {
	Header  core.Row
	Content []core.Row
	props   props.List
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

// BuildListWithFixedHeader is responsible for adding list to table with fixed header
func (l *List) BuildListWithFixedHeader(m core.Maroto) error {
	if m == nil {
		return errors.New("maroto instance cannot be null")
	}
	totalRowsAdded := 0
	lastAddedStatus := true

	for {
		contentRowsAdded := l.addListToCurrentPage(&m, l.Header, l.Content[totalRowsAdded:]...)
		if contentRowsAdded == 0 && !lastAddedStatus {
			return errors.New("the rows cannot be adjusted to the document, check if MinimumRowsBypage is less than the rows capacity of the page")
		}

		totalRowsAdded += contentRowsAdded
		lastAddedStatus = contentRowsAdded != 0

		if totalRowsAdded >= len(l.Content) {
			break
		}
		m.FillPageToAddNew()
	}
	return nil
}

// addListToCurrentPage is responsible for adding the list to the current page and returning
// the number of lines of content inserted.
func (l *List) addListToCurrentPage(m *core.Maroto, header core.Row, list ...core.Row) int {
	list = append([]core.Row{header}, list...)
	amountRows := (*m).FitsOnCurrentPage(list...)
	if amountRows >= l.props.MinimumRowsBypage {
		(*m).AddRows(list[:amountRows]...)
		return amountRows - 1
	}
	return 0
}

// New is responsible to create an instance of a List.
func New(header core.Row, ps ...props.List) *List {
	listProps := props.List{}
	if len(ps) > 0 {
		listProps = ps[0]
	}
	if header == nil {
		header = row.New(0)
	}
	listProps.MakeValid()

	return &List{
		Header:  header,
		Content: make([]core.Row, 0),
		props:   listProps,
	}
}
