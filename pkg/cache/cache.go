package cache

import (
	"encoding/base64"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"os"
	"sync"
)

type Cache interface {
	Load(string, extension.Type) (*Image, error)
}

type Image struct {
	Value     string
	Extension extension.Type
}

type cache struct {
	images map[string]*Image
	mutex  sync.RWMutex
}

func New() Cache {
	return &cache{
		images: make(map[string]*Image),
		mutex:  sync.RWMutex{},
	}
}

func (c *cache) Load(value string, extension extension.Type) (*Image, error) {
	if _, ok := c.images[value]; !ok {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		if _, err := base64.StdEncoding.DecodeString(value); err == nil {
			img := &Image{Value: value, Extension: extension}
			c.images[value] = img
			return img, nil
		} else {
			imageBytes, err := os.ReadFile(value)
			if err != nil {
				return nil, err
			}
			valueString := base64.StdEncoding.EncodeToString(imageBytes)
			img := &Image{Value: valueString, Extension: extension}
			c.images[value] = img
		}
	}
	return c.images[value], nil
}
