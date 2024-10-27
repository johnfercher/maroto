// Package image implements creation of images from file and bytes.
package imagemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Image struct {
	SourceKey string
	Props     *propsmapper.Rect
}

func NewImage(templateImage interface{}) (*Image, error) {
	imageMap, ok := templateImage.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure image can be converted to map[string] interface{}")
	}

	image := &Image{}
	if err := image.addFields(imageMap); err != nil {
		return nil, err
	}
	if image.SourceKey == "" {
		return nil, fmt.Errorf("no value passed for image. Add the 'source_key'")
	}

	return image, nil

}

// addFields is responsible for adding the barcode fields according to
// the properties informed in the map
func (i *Image) addFields(imageMap map[string]interface{}) error {
	fieldMappers := i.getFieldMappers()

	for templateName, template := range imageMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the image cannot be mapped to any valid field", templateName)
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
func (i *Image) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": i.setSourceKey,
		"props":      i.setProps,
	}
}

func (i *Image) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	i.SourceKey = sourceKey
	return nil
}

func (b *Image) setProps(template interface{}) error {
	props, err := propsmapper.NewRect(template)
	if err != nil {
		return err
	}
	b.Props = props
	return nil
}

func (b *Image) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
