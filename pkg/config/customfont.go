package config

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

type CustomFont struct {
	Family string
	Style  fontstyle.Type
	File   string
}
