// Package codemapper implements creation of matrixCode, MatrixCode and QrCode.
// nolint:dupl
package codemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/order"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Matrixcode struct {
	SourceKey string
	Code      string
	Props     *propsmapper.Rect
	Order     int
}

func NewMatrixcode(code interface{}) (*Matrixcode, error) {
	matrixcodeMap, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure matrixcode can be converted to map[string] interface{}")
	}

	matrixCode := &Matrixcode{}
	if err := matrixCode.addFields(matrixcodeMap); err != nil {
		return nil, err
	}
	if err := matrixCode.validateFields(); err != nil {
		return nil, err
	}

	return matrixCode, nil
}

// addFields is responsible for adding the matrix code fields according to
// the properties informed in the map
func (m *Matrixcode) addFields(codeMap map[string]interface{}) error {
	order, err := order.SetPageOrder(&codeMap, "matrixcode", m.SourceKey)
	if err != nil {
		return err
	}
	m.Order = order
	fieldMappers := m.getFieldMappers()

	for templateName, template := range codeMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the matrix code cannot be mapped to any valid field", templateName)
		}
		err := mapper(template)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetOrder is responsible for returning the component's defined order
func (m *Matrixcode) GetOrder() int {
	return m.Order
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (m *Matrixcode) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": m.setSourceKey,
		"value":      m.setCode,
		"props":      m.setProps,
	}
}

func (m *Matrixcode) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	m.SourceKey = sourceKey
	return nil
}

func (m *Matrixcode) setCode(template interface{}) error {
	code, ok := template.(string)
	if !ok {
		return fmt.Errorf("code cannot be converted to a string")
	}
	m.Code = code
	return nil
}

func (m *Matrixcode) setProps(template interface{}) error {
	props, err := propsmapper.NewRect(template)
	if err != nil {
		return err
	}
	m.Props = props
	return nil
}

func (m *Matrixcode) validateFields() error {
	if m.Code == "" && m.SourceKey == "" {
		return fmt.Errorf("no value passed for matrixCode. Add the 'source_key' or a code")
	}
	return nil
}

func (m *Matrixcode) getCode(content map[string]interface{}) (string, error) {
	if m.Code != "" {
		return m.Code, nil
	}
	codeFound, ok := content[m.SourceKey]
	if !ok {
		return "", fmt.Errorf("matrixcode requires a source key named %s, but it was not found", m.SourceKey)
	}
	codeValid, ok := codeFound.(string)
	if !ok {
		return "", fmt.Errorf("unable to generate matrixcode, invalid code. source key %s", m.SourceKey)
	}
	return codeValid, nil
}

func (m *Matrixcode) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	code, err := m.getCode(content)
	if err != nil {
		return nil, err
	}
	m.Code = code

	if m.Props != nil {
		return []processorprovider.ProviderComponent{provider.CreateMatrixCode(m.Code, m.Props)}, nil
	}
	return []processorprovider.ProviderComponent{provider.CreateMatrixCode(m.Code)}, nil
}
