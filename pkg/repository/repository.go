// Package repository implements font repository.
package repository

import (
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// Repository is the abstraction to load custom fonts.
type Repository interface {
	AddUTF8Font(family string, style fontstyle.Type, file string) Repository
	AddUTF8FontFromBytes(family string, style fontstyle.Type, bytes []byte) Repository
	Load() ([]*entity.CustomFont, error)
}

type FontRepository struct {
	customFonts []*entity.CustomFont
}

// New creates a new repository.
func New() Repository {
	return &FontRepository{}
}

// AddUTF8Font adds a custom font to the repository.
func (r *FontRepository) AddUTF8Font(family string, style fontstyle.Type, file string) Repository {
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

// AddUTF8FontFromBytes adds a custom font to the repository from a byte slice.
func (r *FontRepository) AddUTF8FontFromBytes(family string, style fontstyle.Type, bytes []byte) Repository {
	if family == "" {
		return r
	}

	if !style.IsValid() {
		return r
	}

	if bytes == nil {
		return r
	}

	r.customFonts = append(r.customFonts, &entity.CustomFont{
		Family: family,
		Style:  style,
		Bytes:  bytes,
	})

	return r
}

// Load loads all custom fonts.
func (r *FontRepository) Load() ([]*entity.CustomFont, error) {
	for _, customFont := range r.customFonts {
		if customFont.File == "" {
			continue
		}
		bytes, err := os.ReadFile(customFont.File)
		if err != nil {
			return nil, err
		}
		customFont.Bytes = bytes
	}
	return r.customFonts, nil
}
