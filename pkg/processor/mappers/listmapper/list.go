// listmapper is the package responsible for mapping row settings
package listmapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type List[T mappers.Componentmapper] struct {
	sourceKey string
	templates T
}

func NewList[T mappers.Componentmapper](document interface{}) (*List[T], error) {
	return nil, nil
}

func (r *List[T]) Generate(content map[string]interface{}) (components.Component, error) {
	return nil, nil
}
