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

type matrixCode struct {
	code string
	prop props.Rect
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
	c := NewMatrixCol(12, code, ps...)
	return row.New(height).Add(c)
}

func (m *matrixCode) Render(provider domain.Provider, cell internal.Cell, config *config.Maroto) {
	provider.AddMatrixCode(m.code, cell, m.prop)
}

func (m *matrixCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "matrixcode",
		Value: m.code,
	}

	return tree.NewNode(str)
}

func (m *matrixCode) GetValue() string {
	return m.code
}
