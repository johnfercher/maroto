// listmapper is the package responsible for mapping row settings
package listmapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type List struct {
	sourceKey string
	templates mappers.Componentmapper
}

func NewList(document interface{}, sourceKey string, generate mappers.GenerateComponent) (*List, error) {
	return &List{}, nil
}

func (r *List) Generate(content map[string]interface{}) (components.Component, error) {
	return nil, nil
}
