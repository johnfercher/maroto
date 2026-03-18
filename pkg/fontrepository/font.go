package fontrepository

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

type customFont struct {
	family string
	style  fontstyle.Type
	file   string
	bytes  []byte
}

func (c *customFont) GetFamily() string        { return c.family }
func (c *customFont) GetStyle() fontstyle.Type { return c.style }
func (c *customFont) GetFile() string          { return c.file }
func (c *customFont) GetBytes() []byte         { return c.bytes }
