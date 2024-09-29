package factory

import (
	"html/template"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
)

type FactoryComponents struct{}

func NewFactoryComponents() *FactoryComponents {
	return &FactoryComponents{}
}

func (f *FactoryComponents) FactoryComponentTree(template template.Template, content map[string]interface{}) (core.Component, error) {
	return nil, nil
}
