// documentmapper is the package responsible for mapping pdf settings
package documentmapper

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/processor/components/pdf"
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
)

type Document struct {
	factory mappers.AbstractFactoryMaps
	Builder buildermapper.Builder
	Header  []mappers.Componentmapper
	Footer  []mappers.Componentmapper
	pages   []mappers.Componentmapper
}

// NewPdf is responsible for creating the pdf template
func NewPdf(document string, deserializer core.Deserializer, factory mappers.AbstractFactoryMaps) (*Document, error) {
	newPdf := Document{factory: factory}
	template, err := deserializer.Deserialize(document)
	if err != nil {
		return nil, err
	}

	err = newPdf.addComponentsToPdf(template)
	if err != nil {
		return nil, err
	}

	return &newPdf, nil
}

// addComponentsToPdf is responsible for adding all the components that are part of the template to the PDF.
// This is the method where the components that are part of the PDF are created.
func (p *Document) addComponentsToPdf(templates map[string]interface{}) error {
	fieldMappers := p.getFieldMappers()

	for field, template := range templates {
		mapper, ok := fieldMappers[field]
		if !ok {
			return fmt.Errorf("the field %s present in the template cannot be mapped to any valid component", field)
		}
		err := mapper(template)
		if err != nil {
			return err
		}
	}
	return nil
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (p *Document) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"builder": p.setBuilder,
		"header":  p.setHeader,
		"footer":  p.setFooter,
		"pages":   p.setPages,
	}
}

// setBuilder is responsible for factories builder information
func (p *Document) setBuilder(builderDoc interface{}) error {
	builder, err := buildermapper.NewBuilder(builderDoc)
	if err != nil {
		return err
	}

	p.Builder = *builder
	return nil
}

// setHeader is responsible for factory the header
func (p *Document) setHeader(rowsDoc interface{}) error {
	rowsTemplate, ok := rowsDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("header cannot be deserialized, ensure header has an array of rows")
	}

	for templateKey, rowTemplate := range rowsTemplate {
		row, err := p.factory.NewRow(rowTemplate, templateKey)
		if err != nil {
			return err
		}
		p.Header = append(p.Header, row)
	}

	return nil
}

// setFooter is responsible for factory the footer
func (p *Document) setFooter(rowsDoc interface{}) error {
	rowsTemplate, ok := rowsDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("footer cannot be deserialized, ensure footer has an array of rows")
	}

	for templateKey, rowTemplate := range rowsTemplate {
		row, err := p.factory.NewRow(rowTemplate, templateKey)
		if err != nil {
			return err
		}
		p.Footer = append(p.Footer, row)
	}

	return nil
}

// setPages is responsible for factory the pages.
// pages can be a list of pages or just one page
func (p *Document) setPages(pagesDoc interface{}) error {
	templatePage, ok := pagesDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("ensure pages can be converted to map[string] interface{}")
	}

	for templateName, template := range templatePage {
		var page mappers.Componentmapper
		var err error

		if strings.HasPrefix(templateName, "list") {
			page, err = p.factory.NewList(template, templateName, p.factory.NewPage)
		} else {
			page, err = p.factory.NewPage(template, templateName)
		}

		if err != nil {
			return err
		}
		p.pages = append(p.pages, page)
	}

	return nil
}

// generate is responsible for the builder pdf according to the submitted content
func (p *Document) Generate(content map[string]interface{}) (*pdf.Pdf, error) {
	return nil, nil
}
