package v2

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type Renderable interface {
	Render(fpdf fpdf.Fpdf, ctx context.Context)
}

type Document interface {
	Generate(file string) error
	Add(rows ...Row)
}

type Row interface {
	Renderable
	GetHeight() float64
	Add(cols ...Col)
}

type Col interface {
	Renderable
	GetSize() int
	Add(components ...Renderable)
}
