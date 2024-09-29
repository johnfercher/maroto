package row

import "github.com/johnfercher/maroto/v2/pkg/processor/mappers/col"

type Row struct {
	List string    `json:"list"`
	Cols []col.Col `json:"cols"`
}
