package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/page"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
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

func (p *Pdf) Generate(provider processorprovider.ProcessorProvider) processorprovider.ProcessorProvider {
	for _, page := range p.Pages {
		page.Generate(provider)
	}

	return provider
}
