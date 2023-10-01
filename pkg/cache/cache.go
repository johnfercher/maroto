package cache

import (
	"errors"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type Cache interface {
	GetImage(value string, extension extension.Type) (*entity.Image, error)
	LoadImage(value string, extension extension.Type) error
	GetCode(code string, codeType string) ([]byte, error)
	SaveCode(code string, codeType string, bytes []byte)
}

type cache struct {
	images map[string]*entity.Image
	codes  map[string][]byte
}

func New() Cache {
	return &cache{
		images: make(map[string]*entity.Image),
		codes:  make(map[string][]byte),
	}
}

func (c *cache) LoadImage(file string, extension extension.Type) error {
	imageBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	img := &entity.Image{Bytes: imageBytes, Extension: extension}
	c.images[file+string(extension)] = img

	return nil
}

func (c *cache) GetImage(file string, extension extension.Type) (*entity.Image, error) {
	image, ok := c.images[file+string(extension)]
	if ok {
		return image, nil
	}

	return nil, errors.New("image not found")
}

func (c *cache) GetCode(code string, codeType string) ([]byte, error) {
	bytes, ok := c.codes[code+codeType]
	if ok {
		return bytes, nil
	}

	return nil, errors.New("code not found")
}

func (c *cache) SaveCode(code string, codeType string, bytes []byte) {
	c.codes[code+codeType] = bytes
}
