// nolint: dupl
package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
)

type qrCode struct {
	code   string
	prop   props.Rect
	config *config.Maroto
}

func NewQr(code string, barcodeProps ...props.Rect) domain.Component {
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

func NewQrCol(size int, code string, ps ...props.Rect) domain.Col {
	qrCode := NewQr(code, ps...)
	return col.New(size).Add(qrCode)
}

func NewQrRow(height float64, code string, ps ...props.Rect) domain.Row {
	qrCode := NewQr(code, ps...)
	c := col.New().Add(qrCode)
	return row.New(height).Add(c)
}

func (q *qrCode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddQrCode(q.code, cell, q.prop)
}

func (q *qrCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "qrcode",
		Value: q.code,
	}

	return tree.NewNode(str)
}

func (q *qrCode) SetConfig(config *config.Maroto) {
	q.config = config
}
