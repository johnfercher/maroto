package row

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
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

func (r *row) Add(components ...v2.Component) {
	for _, component := range components {
		if r._type.Accept(component.GetType()) {
			r.components = append(r.components, component)
		}
	}
}

func (r *row) Render(fpdf fpdf.Fpdf, ctx *v2.Context) {
	ctx.Print(r.height)
	ctx = ctx.WithDimension(ctx.MaxWidth(), r.height)
	r.render(fpdf, ctx)
	for _, component := range r.components {
		component.Render(fpdf, ctx)
	}
}

func (r *row) render(fpdf fpdf.Fpdf, ctx *v2.Context) {
	fpdf.SetFont("Arial", "B", 16)
	x, y := ctx.GetX(), ctx.GetY()
	fpdf.CellFormat(x, y, "", "1", 0, "C", false, 0, "")
}
