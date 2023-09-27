package image

import (
	"strings"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type fileImage struct {
	path   string
	prop   props.Rect
	config *config.Config
}

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

func NewFromFileCol(size int, path string, ps ...props.Rect) core.Col {
	image := NewFromFile(path, ps...)
	return col.New(size).Add(image)
}

func NewFromFileRow(height float64, path string, ps ...props.Rect) core.Row {
	image := NewFromFile(path, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

func (f *fileImage) Render(provider core.Provider, cell *core.Cell) {
	extensionStr := strings.Split(f.path, ".")[1]
	provider.AddImage(f.path, cell, &f.prop, extension.Type(extensionStr))
}

func (f *fileImage) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:  "fileimage",
		Value: f.path,
	}

	return node.New(str)
}

func (f *fileImage) SetConfig(config *config.Config) {
	f.config = config
}
