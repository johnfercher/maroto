package page

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/row"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Page struct {
	Rows []row.Row
}

func NewPage(header, rows []row.Row) *Page {
	return &Page{
		Rows: rows,
	}
}

func (p *Page) Generate(provider processorprovider.ProcessorProvider) {
	rows := make([]core.Row, len(p.Rows))

	for i, row := range p.Rows {
		rows[i] = row.Generate(provider)
	}
	provider.CreatePage(rows...)
}
