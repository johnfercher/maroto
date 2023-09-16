package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type qrCode struct {
	code string
	prop props.Rect
}

func NewQr(code string, barcodeProps ...props.Rect) domain.Node {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &qrCode{
		code: code,
		prop: prop,
	}
}

func (q *qrCode) Render(provider domain.Provider, cell internal.Cell, config *config.Maroto) {
	provider.AddQrCode(q.code, cell, q.prop)
}

func (q *qrCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "qrcode",
		Value: q.code,
	}

	return tree.NewNode(str)
}
