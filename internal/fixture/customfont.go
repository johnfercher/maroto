package fixture

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type TestFont struct {
	Family string
	Style  fontstyle.Type
	File   string
	Bytes  []byte
}

func (t TestFont) GetFamily() string {
	return t.Family
}

func (t TestFont) GetStyle() fontstyle.Type {
	return t.Style
}

func (t TestFont) GetFile() string {
	return t.File
}

func (t TestFont) GetBytes() []byte {
	return t.Bytes
}
