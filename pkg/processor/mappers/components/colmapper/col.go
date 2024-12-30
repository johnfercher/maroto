package colmapper

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type factoryComponent = func(interface{}) (mappers.OrderedComponents, error)

type Col struct {
	Size       int
	Components []mappers.OrderedComponents
	factory    mappers.AbstractFactoryMaps
	props      *propsmapper.Cell
}

func NewCol(templateCols interface{}, factory mappers.AbstractFactoryMaps) (*Col, error) {
	mapCols, ok := templateCols.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure that rows can be converted to map[string] interface{}")
	}

	col := &Col{Size: 0, Components: make([]mappers.OrderedComponents, 0), factory: factory}

	if err := col.setProps(&mapCols); err != nil {
		return nil, err
	}

	if err := col.setSize(&mapCols); err != nil {
		return nil, err
	}

	if err := col.addComponents(mapCols); err != nil {
		return nil, err
	}
	return col, nil
}

func (c *Col) addComponents(mapComponents map[string]interface{}) error {
	fieldMappers := c.getFieldMappers()

	c.Components = make([]mappers.OrderedComponents, len(mapComponents))
	for templateName, template := range mapComponents {
		mapper, err := c.getFieldMapperByTemplateName(templateName, fieldMappers)
		if err != nil {
			return err
		}
		component, err := mapper(template)
		if err != nil {
			return err
		}

		if err := c.addComponent(component); err != nil {
			return err
		}
	}
	return nil
}

// addComponent is responsible for validating and adding the component to the template
func (c *Col) addComponent(component mappers.OrderedComponents) error {
	order := component.GetOrder()
	if order > len(c.Components) {
		return fmt.Errorf("component order cannot be greater than %d, this is the number of components in the template", len(c.Components))
	}
	if c.Components[order-1] != nil {
		return fmt.Errorf("cannot create col template, component order cannot be repeated")
	}

	c.Components[order-1] = component
	return nil
}

func (c *Col) setSize(template *map[string]interface{}) error {
	defer delete(*template, "size")
	templateSize, ok := (*template)["size"]
	if ok {
		size, ok := templateSize.(float64)
		if !ok {
			return fmt.Errorf("ensure that size can be converted to int")
		}
		c.Size = int(size)
	}
	return nil
}

// setProps is responsible for creating template col props
func (c *Col) setProps(template *map[string]interface{}) error {
	props, ok := (*template)["props"]
	if !ok {
		return nil
	}
	defer delete(*template, "props")

	propsRow, err := propsmapper.NewCell(props)
	if err != nil {
		return err
	}
	c.props = propsRow
	return nil
}

func (c *Col) getFieldMapperByTemplateName(templateName string, mappers map[string]factoryComponent) (factoryComponent, error) {
	for mapperName, mapper := range mappers {
		if strings.HasPrefix(templateName, mapperName) {
			return mapper, nil
		}
	}
	return nil, fmt.Errorf(
		"the field \"%s\" present in the col cannot be mapped to any valid component, ensure the field name starts with a valid component",
		templateName)
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (c *Col) getFieldMappers() map[string]factoryComponent {
	return map[string]factoryComponent{
		"bar_code":    c.factory.NewBarcode,
		"matrix_code": c.factory.NewMatrixcode,
		"qr_code":     c.factory.NewQrcode,
		"image":       c.factory.NewImage,
		"line":        c.factory.NewLine,
		"signature":   c.factory.NewSignature,
		"text":        c.factory.NewText,
	}
}

// Generate is responsible for generating the col component, it will generate all the internal components and add them to the col
//   - The content is a map with the properties of the col components
func (c *Col) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	components := make([]processorprovider.ProviderComponent, 0, len(c.Components))
	for _, component := range c.Components {
		newComponent, err := component.Generate(content, provider)
		if err != nil {
			return nil, err
		}
		components = append(components, newComponent...)
	}

	col, err := provider.CreateCol(c.Size, c.props, components...)
	if err != nil {
		return nil, err
	}
	return []processorprovider.ProviderComponent{col}, nil
}
