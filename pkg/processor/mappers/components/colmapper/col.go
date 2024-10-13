package colmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type Col struct {
	Size       int
	Components []mappers.Componentmapper
	factory    mappers.AbstractFactoryMaps
}

func NewCol(templateCols interface{}, factory mappers.AbstractFactoryMaps) (*Col, error) {
	mapCols, ok := templateCols.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure that rows can be converted to map[string] interface{}")
	}

	col := &Col{Size: 0, Components: make([]mappers.Componentmapper, 0), factory: factory}

	if err := col.setSize(&mapCols); err != nil {
		return nil, err
	}

	if err := col.addComponents(mapCols); err != nil {
		return nil, err
	}
	return col, nil
}

func (c *Col) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}

func (c *Col) addComponents(mapComponents map[string]interface{}) error {
	fieldMappers := c.getFieldMappers()

	for templateName, template := range mapComponents {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the col cannot be mapped to any valid component", templateName)
		}
		component, err := mapper(template)
		if err != nil {
			return err
		}
		c.Components = append(c.Components, component)
	}
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

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (c *Col) getFieldMappers() map[string]func(interface{}) (mappers.Componentmapper, error) {
	return map[string]func(interface{}) (mappers.Componentmapper, error){
		"bar_code":    c.factory.NewBarcode,
		"matrix_code": c.factory.NewMatrixcode,
		"qr_code":     c.factory.NewQrcode,
		"image":       c.factory.NewImage,
		"line":        c.factory.NewLine,
		"signature":   c.factory.NewSignature,
		"text":        c.factory.NewText,
	}
}
