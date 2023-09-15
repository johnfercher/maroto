package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type base64Image struct {
	base64    string
	extension consts.Extension
	prop      props.Rect
}

func NewFromBase64(path string, extension consts.Extension, imageProps ...props.Rect) domain.Node {
	prop := props.Rect{}
	if len(imageProps) > 0 {
		prop = imageProps[0]
	}
	prop.MakeValid()

	return &base64Image{
		base64:    path,
		prop:      prop,
		extension: extension,
	}
}

func (b *base64Image) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddImageFromBase64(b.base64, cell, b.prop, b.extension)
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
