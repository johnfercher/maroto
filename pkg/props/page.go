package props

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type Place string

const (
	NorthWest Place = "north_west"
	North     Place = "north"
	NorthEast Place = "north_east"
	SouthWest Place = "south_west"
	South     Place = "south"
	SouthEast Place = "south_east"
)

func (p Place) IsValid() bool {
	return p == NorthWest || p == North || p == NorthEast ||
		p == SouthWest || p == South || p == SouthEast
}

type Page struct {
	Pattern string
	Place   Place
	Family  string
	Style   fontstyle.Type
	Size    float64
	Color   *Color
}

func (p *Page) GetNumberTextProp(height float64) *Text {
	text := &Text{
		Family: p.Family,
		Style:  p.Style,
		Size:   p.Size,
		Color:  p.Color,
		Align:  align.Center,
	}

	if p.Place == SouthWest || p.Place == NorthWest {
		text.Align = align.Left
	} else if p.Place == SouthEast || p.Place == NorthEast {
		text.Align = align.Right
	}

	if p.Place == SouthEast || p.Place == South || p.Place == SouthWest {
		text.Top = height
	}

	return text
}

func (p *Page) GetPageString(current, total int) string {
	pattern := strings.ReplaceAll(p.Pattern, "{current}", fmt.Sprintf("%d", current))
	return strings.ReplaceAll(pattern, "{total}", fmt.Sprintf("%d", total))
}
