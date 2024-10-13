// listmapper is the package responsible for mapping row settings
package listmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

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

func (r *List) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
