package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type matrixCode struct {
	code string
	prop props.Rect
}

func NewMatrixCode(code string, barcodeProps ...props.Rect) domain.Node {
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
