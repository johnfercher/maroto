package page

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type page struct {
	_type      v2.DocumentType
	components []v2.Component
}

func New() *page {
	return &page{
		_type: v2.Page,
	}
}

func (p *page) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.AddPage()
	ctx = ctx.NewPage(fpdf.PageNo())
	for _, component := range p.components {
		component.Render(fpdf, ctx)
	}
}

func (p *page) GetType() string {
	return p._type.String()
}

func (p *page) Add(components ...v2.Component) v2.Component {
	for _, component := range components {
		if p._type.Accept(component.GetType()) {
			p.components = append(p.components, component)
		}
	}
	return p
}
