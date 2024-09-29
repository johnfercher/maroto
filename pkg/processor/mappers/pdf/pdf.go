package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/page"
)

type Pdf struct {
	Builder builder.Builder      `json:"builder"`
	Pages   map[string]page.Page `json:"pages"`
}
