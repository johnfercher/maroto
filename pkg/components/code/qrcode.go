// Package implements creation of barcode, matrixcode and qrcode.
// nolint: dupl
package code

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type qrCode struct {
	code   string
	prop   props.Rect
	config *entity.Config
}

// NewQr is responsible to create an instance of a QrCode.
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

// NewQrCol is responsible to create an instance of a QrCode wrapped in a Col.
func NewQrCol(size int, code string, ps ...props.Rect) core.Col {
	qrCode := NewQr(code, ps...)
	return col.New(size).Add(qrCode)
}

// NewQrRow is responsible to create an instance of a QrCode wrapped in a Row.
func NewQrRow(height float64, code string, ps ...props.Rect) core.Row {
	qrCode := NewQr(code, ps...)
	c := col.New().Add(qrCode)
	return row.New(height).Add(c)
}

func (q *qrCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddQrCode(q.code, cell, &q.prop)
}

func (q *qrCode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "qrcode",
		Value:   q.code,
		Details: q.prop.ToMap(),
	}

	return node.New(str)
}

func (q *qrCode) SetConfig(config *entity.Config) {
	q.config = config
}
