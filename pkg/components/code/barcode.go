// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl // It's similar to Barcode.go and it's hard to extract common code.
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Barcode struct {
	code   string
	prop   props.Barcode
	config *entity.Config
}

// NewBar is responsible to create an instance of a Barcode.
func NewBar(code string, ps ...props.Barcode) core.Component {
	prop := props.Barcode{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &Barcode{
		code: code,
		prop: prop,
	}
}

// NewBarCol is responsible to create an instance of a Barcode wrapped in a Col.
func NewBarCol(size int, code string, ps ...props.Barcode) core.Col {
	bar := NewBar(code, ps...)
	return col.New(size).Add(bar)
}

// NewBarRow is responsible to create an instance of a Barcode wrapped in a Row.
func NewBarRow(height float64, code string, ps ...props.Barcode) core.Row {
	bar := NewBar(code, ps...)
	c := col.New().Add(bar)
	return row.New(height).Add(c)
}

// Render renders a Barcode into a PDF context.
func (b *Barcode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddBarCode(b.code, cell, &b.prop)
}

// GetStructure returns the Structure of a Barcode.
func (b *Barcode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "barcode",
		Value:   b.code,
		Details: b.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the configuration of a Barcode.
func (b *Barcode) SetConfig(config *entity.Config) {
	b.config = config
}
