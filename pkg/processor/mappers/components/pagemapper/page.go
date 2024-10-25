// pagemapper is the package responsible for mapping page settings
package pagemapper

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type Page struct {
	SourceKey string
	Rows      []mappers.Componentmapper
	factory   mappers.AbstractFactoryMaps
}

func NewPage(rows interface{}, sourceKey string, factory mappers.AbstractFactoryMaps) (*Page, error) {
	newPage := &Page{
		SourceKey: sourceKey,
		factory:   factory,
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
		return fmt.Errorf("ensure rows can be converted to map[string] interface{}")
	}

	for templateName, template := range templateRows {
		var rows mappers.Componentmapper
		var err error

		if strings.HasPrefix(templateName, "list") {
			rows, err = p.factory.NewList(template, templateName, p.factory.NewRow)
		} else {
			rows, err = p.factory.NewRow(template, templateName)
		}

		if err != nil {
			return err
		}
		p.Rows = append(p.Rows, rows)
	}

	return nil
}

func (p *Page) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
