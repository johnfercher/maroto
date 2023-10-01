package image

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type bytesImage struct {
	bytes     []byte
	extension extension.Type
	prop      props.Rect
	config    *entity.Config
}

func NewFromBytes(bytes []byte, extension extension.Type, ps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &bytesImage{
		bytes:     bytes,
		prop:      prop,
		extension: extension,
	}
}

func NewFromBytesCol(size int, bytes []byte, extension extension.Type, ps ...props.Rect) core.Col {
	image := NewFromBytes(bytes, extension, ps...)
	return col.New(size).Add(image)
}

func NewFromBytesRow(height float64, bytes []byte, extension extension.Type, ps ...props.Rect) core.Row {
	image := NewFromBytes(bytes, extension, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

func (b *bytesImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromBytes(b.bytes, cell, &b.prop, b.extension)
}

func (b *bytesImage) GetStructure() *node.Node[core.Structure] {
	trimLength := 10
	if len(b.bytes) < trimLength {
		trimLength = len(b.bytes)
	}

	str := core.Structure{
		Type:  "bytesImage",
		Value: fmt.Sprintf("%v", b.bytes[:trimLength]),
	}

	return node.New(str)
}

func (b *bytesImage) SetConfig(config *entity.Config) {
	b.config = config
}
