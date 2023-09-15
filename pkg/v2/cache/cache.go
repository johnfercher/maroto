package cache

import (
	"encoding/base64"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"os"
)

type Cache interface {
	Load(string, consts.Extension) (*domain.Image, error)
}

type cache struct {
	images map[string]*domain.Image
}

func New() Cache {
	return &cache{
		images: make(map[string]*domain.Image),
	}
}

func (c *cache) Load(value string, extension consts.Extension) (*domain.Image, error) {
	if _, ok := c.images[value]; !ok {
		if _, err := base64.StdEncoding.DecodeString(value); err == nil {
			img := &domain.Image{Value: value, Extension: extension}
			c.images[value] = img
			return img, nil
		} else {
			imageBytes, err := os.ReadFile(value)
			if err != nil {
				return nil, err
			}
			valueString := base64.StdEncoding.EncodeToString(imageBytes)
			img := &domain.Image{Value: valueString, Extension: extension}
			c.images[value] = img
		}
	}
	return c.images[value], nil
}
