package props

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Place is the representation of a place in a page.
type Place string

const (
	// LeftTop is the place in the left top of the page.
	LeftTop Place = "left_top"
	// Top is the place in the top of the page.
	Top Place = "top"
	// RightTop is the place in the right top of the page.
	RightTop Place = "right_top"
	// LeftBottom is the place in the left bottom of the page.
	LeftBottom Place = "left_bottom"
	// Bottom is the place in the bottom of the page.
	Bottom Place = "bottom"
	// RightBottom is the place in the right bottom of the page.
	RightBottom Place = "right_bottom"
)

// IsValid checks if the place is valid.
func (p Place) IsValid() bool {
	return p == LeftTop || p == Top || p == RightTop ||
		p == LeftBottom || p == Bottom || p == RightBottom
}

// Page is the representation of a page.
type Page struct {
	Pattern string
	Place   Place
	Family  string
	Style   fontstyle.Type
	Size    float64
	Color   *Color
}

// GetNumberTextProp returns the Text properties of the page number.
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

	text.BreakLineStrategy = breakline.EmptySpaceStrategy

	return text
}

// GetPageString returns the page string.
func (p *Page) GetPageString(current, total int) string {
	pattern := strings.ReplaceAll(p.Pattern, "{current}", fmt.Sprintf("%d", current))
	return strings.ReplaceAll(pattern, "{total}", fmt.Sprintf("%d", total))
}
