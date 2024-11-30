// pagemapper is the package responsible for mapping page settings
package pagemapper

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Page struct {
	SourceKey string
	Rows      []mappers.Componentmapper
	Factory   mappers.AbstractFactoryMaps
}

func NewPage(rows interface{}, sourceKey string, factory mappers.AbstractFactoryMaps) (*Page, error) {
	newPage := &Page{
		SourceKey: sourceKey,
		Factory:   factory,
	}
	if err := newPage.setRows(rows); err != nil {
		return nil, err
	}

	return newPage, nil
}

// setPages is responsible for factory the pages.
// pages can be a list of pages or just one page
func (p *Page) setRows(rowsDoc interface{}) error {
	templateRows, ok := rowsDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("could not parse template, ensure rows can be converted to map[string] interface{}")
	}

	for templateName, template := range templateRows {
		var rows mappers.Componentmapper
		var err error

		if strings.HasPrefix(templateName, "list") {
			rows, err = p.Factory.NewList(template, templateName, p.Factory.NewRow)
		} else {
			rows, err = p.Factory.NewRow(template, templateName)
		}

		if err != nil {
			return err
		}
		p.Rows = append(p.Rows, rows)
	}

	return nil
}

func (p *Page) getPageContent(content map[string]interface{}) (map[string]interface{}, error) {
	pageContent, ok := content[p.SourceKey]
	if !ok {
		return nil, fmt.Errorf("could not parse template,the page needs the source key \"%s\", but it was not found", p.SourceKey)
	}
	if mapPage, ok := pageContent.(map[string]interface{}); ok {
		return mapPage, nil
	}
	return nil, fmt.Errorf(
		"could not parse template, ensure that the contents of the page \"%s\" can be converted to map[string]interface{}",
		p.SourceKey,
	)
}

func (p *Page) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	pageContent, err := p.getPageContent(content)
	if err != nil {
		return nil, err
	}

	rows := make([]processorprovider.ProviderComponent, 0, len(p.Rows))
	for _, row := range p.Rows {
		newRow, err := row.Generate(pageContent, provider)
		if err != nil {
			return nil, err
		}
		rows = append(rows, newRow...)
	}

	page, err := provider.CreatePage(rows...)
	if err != nil {
		return nil, err
	}
	return []processorprovider.ProviderComponent{page}, nil
}
