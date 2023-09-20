package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/line"
)

type Cell struct {
	BackgroundColor *Color
	BorderColor     *Color
	BorderType      border.Type
	BorderThickness float64
	LineStyle       line.Style
}
