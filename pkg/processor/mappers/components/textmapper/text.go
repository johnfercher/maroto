package textmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Text struct {
	SourceKey string
	Value     string
	Props     *propsmapper.Text
}

func NewText(templateText interface{}) (*Text, error) {
	textMap, ok := templateText.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure text can be converted to map[string] interface{}")
	}

	text := &Text{}
	if err := text.addFields(textMap); err != nil {
		return nil, err
	}
	if err := text.validateFields(); err != nil {
		return nil, err
	}

	return text, nil
}

// addFields is responsible for adding the text fields according to
// the properties informed in the map
func (t *Text) addFields(valueMap map[string]interface{}) error {
	fieldMappers := t.getFieldMappers()

	for templateName, template := range valueMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the text cannot be mapped to any valid field", templateName)
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
func (t *Text) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": t.setSourceKey,
		"value":      t.setValue,
		"props":      t.setProps,
	}
}

func (t *Text) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	t.SourceKey = sourceKey
	return nil
}

func (t *Text) setValue(template interface{}) error {
	value, ok := template.(string)
	if !ok {
		return fmt.Errorf("value cannot be converted to a string")
	}
	t.Value = value
	return nil
}

func (t *Text) setProps(template interface{}) error {
	props, err := propsmapper.NewText(template)
	if err != nil {
		return err
	}
	t.Props = props
	return nil
}

func (t *Text) validateFields() error {
	if t.Value == "" && t.SourceKey == "" {
		return fmt.Errorf("no value passed for text. Add the 'source_key' or a value")
	}
	return nil
}

func (t *Text) getValue(content map[string]interface{}) (string, error) {
	if t.Value != "" {
		return t.Value, nil
	}
	textFound, ok := content[t.SourceKey]
	if !ok {
		return "", fmt.Errorf("text requires a source key named %s, but it was not found", t.SourceKey)
	}
	textValid, ok := textFound.(string)
	if !ok {
		return "", fmt.Errorf("unable to generate text, invalid value. source key %s", t.SourceKey)
	}
	return textValid, nil
}

func (t *Text) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (processorprovider.PDFComponent, error) {
	signature, err := t.getValue(content)
	if err != nil {
		return nil, err
	}
	t.Value = signature

	if t.Props != nil {
		return provider.CreateText(t.Value, t.Props), nil
	}
	return provider.CreateText(t.Value), nil
}
