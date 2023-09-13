package image

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/text"
)

type image struct {
	path       string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Rect
}

func NewFromFile(path string, imageProps ...props.Rect) *image {
	prop := props.Rect{}
	if len(imageProps) > 0 {
		prop = imageProps[0]
	}
	prop.MakeValid()

	return &image{
		_type: v2.Image,
		path:  path,
		prop:  prop,
	}
}

func (i *image) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)
	img := internal.NewImage(fpdf, math)
	err := img.AddFromFile(
		i.path,
		internal.Cell{fpdf.GetX() - ctx.Margins.Left,
			fpdf.GetY() - ctx.Margins.Top,
			ctx.Dimensions.Width,
			ctx.Dimensions.Height},
		i.prop)
	if err != nil {
		fpdf.ClearError()
		txt := text.New("Failed to render image")
		txt.Render(fpdf, ctx)
	}
}

func (i *image) GetType() string {
	return i._type.String()
}

func (i *image) Add(_ ...v2.Component) {
	return
}
