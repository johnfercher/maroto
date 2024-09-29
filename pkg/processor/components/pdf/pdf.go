package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/page"
)

type Pdf struct {
	Builder *builder.Builder
	Pages   []*page.Page
}

func NewPdf(builder *builder.Builder, pages ...*page.Page) *Pdf {
	return &Pdf{
		Builder: builder,
		Pages:   pages,
	}
}
