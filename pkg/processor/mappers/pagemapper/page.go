// pagemapper is the package responsible for mapping page settings
package pagemapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components/page"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/row"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/rowmapper"
)

type Page struct {
	Header []rowmapper.Row `json:"header"`
	Rows   []rowmapper.Row `json:"rows"`
}

// generate is responsible for the builder page according to the submitted content
func (p *Page) Generate(content map[string]interface{}) (*page.Page, error) {
	header, err := p.generateRows(content, p.Header)
	if err != nil {
		return nil, err
	}

	rows, err := p.generateRows(content, p.Header)
	if err != nil {
		return nil, err
	}

	return page.NewPage(header, rows), nil
}

func (p *Page) generateRows(content map[string]interface{}, templates []rowmapper.Row) ([]row.Row, error) {
	rows := make([]row.Row, len(content))

	for _, headerRow := range templates {
		generatedRow, err := headerRow.Generate(content)
		if err != nil {
			return nil, err
		}
		rows = append(rows, *generatedRow)
	}

	return rows, nil
}
