package components

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type FactoryComponents struct{}

func NewFactoryComponents() *FactoryComponents {
	return &FactoryComponents{}
}

func (f *FactoryComponents) FactoryComponentTree(template mappers.Template, content mappers.Content) (core.Component, error) {
	return nil, nil
}
