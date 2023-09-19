package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
)

type Cell struct {
	BackgroundColor *Color
	BorderColor     *Color
	Border          border.Type
}
