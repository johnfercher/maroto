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
	Rows      []mappers.OrderedComponents
	Factory   mappers.AbstractFactoryMaps
	order     int
}

// NewPage is responsible for create an template page
func NewPage(templatePage interface{}, sourceKey string, factory mappers.AbstractFactoryMaps) (*Page, error) {
	newPage := &Page{
		SourceKey: sourceKey,
		Factory:   factory,
	}
	if err := newPage.setRows(templatePage); err != nil {
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
	err := p.setPageOrder(&templateRows)
	if err != nil {
		return err
	}

	p.Rows = make([]mappers.OrderedComponents, len(templateRows))
	for templateName, template := range templateRows {
		var row mappers.OrderedComponents

		if strings.HasPrefix(templateName, "list") {
			row, err = p.Factory.NewList(template, templateName, p.Factory.NewRow)
		} else {
			row, err = p.Factory.NewRow(template, templateName)
		}

		if err != nil {
			return err
		}

		if err := p.addComponent(row); err != nil {
			return err
		}
	}
	return nil
}

// addPage is responsible for validating and adding the page to the template
func (p *Page) addComponent(row mappers.OrderedComponents) error {
	order := row.GetOrder()
	if order > len(p.Rows) {
		return fmt.Errorf("component order cannot be greater than %d, this is the number of components in the template", len(p.Rows))
	}
	if p.Rows[order-1] != nil {
		return fmt.Errorf("cannot create page template, component order cannot be repeated")
	}

	p.Rows[order-1] = row
	return nil
}

// GetOrder is responsible for returning the component's defined order
func (p *Page) GetOrder() int {
	return p.order
}

// setPageOrder is responsible for validating the component order and adding the order to the page
func (p *Page) setPageOrder(template *map[string]interface{}) error {
	order, ok := (*template)["order"]
	if !ok {
		return fmt.Errorf("could not find field order on page \"%s\"", p.SourceKey)
	}
	validOrder, ok := order.(float64)
	if !ok || validOrder < 1 {
		return fmt.Errorf("the order field passed on page \"%s\" is not valid", p.SourceKey)
	}

	p.order = int(validOrder)
	delete(*template, "order")
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

// Generate is responsible for computing the page component with shipping data
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
