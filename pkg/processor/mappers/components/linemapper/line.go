// Package line implements creation of lines.
package linemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/order"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Line struct {
	Props *propsmapper.Line
	Order int
}

func NewLine(code interface{}) (*Line, error) {
	lineMapper, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure line can be converted to map[string] interface{}")
	}
	line := &Line{}

	if err := line.addFields(lineMapper); err != nil {
		return nil, err
	}
	return line, nil
}

// addFields is responsible for adding the barcode fields according to
// the properties informed in the map
func (l *Line) addFields(lineMapper map[string]interface{}) error {
	order, err := order.SetPageOrder(&lineMapper, "line", "")
	if err != nil {
		return err
	}
	l.Order = order
	fieldMappers := l.getFieldMappers()

	for templateName, template := range lineMapper {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the line cannot be mapped to any valid field", templateName)
		}
		err := mapper(template)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetOrder is responsible for returning the component's defined order
func (l *Line) GetOrder() int {
	return l.Order
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (l *Line) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"props": l.setProps,
	}
}

func (l *Line) setProps(templateProps interface{}) error {
	propsLine, err := propsmapper.NewLine(templateProps)
	if err != nil {
		return err
	}
	l.Props = propsLine
	return nil
}

func (l *Line) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	if l.Props != nil {
		return []processorprovider.ProviderComponent{provider.CreateLine(l.Props)}, nil
	}
	return []processorprovider.ProviderComponent{provider.CreateLine()}, nil
}
