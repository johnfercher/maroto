// Package contains all props used to customize components.
package props

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type Place string

const (
	LeftTop     Place = "left_top"
	Top         Place = "top"
	RightTop    Place = "right_top"
	LeftBottom  Place = "left_bottom"
	Bottom      Place = "bottom"
	RightBottom Place = "right_bottom"
)

func (p Place) IsValid() bool {
	return p == LeftTop || p == Top || p == RightTop ||
		p == LeftBottom || p == Bottom || p == RightBottom
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

	if p.Place == LeftBottom || p.Place == LeftTop {
		text.Align = align.Left
	} else if p.Place == RightBottom || p.Place == RightTop {
		text.Align = align.Right
	}

	if p.Place == RightBottom || p.Place == Bottom || p.Place == LeftBottom {
		text.Top = height
	}

	text.BreakLineStrategy = breakline.EmptyLineStrategy

	return text
}

func (p *Page) GetPageString(current, total int) string {
	pattern := strings.ReplaceAll(p.Pattern, "{current}", fmt.Sprintf("%d", current))
	return strings.ReplaceAll(pattern, "{total}", fmt.Sprintf("%d", total))
}
