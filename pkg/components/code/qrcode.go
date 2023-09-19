// nolint: dupl
package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type qrCode struct {
	code   string
	prop   props.Rect
	config *config.Config
}

func NewQr(code string, barcodeProps ...props.Rect) core.Component {
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

func NewQrCol(size int, code string, ps ...props.Rect) core.Col {
	qrCode := NewQr(code, ps...)
	return col.New(size).Add(qrCode)
}

func NewQrRow(height float64, code string, ps ...props.Rect) core.Row {
	qrCode := NewQr(code, ps...)
	c := col.New().Add(qrCode)
	return row.New(height).Add(c)
}

func (q *qrCode) Render(provider core.Provider, cell core.Cell) {
	provider.AddQrCode(q.code, cell, q.prop)
}

func (q *qrCode) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "qrcode",
		Value: q.code,
	}

	return tree.NewNode(str)
}

func (q *qrCode) SetConfig(config *config.Config) {
	q.config = config
}
