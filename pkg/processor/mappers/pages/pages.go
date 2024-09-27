package pages

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/foreach"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/rows"
)

type Page struct {
	ForEach foreach.ForEach[rows.Row]
}
