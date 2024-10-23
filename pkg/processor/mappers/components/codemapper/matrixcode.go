// Package codemapper implements creation of matrixCode, MatrixCode and QrCode.
package codemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Matrixcode struct {
	SourceKey string
	Code      string
	Props     *propsmapper.Rect
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
func (m *Matrixcode) addFields(mapRows map[string]interface{}) error {
	fieldMappers := m.getFieldMappers()

	for templateName, template := range mapRows {
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

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (m *Matrixcode) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": m.setSourceKey,
		"code":       m.setCode,
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

func (m *Matrixcode) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
