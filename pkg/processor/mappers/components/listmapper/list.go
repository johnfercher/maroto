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
	Templates []mappers.Componentmapper
}

func NewList(list interface{}, sourceKey string, generate mappers.GenerateComponent) (*List, error) {
	listMapper, ok := list.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure list can be converted to map[string] interface{}")
	}

	components, err := createComponents(listMapper, generate)
	if err != nil {
		return nil, err
	}
	return &List{
		Templates: components,
		SourceKey: sourceKey,
	}, nil
}

func createComponents(listMapper map[string]interface{}, generate mappers.GenerateComponent) ([]mappers.Componentmapper, error) {
	components := make([]mappers.Componentmapper, len(listMapper))
	cont := 0
	for templateName, template := range listMapper {
		component, err := generate(template, templateName)
		if err != nil {
			return nil, err
		}
		components[cont] = component
		cont++
	}
	return components, nil
}

func (l *List) getListContent(content map[string]interface{}) ([]map[string]interface{}, error) {
	listContent, ok := content[l.SourceKey]
	if !ok {
		return nil, fmt.Errorf("the list needs the source key \"%s\", but it was not found", l.SourceKey)
	}
	if contents, ok := listContent.([]map[string]interface{}); ok {
		return contents, nil
	}
	return nil, fmt.Errorf("ensure that the contents of the list \"%s\" can be converted to []map[string]interface{}", l.SourceKey)
}

func (l *List) generateTemplates(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error) {
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

func (l *List) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error) {
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
