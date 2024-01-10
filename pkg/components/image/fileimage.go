// Package implements creation of images from file and bytes.
package image

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type fileImage struct {
	path   string
	prop   props.Rect
	config *entity.Config
}

// NewFromFile is responsible to create an instance of an Image.
func NewFromFile(path string, ps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &fileImage{
		path: path,
		prop: prop,
	}
}

// NewFromFileCol is responsible to create an instance of an Image wrapped in a Col.
func NewFromFileCol(size int, path string, ps ...props.Rect) core.Col {
	image := NewFromFile(path, ps...)
	return col.New(size).Add(image)
}

// NewFromFileRow is responsible to create an instance of an Image wrapped in a Row.
func NewFromFileRow(height float64, path string, ps ...props.Rect) core.Row {
	image := NewFromFile(path, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

func (f *fileImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromFile(f.path, cell, &f.prop)
}

func (f *fileImage) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "fileImage",
		Value:   f.path,
		Details: f.prop.ToMap(),
	}

	return node.New(str)
}

func (f *fileImage) SetConfig(config *entity.Config) {
	f.config = config
}
