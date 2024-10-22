// codemapper code implements creation of Barcode, MatrixCode and QrCode.
package codemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Barcode struct {
	SourceKey string
	Code      string
	Props     *propsmapper.Barcode
}

func NewBarcode(code interface{}) (*Barcode, error) {
	barcodeMap, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure barcode can be converted to map[string] interface{}")
	}

	barcode := &Barcode{}
	if err := barcode.addFields(barcodeMap); err != nil {
		return nil, err
	}
	if err := barcode.validateFields(); err != nil {
		return nil, err
	}

	return barcode, nil
}

// addFields is responsible for adding the barcode fields according to
// the properties informed in the map
func (b *Barcode) addFields(mapRows map[string]interface{}) error {
	fieldMappers := b.getFieldMappers()

	for templateName, template := range mapRows {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the barcode cannot be mapped to any valid field", templateName)
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
func (b *Barcode) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": b.setSourceKey,
		"code":       b.setCode,
		"props":      b.setProps,
	}
}

func (b *Barcode) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	b.SourceKey = sourceKey
	return nil
}

func (b *Barcode) setCode(template interface{}) error {
	code, ok := template.(string)
	if !ok {
		return fmt.Errorf("code cannot be converted to a string")
	}
	b.Code = code
	return nil
}

func (b *Barcode) setProps(template interface{}) error {
	props, err := propsmapper.NewBarcode(template)
	if err != nil {
		return err
	}
	b.Props = props
	return nil
}

func (b *Barcode) validateFields() error {
	if b.Code == "" && b.SourceKey == "" {
		return fmt.Errorf("no value passed for barcode. Add the 'source_key' or a code")
	}
	return nil
}

func (b *Barcode) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
