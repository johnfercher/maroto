package repository

import (
	"os"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type Repository interface {
	AddUTF8Font(family string, style fontstyle.Type, file string) Repository
	Load() ([]*entity.CustomFont, error)
}

type repository struct {
	customFonts []*entity.CustomFont
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) AddUTF8Font(family string, style fontstyle.Type, file string) Repository {
	if family == "" {
		return r
	}

	if !style.IsValid() {
		return r
	}

	if file == "" {
		return r
	}

	r.customFonts = append(r.customFonts, &entity.CustomFont{
		Family: family,
		Style:  style,
		File:   file,
	})

	return r
}

func (r *repository) Load() ([]*entity.CustomFont, error) {
	for _, customFont := range r.customFonts {
		bytes, err := os.ReadFile(customFont.File)
		if err != nil {
			return nil, err
		}
		customFont.Bytes = bytes
	}
	return r.customFonts, nil
}
