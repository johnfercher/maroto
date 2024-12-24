// documentmapper is the package responsible for mapping pdf settings
package documentmapper

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Document struct {
	factory mappers.AbstractFactoryMaps
	Builder buildermapper.Builder
	Header  []mappers.Componentmapper
	Footer  []mappers.Componentmapper
	Pages   []mappers.Componentmapper
}

// NewPdf is responsible for creating the pdf template
// parse the model and create the pdf object
func NewPdf(template map[string]any, factory mappers.AbstractFactoryMaps) (*Document, error) {
	newPdf := Document{
		factory: factory, Builder: buildermapper.Builder{}, Pages: make([]mappers.Componentmapper, 0),
		Header: make([]mappers.Componentmapper, 0), Footer: make([]mappers.Componentmapper, 0),
	}

	err := newPdf.addComponentsToPdf(template)
	if err != nil {
		return nil, err
	}

	return &newPdf, nil
}

// addComponentsToPdf is responsible for adding all the components that are part of the template to the PDF.
// This is the method where the components that are part of the PDF are created.
func (d *Document) addComponentsToPdf(templates map[string]interface{}) error {
	fieldMappers := d.getFieldMappers()

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
func (d *Document) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"builder": d.setBuilder,
		"header":  d.setHeader,
		"footer":  d.setFooter,
		"pages":   d.setPages,
	}
}

// setBuilder is responsible for factories builder information
func (d *Document) setBuilder(builderDoc interface{}) error {
	builder, err := buildermapper.NewBuilder(builderDoc)
	if err != nil {
		return err
	}

	d.Builder = *builder
	return nil
}

// setHeader is responsible for factory the header
func (d *Document) setHeader(rowsDoc interface{}) error {
	rowsTemplate, ok := rowsDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("header cannot be deserialized, ensure header has an array of rows")
	}

	for templateKey, rowTemplate := range rowsTemplate {
		row, err := d.factory.NewRow(rowTemplate, templateKey)
		if err != nil {
			return err
		}
		d.Header = append(d.Header, row)
	}

	return nil
}

// setFooter is responsible for factory the footer
func (d *Document) setFooter(rowsDoc interface{}) error {
	rowsTemplate, ok := rowsDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("footer cannot be deserialized, ensure footer has an array of rows")
	}

	for templateKey, rowTemplate := range rowsTemplate {
		row, err := d.factory.NewRow(rowTemplate, templateKey)
		if err != nil {
			return err
		}
		d.Footer = append(d.Footer, row)
	}

	return nil
}

// setPages is responsible for factory the pages.
// pages can be a list of pages or just one page
func (d *Document) setPages(pagesDoc interface{}) error {
	templatePage, ok := pagesDoc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("ensure pages can be converted to map[string] interface{}")
	}

	d.Pages = make([]mappers.Componentmapper, len(templatePage))
	for templateName, template := range templatePage {
		page, err := d.factoryPage(template, templateName)
		if err != nil {
			return err
		}
		if err := d.addPage(page); err != nil {
			return err
		}
	}

	return nil
}

// addPage is responsible for validating and adding the page to the template
func (d *Document) addPage(page mappers.OrderedComponents) error {
	order := page.GetOrder()
	if page.GetOrder() > len(d.Pages) {
		return fmt.Errorf("component order cannot be greater than %d, this is the number of components in the template", len(d.Pages))
	}
	if d.Pages[order-1] != nil {
		return fmt.Errorf("cannot create document template, component order cannot be repeated")
	}

	d.Pages[order-1] = page
	return nil
}

// factoryPage is responsible for making a template of a page or a list of pages
func (d *Document) factoryPage(template interface{}, templateName string) (mappers.OrderedComponents, error) {
	var page mappers.OrderedComponents
	var err error

	if strings.HasPrefix(templateName, "list") {
		page, err = d.factory.NewList(template, templateName, d.factory.NewPage)
	} else {
		page, err = d.factory.NewPage(template, templateName)
	}

	if err != nil {
		return nil, err
	}
	return page, nil
}

func (d *Document) GetBuilderCfg() *buildermapper.Builder {
	return &d.Builder
}

// getContent is responsible for obtaining a content with a matching key, when the content is not found,
// an empty array is returned
func (d *Document) getContent(content map[string]interface{}, key string) (map[string]interface{}, error) {
	doc, ok := content[key]
	if ok {
		docArr, ok := doc.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("ensure that the contents of the %s can be converted to map[string]interface{}", key)
		}
		return docArr, nil
	}
	return make(map[string]interface{}, 0), nil
}

// generatePages is responsible for creating and adding pages, if the content does not have a page field
// an empty array is sent to the page
func (d *Document) generatePages(content map[string]interface{}, provider processorprovider.ProcessorProvider) error {
	pageContent, err := d.getContent(content, "pages")
	if err != nil {
		return err
	}

	pagesComponents := make([]processorprovider.ProviderComponent, 0, len(d.Pages))
	for _, pageTemplate := range d.Pages {
		page, err := pageTemplate.Generate(pageContent, provider)
		if err != nil {
			return err
		}
		pagesComponents = append(pagesComponents, page...)
	}

	_, err = provider.AddPages(pagesComponents...)
	return err
}

// generateRows is responsible for creating row components, it will extract the content (if it exists)
// and send this content to rows
func (d *Document) generateRows(content map[string]interface{}, provider processorprovider.ProcessorProvider,
	sourceKey string, templateRows ...mappers.Componentmapper,
) ([]processorprovider.ProviderComponent, error) {
	headerContent, err := d.getContent(content, sourceKey)
	if err != nil {
		return nil, err
	}

	rows := make([]processorprovider.ProviderComponent, 0, len(templateRows))
	for _, row := range templateRows {
		componentRow, err := row.Generate(headerContent, provider)
		if err != nil {
			return nil, err
		}
		rows = append(rows, componentRow...)
	}
	return rows, nil
}

// generate is responsible for the builder pdf according to the submitted content
func (d *Document) Generate(content map[string]interface{},
	provider processorprovider.ProcessorProvider) (*processorprovider.ProcessorProvider, error,
) {
	if len(d.Header) > 0 {
		header, err := d.generateRows(content, provider, "header", d.Header...)
		if err != nil {
			return nil, err
		}
		_, err = provider.AddHeader(header...)
		if err != nil {
			return nil, err
		}
	}

	if len(d.Footer) > 0 {
		footer, err := d.generateRows(content, provider, "footer", d.Footer...)
		if err != nil {
			return nil, err
		}
		_, err = provider.AddFooter(footer...)
		if err != nil {
			return nil, err
		}
	}

	if len(d.Pages) > 0 {
		if err := d.generatePages(content, provider); err != nil {
			return nil, err
		}
	}
	return &provider, nil
}
