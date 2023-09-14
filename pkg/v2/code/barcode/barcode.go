package barcode

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type barcode struct {
	code       string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Barcode
}

func New(code string, barcodeProps ...props.Barcode) domain.Component {
	prop := props.Barcode{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &barcode{
		_type: types.Barcode,
		code:  code,
		prop:  prop,
	}
}

func (b *barcode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddBarCode(b.code, cell, b.prop)
}

func (b *barcode) GetType() string {
	return b._type.String()
}

func (b *barcode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(b._type),
		Value: b.code,
	}

	return tree.NewNode(str)
}
