// Package codemapper implements creation of qrcode, MatrixCode and QrCode.
// nolint:dupl
package codemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Qrcode struct {
	SourceKey string
	Code      string
	Props     *propsmapper.Rect
}

func NewQrcode(code interface{}) (*Qrcode, error) {
	qrcodeMap, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure qrcode can be converted to map[string] interface{}")
	}

	qrcode := &Qrcode{}
	if err := qrcode.addFields(qrcodeMap); err != nil {
		return nil, err
	}
	if err := qrcode.validateFields(); err != nil {
		return nil, err
	}

	return qrcode, nil
}

// addFields is responsible for adding the qrcode fields according to
// the properties informed in the map
func (q *Qrcode) addFields(codeMap map[string]interface{}) error {
	fieldMappers := q.getFieldMappers()

	for templateName, template := range codeMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the qrcode cannot be mapped to any valid field", templateName)
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
func (q *Qrcode) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": q.setSourceKey,
		"code":       q.setCode,
		"props":      q.setProps,
	}
}

func (q *Qrcode) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	q.SourceKey = sourceKey
	return nil
}

func (q *Qrcode) setCode(template interface{}) error {
	code, ok := template.(string)
	if !ok {
		return fmt.Errorf("code cannot be converted to a string")
	}
	q.Code = code
	return nil
}

func (q *Qrcode) setProps(template interface{}) error {
	props, err := propsmapper.NewRect(template)
	if err != nil {
		return err
	}
	q.Props = props
	return nil
}

func (q *Qrcode) validateFields() error {
	if q.Code == "" && q.SourceKey == "" {
		return fmt.Errorf("no value passed for qrcode. Add the 'source_key' or a code")
	}
	return nil
}

func (q *Qrcode) getCode(content map[string]interface{}) (string, error) {
	if q.Code != "" {
		return q.Code, nil
	}
	codeFound, ok := content[q.SourceKey]
	if !ok {
		return "", fmt.Errorf("qrcode requires a source key named %s, but it was not found", q.SourceKey)
	}
	codeValid, ok := codeFound.(string)
	if !ok {
		return "", fmt.Errorf("unable to generate qrcode, invalid code. source key %s", q.SourceKey)
	}
	return codeValid, nil
}

func (q *Qrcode) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error) {
	code, err := q.getCode(content)
	if err != nil {
		return nil, err
	}
	q.Code = code

	if q.Props != nil {
		return []processorprovider.ProviderComponent{provider.CreateQrCode(q.Code, q.Props)}, nil
	}
	return []processorprovider.ProviderComponent{provider.CreateQrCode(q.Code)}, nil
}
