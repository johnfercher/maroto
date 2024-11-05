// Package image implements creation of images from file and bytes.
package imagemapper

import (
	"fmt"

	processorcore "github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Image struct {
	Path       string
	SourceKey  string
	Props      *propsmapper.Rect
	Repository processorcore.ProcessorRepository
}

func NewImage(templateImage interface{}, repository processorcore.ProcessorRepository) (*Image, error) {
	imageMap, ok := templateImage.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure image can be converted to map[string] interface{}")
	}

	image := &Image{Repository: repository}
	if err := image.addFields(imageMap); err != nil {
		return nil, err
	}
	if image.SourceKey == "" {
		return nil, fmt.Errorf("no value passed for image. Add the 'source_key' or a path")
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
		"path":       i.setPath,
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
	i.Path = path
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

func (i *Image) getImagePath(content map[string]interface{}) (string, error) {
	if i.Path != "" {
		return i.Path, nil
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

func (i *Image) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error) {
	path, err := i.getImagePath(content)
	if err != nil {
		return nil, err
	}
	i.Path = path
	extension, img, err := i.Repository.GetDocument(i.Path)
	if err != nil {
		return nil, err
	}

	if i.Props != nil {
		return []processorprovider.ProviderComponent{provider.CreateImage(img, extension, i.Props)}, nil
	}
	return []processorprovider.ProviderComponent{provider.CreateImage(img, extension)}, nil
}
