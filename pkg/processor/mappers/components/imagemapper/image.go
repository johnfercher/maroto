// Package image implements creation of images from file and bytes.
package imagemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/order"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Image struct {
	Value     string
	SourceKey string
	Props     *propsmapper.Rect
	Order     int
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
	if image.SourceKey == "" && image.Value == "" {
		return nil, fmt.Errorf("no value passed for image. Add the 'source_key' or a value")
	}

	return image, nil
}

// addFields is responsible for adding the barcode fields according to
// the properties informed in the map
func (i *Image) addFields(imageMap map[string]interface{}) error {
	order, err := order.SetPageOrder(&imageMap, "image", i.SourceKey)
	if err != nil {
		return err
	}
	i.Order = order
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

// GetOrder is responsible for returning the component's defined order
func (i *Image) GetOrder() int {
	return i.Order
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (i *Image) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": i.setSourceKey,
		"props":      i.setProps,
		"value":      i.setPath,
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

func (i *Image) setPath(template interface{}) error {
	path, ok := template.(string)
	if !ok {
		return fmt.Errorf("path cannot be converted to a string")
	}
	i.Value = path
	return nil
}

func (i *Image) setProps(template interface{}) error {
	props, err := propsmapper.NewRect(template)
	if err != nil {
		return err
	}
	i.Props = props
	return nil
}

func (i *Image) getImagePath(content map[string]interface{}) (string, error) {
	if i.Value != "" {
		return i.Value, nil
	}
	imageFound, ok := content[i.SourceKey]
	if !ok {
		return "", fmt.Errorf("image requires a source key named %s, but it was not found", i.SourceKey)
	}
	imageValid, ok := imageFound.(string)
	if !ok {
		return "", fmt.Errorf("unable to generate image, invalid path. source key %s", i.SourceKey)
	}
	return imageValid, nil
}

func (i *Image) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) (
	[]processorprovider.ProviderComponent, error,
) {
	var err error
	i.Value, err = i.getImagePath(content)
	if err != nil {
		return nil, err
	}

	var img processorprovider.ProviderComponent
	if i.Props != nil {
		img, err = provider.CreateImage(i.Value, i.Props)
	} else {
		img, err = provider.CreateImage(i.Value)
	}

	return []processorprovider.ProviderComponent{img}, err
}
