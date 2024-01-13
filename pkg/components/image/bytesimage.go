// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type BytesImage struct {
	bytes     []byte
	extension extension.Type
	prop      props.Rect
	config    *entity.Config
}

// NewFromBytes is responsible to create an instance of an Image.
func NewFromBytes(bytes []byte, extension extension.Type, ps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &BytesImage{
		bytes:     bytes,
		prop:      prop,
		extension: extension,
	}
}

// NewFromBytesCol is responsible to create an instance of an Image wrapped in a Col.
func NewFromBytesCol(size int, bytes []byte, extension extension.Type, ps ...props.Rect) core.Col {
	image := NewFromBytes(bytes, extension, ps...)
	return col.New(size).Add(image)
}

// NewFromBytesRow is responsible to create an instance of an Image wrapped in a Row.
func NewFromBytesRow(height float64, bytes []byte, extension extension.Type, ps ...props.Rect) core.Row {
	image := NewFromBytes(bytes, extension, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

// Render renders an Image into a PDF context.
func (b *BytesImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromBytes(b.bytes, cell, &b.prop, b.extension)
}

// GetStructure returns the Structure of an Image.
func (b *BytesImage) GetStructure() *node.Node[core.Structure] {
	trimLength := 10
	if len(b.bytes) < trimLength {
		trimLength = len(b.bytes)
	}

	str := core.Structure{
		Type:    "bytesImage",
		Value:   b.bytes[:trimLength],
		Details: b.prop.ToMap(),
	}

	str.Details["extension"] = b.extension
	str.Details["bytes_size"] = len(b.bytes)

	return node.New(str)
}

// SetConfig sets the pdf config.
func (b *BytesImage) SetConfig(config *entity.Config) {
	b.config = config
}
