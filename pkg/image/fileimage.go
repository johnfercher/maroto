package image

import (
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/domain"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"strings"

	"github.com/johnfercher/go-tree/tree"
)

type fileImage struct {
	path   string
	prop   props.Rect
	config *config.Maroto
}

func NewFromFile(path string, ps ...props.Rect) domain.Component {
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

func NewFromFileCol(size int, path string, ps ...props.Rect) domain.Col {
	image := NewFromFile(path, ps...)
	return col.New(size).Add(image)
}

func NewFromFileRow(height float64, path string, ps ...props.Rect) domain.Row {
	image := NewFromFile(path, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

func (f *fileImage) Render(provider domain.Provider, cell internal.Cell) {
	extensionStr := strings.Split(f.path, ".")[1]
	provider.AddImage(f.path, cell, f.prop, extension.Type(extensionStr))
}

func (f *fileImage) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "fileimage",
		Value: f.path,
	}

	return tree.NewNode(str)
}

func (f *fileImage) SetConfig(config *config.Maroto) {
	f.config = config
}
