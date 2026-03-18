// Package fontrepository implements font repository.
package fontrepository

import (
	"errors"
	"fmt"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var ErrCannotReadFile = errors.New("cannot read file")

// Repository is the abstraction to load custom fonts.
type Repository interface {
	AddUTF8Font(family string, style fontstyle.Type, file string) Repository
	AddUTF8FontFromBytes(family string, style fontstyle.Type, bytes []byte) Repository
	Load() ([]entity.CustomFont, error)
}

// FontRepository manages custom fonts to be loaded into a document.
type FontRepository struct {
	customFonts []*customFont
}

// New creates a new Repository.
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

	r.customFonts = append(r.customFonts, &customFont{
		family: family,
		style:  style,
		file:   file,
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

	r.customFonts = append(r.customFonts, &customFont{
		family: family,
		style:  style,
		bytes:  bytes,
	})

	return r
}

// Load loads all custom fonts, reading file contents from disk where needed.
func (r *FontRepository) Load() ([]entity.CustomFont, error) {
	for _, customFont := range r.customFonts {
		if customFont.file == "" {
			continue
		}
		bytes, err := os.ReadFile(customFont.file)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrCannotReadFile, err)
		}
		customFont.bytes = bytes
	}

	var customFonts []entity.CustomFont
	for _, customFont := range r.customFonts {
		customFonts = append(customFonts, customFont)
	}

	return customFonts, nil
}
