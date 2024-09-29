package page

import "github.com/johnfercher/maroto/v2/pkg/processor/components/row"

type Page struct {
	Header []row.Row
	Rows   []row.Row
}

func NewPage(header, rows []row.Row) *Page {
	return &Page{
		Header: header,
		Rows:   rows,
	}
}
