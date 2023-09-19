package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/context"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type barcode struct {
	code   string
	prop   props.Barcode
	config *config.Maroto
}

func NewBar(code string, ps ...props.Barcode) core.Component {
	prop := props.Barcode{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &barcode{
		code: code,
		prop: prop,
	}
}

func NewBarCol(size int, code string, ps ...props.Barcode) core.Col {
	bar := NewBar(code, ps...)
	return col.New(size).Add(bar)
}

func NewBarRow(height float64, code string, ps ...props.Barcode) core.Row {
	bar := NewBar(code, ps...)
	c := col.New().Add(bar)
	return row.New(height).Add(c)
}

func (b *barcode) Render(provider core.Provider, cell context.Cell) {
	provider.AddBarCode(b.code, cell, b.prop)
}

func (b *barcode) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "barcode",
		Value: b.code,
	}

	return tree.NewNode(str)
}

func (b *barcode) GetValue() string {
	return b.code
}

func (b *barcode) SetConfig(config *config.Maroto) {
	b.config = config
}
