package gofpdf

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var ErrInvalidImageFormat = errors.New("invalid image format")

func FromBytes(bytes []byte, ext extension.Type) (*entity.Image, error) {
	if !ext.IsValid() {
		return nil, ErrInvalidImageFormat
	}

	return &entity.Image{
		Bytes:     bytes,
		Extension: ext,
	}, nil
}
