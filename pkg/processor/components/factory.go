package components

import "github.com/johnfercher/maroto/v2/pkg/processor/core"

type FactoryComponents struct{}

func NewFactoryComponents() *FactoryComponents {
	return &FactoryComponents{}
}

func (f *FactoryComponents) FactoryComponentTree(template interface{}, content map[string]interface{}) (core.Component, error) {
	return nil, nil
}
