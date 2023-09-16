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
)

type base64Image struct {
	base64    string
	extension consts.Extension
	prop      props.Rect
}

func NewFromBase64(path string, extension consts.Extension, ps ...props.Rect) domain.Component {
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

func NewFromBase64Col(size int, path string, extension consts.Extension, ps ...props.Rect) domain.Col {
	image := NewFromBase64(path, extension, ps...)
	return col.New(size).Add(image)
}

func NewFromBase64eRow(height float64, path string, extension consts.Extension, ps ...props.Rect) domain.Row {
	c := NewFromBase64Col(12, path, extension, ps...)
	return row.New(height).Add(c)
}

func (b *base64Image) Render(provider domain.Provider, cell internal.Cell, config *config.Maroto) {
	provider.AddImage(b.base64, cell, b.prop, b.extension)
}

func (b *base64Image) GetStructure() *tree.Node[domain.Structure] {
	trimLength := 10
	if len(b.base64) < trimLength {
		trimLength = len(b.base64)
	}

	str := domain.Structure{
		Type:  "base64image",
		Value: b.base64[:trimLength],
	}

	return tree.NewNode(str)
}

func (b *base64Image) GetValue() string {
	return b.base64
}
