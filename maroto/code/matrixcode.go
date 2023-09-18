// nolint: dupl
package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/maroto/config"
	"github.com/johnfercher/maroto/v2/maroto/domain"
	"github.com/johnfercher/maroto/v2/maroto/grid/col"
	"github.com/johnfercher/maroto/v2/maroto/grid/row"
	"github.com/johnfercher/maroto/v2/maroto/props"
)

type matrixCode struct {
	code   string
	prop   props.Rect
	config *config.Maroto
}

func NewMatrix(code string, barcodeProps ...props.Rect) domain.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &matrixCode{
		code: code,
		prop: prop,
	}
}

func NewMatrixCol(size int, code string, ps ...props.Rect) domain.Col {
	matrixCode := NewMatrix(code, ps...)
	return col.New(size).Add(matrixCode)
}

func NewMatrixRow(height float64, code string, ps ...props.Rect) domain.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New(height).Add(c)
}

func (m *matrixCode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddMatrixCode(m.code, cell, m.prop)
}

func (m *matrixCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "matrixcode",
		Value: m.code,
	}

	return tree.NewNode(str)
}

func (m *matrixCode) SetConfig(config *config.Maroto) {
	m.config = config
}
