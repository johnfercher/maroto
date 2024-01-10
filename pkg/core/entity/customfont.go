// Package contains all core entities.
package entity

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

type CustomFont struct {
	Family string
	Style  fontstyle.Type
	File   string
	Bytes  []byte
}
