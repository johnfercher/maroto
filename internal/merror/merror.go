package merror

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// DefaultErrorText is the default error text properties.
var DefaultErrorText = &props.Text{
	Family: fontfamily.Arial,
	Style:  fontstyle.Bold,
	Size:   10,
	Color: &props.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	},
}
