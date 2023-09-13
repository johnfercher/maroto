package image

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/text"
)

type base64Image struct {
	path       string
	extension  consts.Extension
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Rect
}

func NewFromBase64(path string, extension consts.Extension, imageProps ...props.Rect) *base64Image {
	prop := props.Rect{}
	if len(imageProps) > 0 {
		prop = imageProps[0]
	}
	prop.MakeValid()

	return &base64Image{
		_type:     v2.Image,
		path:      path,
		prop:      prop,
		extension: extension,
	}
}

func (b *base64Image) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)
	img := internal.NewImage(fpdf, math)
	err := img.AddFromBase64(
		b.path,
		internal.Cell{fpdf.GetX() - ctx.Margins.Left,
			fpdf.GetY() - ctx.Margins.Top,
			ctx.Dimensions.Width,
			ctx.Dimensions.Height},
		b.prop,
		b.extension,
	)
	if err != nil {
		fpdf.ClearError()
		txt := text.New("Failed to render fileImage")
		txt.Render(fpdf, ctx)
	}
}

func (b *base64Image) GetType() string {
	return b._type.String()
}

func (b *base64Image) Add(_ ...v2.Component) v2.Component {
	return b
}
