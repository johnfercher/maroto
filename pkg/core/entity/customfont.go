package entity

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

// CustomFont is the representation of a font that can be added to the pdf.
type CustomFont interface {
	GetFamily() string
	GetStyle() fontstyle.Type
	GetFile() string
	GetBytes() []byte
}
