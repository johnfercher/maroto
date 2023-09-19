package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/context"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type base64Image struct {
	base64    string
	extension extension.Type
	prop      props.Rect
	config    *config.Config
}

func NewFromBase64(path string, extension extension.Type, ps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &base64Image{
		base64:    path,
		prop:      prop,
		extension: extension,
	}
}

func NewFromBase64Col(size int, path string, extension extension.Type, ps ...props.Rect) core.Col {
	image := NewFromBase64(path, extension, ps...)
	return col.New(size).Add(image)
}

func NewFromBase64Row(height float64, path string, extension extension.Type, ps ...props.Rect) core.Row {
	image := NewFromBase64(path, extension, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

func (b *base64Image) Render(provider core.Provider, cell context.Cell) {
	provider.AddImage(b.base64, cell, b.prop, b.extension)
}

func (b *base64Image) GetStructure() *tree.Node[core.Structure] {
	trimLength := 10
	if len(b.base64) < trimLength {
		trimLength = len(b.base64)
	}

	str := core.Structure{
		Type:  "base64image",
		Value: b.base64[:trimLength],
	}

	return tree.NewNode(str)
}

func (b *base64Image) SetConfig(config *config.Config) {
	b.config = config
}
