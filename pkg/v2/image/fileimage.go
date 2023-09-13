package image

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/text"
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

func (f *fileImage) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)
	img := internal.NewImage(fpdf, math)
	x := fpdf.GetX() - ctx.Margins.Left - ctx.Dimensions.Width
	y := fpdf.GetY() - ctx.Margins.Top
	err := img.AddFromFile(
		f.path,
		internal.Cell{x, y, ctx.Dimensions.Width, ctx.Dimensions.Height},
		f.prop)
	if err != nil {
		fpdf.ClearError()
		txt := text.New("Failed to render fileImage")
		txt.Render(fpdf, ctx)
	}
}

func (f *fileImage) GetType() string {
	return f._type.String()
}

func (f *fileImage) Add(_ ...domain.Node) domain.Node {
	return f
}

func (f *fileImage) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(f._type),
		Value: f.path,
	}

	return tree.NewNode(0, str)
}
