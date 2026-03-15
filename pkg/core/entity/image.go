package entity

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
)

const trimSizeDefault = 10

// Image is the representation of an image that can be added to the pdf.
type Image struct {
	Bytes      []byte
	Extension  extension.Type
	Dimensions *Dimensions
}

// AppendMap adds the Image fields to the map.
func (i *Image) AppendMap(m map[string]any) map[string]any {
	lenBytes := len(i.Bytes)
	if lenBytes != 0 {
		trimSize := min(lenBytes, trimSizeDefault)
		m["entity_image_bytes"] = fmt.Sprintf("%v", i.Bytes[:trimSize])
	}

	if i.Extension != "" {
		m["entity_extension"] = i.Extension
	}

	if i.Dimensions != nil {
		m = i.Dimensions.AppendMap("background", m)
	}

	return m
}
