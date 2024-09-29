package barcode

import "github.com/johnfercher/maroto/v2/pkg/processor/components/props"

type BarCode struct {
	Props props.BarCodeProps
	Code  string
}

func NewBarCode(props props.BarCodeProps, code string) *BarCode {
	return &BarCode{
		Code:  code,
		Props: props,
	}
}
