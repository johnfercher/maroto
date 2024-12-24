// listmapper is the package responsible for mapping row settings
package listmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

// The List component is responsible for adding a list behavior to a component.
// It will repeat a component for each content sent in the generate method
type List struct {
	SourceKey string
	Templates []mappers.OrderedComponents
	order     int
}

func NewList(list interface{}, sourceKey string, generate mappers.GenerateComponent) (*List, error) {
	listMapper, ok := list.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure list can be converted to map[string] interface{}")
	}

	newList := List{SourceKey: sourceKey}
	if err := newList.setListOrder(&listMapper, sourceKey); err != nil {
		return nil, err
	}

	if err := newList.setComponents(listMapper, generate); err != nil {
		return nil, err
	}
	return &newList, nil
}

// setComponents is responsible for generating the list component. Components will be generated through the generate method
func (l *List) setComponents(listMapper map[string]interface{}, generate mappers.GenerateComponent) error {
	l.Templates = make([]mappers.OrderedComponents, len(listMapper))

	for templateName, template := range listMapper {
		component, err := generate(template, templateName)
		if err != nil {
			return err
		}
		if err := l.addComponent(component); err != nil {
			return err
		}
	}
	return nil
}

// addComponent is responsible for validating and adding the component to the template
func (l *List) addComponent(component mappers.OrderedComponents) error {
	order := component.GetOrder()
	if component.GetOrder() > len(l.Templates) {
		return fmt.Errorf("component order cannot be greater than %d, this is the number of components in the template", len(l.Templates))
	}
	if l.Templates[order-1] != nil {
		return fmt.Errorf("cannot create list template, component order cannot be repeated")
	}

	l.Templates[order-1] = component
	return nil
}

func (l *List) GetOrder() int {
	return l.order
}

// formatListContent is responsible for converting content into []map[string]interface{}
func (l *List) formatListContent(content interface{}) ([]map[string]interface{}, error) {
	listContent, ok := content.([]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure that the contents of the list \"%s\" can be converted to []map[string]interface{}", l.SourceKey)
	}

	contentMaps := make([]map[string]interface{}, 0, len(listContent))
	for _, content := range listContent {
		contentMap, ok := content.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("ensure that the contents of the list \"%s\" can be converted to []map[string]interface{}", l.SourceKey)
		}
		contentMaps = append(contentMaps, contentMap)
	}
	return contentMaps, nil
}

func (l *List) getListContent(content map[string]interface{}) ([]map[string]interface{}, error) {
	listContent, ok := content[l.SourceKey]
	if !ok {
		return nil, fmt.Errorf("the list needs the source key \"%s\", but it was not found", l.SourceKey)
	}

	return l.formatListContent(listContent)
}

func (l *List) generateTemplates(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	components := make([]processorprovider.ProviderComponent, 0, len(l.Templates))
	for _, template := range l.Templates {
		component, err := template.Generate(content, provider)
		if err != nil {
			return nil, err
		}
		components = append(components, component...)
	}
	return components, nil
}

func (l *List) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	listContent, err := l.getListContent(content)
	if err != nil {
		return nil, err
	}
	newComponents := make([]processorprovider.ProviderComponent, 0, len(l.Templates)*len(listContent))

	for _, content := range listContent {
		components, err := l.generateTemplates(content, provider)
		if err != nil {
			return nil, err
		}
		newComponents = append(newComponents, components...)
	}

	return newComponents, nil
}

// setListOrder is responsible for validating the component order and adding the order to the list
func (l *List) setListOrder(template *map[string]interface{}, sourceKey string) error {
	order, ok := (*template)["order"]
	if !ok {
		return fmt.Errorf("could not find field order on list \"%s\"", sourceKey)
	}
	validOrder, ok := order.(float64)
	if !ok || validOrder < 1 {
		return fmt.Errorf("the order field passed on list \"%s\" is not valid", sourceKey)
	}

	delete(*template, "order")
	l.order = int(validOrder)
	return nil
}
