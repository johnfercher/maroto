package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type fileImage struct {
	path string
	prop props.Rect
}

func NewFromFile(path string, imageProps ...props.Rect) domain.Node {
	prop := props.Rect{}
	if len(imageProps) > 0 {
		prop = imageProps[0]
	}
	prop.MakeValid()

	return &fileImage{
		path: path,
		prop: prop,
	}
}

func (f *fileImage) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddImageFromFile(f.path, cell, f.prop)
}

func (f *fileImage) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "fileimage",
		Value: f.path,
	}

	return tree.NewNode(str)
}
