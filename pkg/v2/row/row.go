package row

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type row struct {
	height float64
	_type  types.DocumentType
	cols   []v2.Col
}

func (r *row) Add(cols ...v2.Col) {
	r.cols = append(r.cols, cols...)
}

func New(height float64) v2.Row {
	return &row{
		_type:  types.Row,
		height: height,
	}
}

func (r *row) GetType() string {
	return r._type.String()
}

func (r *row) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
		Type:  string(r._type),
		Value: fmt.Sprintf("%2.f", r.height),
	}

	node := tree.NewNode(0, str)

	for _, c := range r.cols {
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
	for _, col := range r.cols {
		col.Render(fpdf, ctx)
	}

	r.render(fpdf, ctx)
	return
}

func (r *row) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.SetDrawColor(0, 0, 0)
	//x, y := ctx.GetXOffset(), ctx.GetYOffset()
	fpdf.Ln(ctx.Dimensions.Height)
}
