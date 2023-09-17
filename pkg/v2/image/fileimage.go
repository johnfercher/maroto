package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"strings"
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
	c := col.New(0).Add(image)
	return row.New(height).Add(c)
}

func (f *fileImage) Render(provider domain.Provider, cell internal.Cell) {
	extension := strings.Split(f.path, ".")[1]
	provider.AddImage(f.path, cell, f.prop, consts.Extension(extension))
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
