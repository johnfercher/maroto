package barcode

import "github.com/johnfercher/maroto/v2/pkg/processor/mappers/props/barcode"

type BarCode struct {
	Props      barcode.BarCodeProps `json:"props"`
	Source_key string               `json:"source_key"`
}
