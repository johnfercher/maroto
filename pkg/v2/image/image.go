package image

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type image struct {
	path       string
	_type      v2.DocumentType
	components []v2.Component
}

func New(path string) *image {
	return &image{
		_type: v2.Image,
		path:  path,
	}
}

func (i *image) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(i.path)
	for _, component := range i.components {
		component.Render(fpdf, ctx)
	}
	return
}

func (i *image) GetType() string {
	return i._type.String()
}

func (i *image) Add(_ ...v2.Component) {
	return
}
