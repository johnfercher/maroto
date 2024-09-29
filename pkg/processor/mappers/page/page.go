package page

import "github.com/johnfercher/maroto/v2/pkg/processor/mappers/row"

type Page struct {
	Header []row.Row `json:"header"`
	Rows   []row.Row `json:"rows"`
}
