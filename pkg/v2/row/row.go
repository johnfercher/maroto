package row

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type row struct {
	height     float64
	_type      v2.DocumentType
	components []v2.Component
}

func New(height float64) *row {
	return &row{
		_type:  v2.Row,
		height: height,
	}
}

func (r *row) GetType() string {
	return r._type.String()
}

func (r *row) Add(components ...v2.Component) v2.Component {
	for _, component := range components {
		if r._type.Accept(component.GetType()) {
			r.components = append(r.components, component)
		}
	}
	return r
}

func (r *row) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
		Type:  string(r._type),
		Value: fmt.Sprintf("%2.f", r.height),
	}

	node := tree.NewNode(0, str)

	for _, c := range r.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (r *row) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(r.height)
	ctx = ctx.WithDimension(ctx.Dimensions.Width, r.height)
	if ctx.GetYOffset() == 0 && ctx.GetCurrentPage() >= fpdf.PageCount() {
		fpdf.AddPage()
		ctx = ctx.NewPage(fpdf.PageNo())
	}
	for _, component := range r.components {
		component.Render(fpdf, ctx)
	}

	r.render(fpdf, ctx)
	return
}

func (r *row) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.SetDrawColor(0, 0, 0)
	//x, y := ctx.GetXOffset(), ctx.GetYOffset()
	fpdf.Ln(ctx.Dimensions.Height)
}
