package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type fileImage struct {
	path       string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Rect
}

func NewFromFile(path string, imageProps ...props.Rect) domain.Component {
	prop := props.Rect{}
	if len(imageProps) > 0 {
		prop = imageProps[0]
	}
	prop.MakeValid()

	return &fileImage{
		_type: types.Image,
		path:  path,
		prop:  prop,
	}
}

func (f *fileImage) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddImageFromFile(f.path, cell, f.prop)
}

func (f *fileImage) GetType() string {
	return f._type.String()
}

func (f *fileImage) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(f._type),
		Value: f.path,
	}

	return tree.NewNode(str)
}
