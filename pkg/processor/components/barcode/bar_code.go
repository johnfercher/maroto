package barcode

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/provider"
)

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

func (b *BarCode) Generate(provider provider.Provider) core.Component {
	return provider.CreateBarCode(b.Code, b.Props)
}
