package entity

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
)

const trimSize = 10

// Image is the representation of an image that can be added to the pdf.
type Image struct {
	Bytes      []byte
	Extension  extension.Type
	Dimensions *Dimensions
}

// AppendMap adds the Image fields to the map.
func (i *Image) AppendMap(m map[string]interface{}) map[string]interface{} {
	lenBytes := len(i.Bytes)
	if lenBytes != 0 {
		trimSize := trimSize
		if lenBytes < trimSize {
			trimSize = lenBytes
		}
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
