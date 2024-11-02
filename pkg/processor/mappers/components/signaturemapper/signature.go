// Package signature implements creation of signatures.
package signaturemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Signature struct {
	SourceKey string
	Value     string
	Props     *propsmapper.Signature
}

func NewSignature(code interface{}) (*Signature, error) {
	signatureMap, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure signature can be converted to map[string] interface{}")
	}

	signature := &Signature{}
	if err := signature.addFields(signatureMap); err != nil {
		return nil, err
	}
	if err := signature.validateFields(); err != nil {
		return nil, err
	}

	return signature, nil
}

// addFields is responsible for adding the signature fields according to
// the properties informed in the map
func (s *Signature) addFields(signatureMap map[string]interface{}) error {
	fieldMappers := s.getFieldMappers()

	for templateName, template := range signatureMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the signature cannot be mapped to any valid field", templateName)
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
func (s *Signature) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": s.setSourceKey,
		"value":      s.setValue,
		"props":      s.setProps,
	}
}

func (s *Signature) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	s.SourceKey = sourceKey
	return nil
}

func (s *Signature) setValue(template interface{}) error {
	value, ok := template.(string)
	if !ok {
		return fmt.Errorf("value cannot be converted to a string")
	}
	s.Value = value
	return nil
}

func (s *Signature) setProps(template interface{}) error {
	props, err := propsmapper.NewSignature(template)
	if err != nil {
		return err
	}
	s.Props = props
	return nil
}

func (s *Signature) validateFields() error {
	if s.Value == "" && s.SourceKey == "" {
		return fmt.Errorf("no value passed for signature. Add the 'source_key' or a value")
	}
	return nil
}

func (s *Signature) getSignature(content map[string]interface{}) (string, error) {
	if s.Value != "" {
		return s.Value, nil
	}
	signatureFound, ok := content[s.SourceKey]
	if !ok {
		return "", fmt.Errorf("signature requires a source key named %s, but it was not found", s.SourceKey)
	}
	signatureValid, ok := signatureFound.(string)
	if !ok {
		return "", fmt.Errorf("unable to generate signature, invalid value. source key %s", s.SourceKey)
	}
	return signatureValid, nil
}

func (s *Signature) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (processorprovider.PDFComponent, error) {
	signature, err := s.getSignature(content)
	if err != nil {
		return nil, err
	}
	s.Value = signature

	if s.Props != nil {
		return provider.CreateSignature(s.Value, s.Props), nil
	}
	return provider.CreateSignature(s.Value), nil
}
