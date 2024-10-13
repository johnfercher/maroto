// Package image implements creation of images from file and bytes.
package imagemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Image struct{}

func NewImage(code interface{}) (*Image, error) {
	return nil, nil
}

func (b *Image) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
