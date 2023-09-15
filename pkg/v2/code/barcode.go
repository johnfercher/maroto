package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type barcode struct {
	code string
	prop props.Barcode
}

func NewBar(code string, barcodeProps ...props.Barcode) domain.Node {
	prop := props.Barcode{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &barcode{
		code: code,
		prop: prop,
	}
}

func (b *barcode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddBarCode(b.code, cell, b.prop)
}

func (b *barcode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "barcode",
		Value: b.code,
	}

	return tree.NewNode(str)
}
