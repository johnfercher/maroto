package matrixcode

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type matrixCode struct {
	code       string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Rect
}

func New(code string, barcodeProps ...props.Rect) domain.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &matrixCode{
		_type: types.MatrixCode,
		code:  code,
		prop:  prop,
	}
}

func (m *matrixCode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddMatrixCode(m.code, cell, m.prop)
}

func (m *matrixCode) GetType() string {
	return m._type.String()
}

func (m *matrixCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(m._type),
		Value: m.code,
	}

	return tree.NewNode(str)
}
