// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type MatrixCode struct {
	code   string
	prop   props.Rect
	config *entity.Config
}

// NewMatrix is responsible to create an instance of a MatrixCode.
func NewMatrix(code string, barcodeProps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &MatrixCode{
		code: code,
		prop: prop,
	}
}

// NewMatrixCol is responsible to create an instance of a MatrixCode wrapped in a Col.
func NewMatrixCol(size int, code string, ps ...props.Rect) core.Col {
	matrixCode := NewMatrix(code, ps...)
	return col.New(size).Add(matrixCode)
}

// NewAutoMatrixRow is responsible to create an instance of a Matrix code wrapped in a Row with automatic height.
//   - code: The value that must be placed in the matrixcode
//   - ps: A set of settings that must be applied to the matrixcode
func NewAutoMatrixRow(code string, ps ...props.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New().Add(c)
}

// NewMatrixRow is responsible to create an instance of a MatrixCode wrapped in a Row.
func NewMatrixRow(height float64, code string, ps ...props.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New(height).Add(c)
}

// Render renders a MatrixCode into a PDF context.
func (m *MatrixCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddMatrixCode(m.code, cell, &m.prop)
}

// GetStructure returns the Structure of a MatrixCode.
func (m *MatrixCode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "matrixcode",
		Value:   m.code,
		Details: m.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the code will have in the PDF
func (b *MatrixCode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	dimensions, err := provider.GetDimensionsByMatrixCode(b.code)
	if err != nil {
		return 0
	}
	proportion := dimensions.Height / dimensions.Width
	width := (b.prop.Percent / 100) * cell.Width
	return proportion * width
}

// SetConfig sets the configuration of a MatrixCode.
func (m *MatrixCode) SetConfig(config *entity.Config) {
	m.config = config
}
