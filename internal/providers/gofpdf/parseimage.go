package gofpdf

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

func FromBytes(bytes []byte, ext extension.Type) (*entity.Image, error) {
	if !ext.IsValid() {
		return nil, errors.New("invalid image format")
	}

	return &entity.Image{
		Bytes:     bytes,
		Extension: ext,
	}, nil
}
