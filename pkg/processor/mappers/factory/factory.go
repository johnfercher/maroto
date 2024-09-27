package factory

import (
	"html/template"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/content"
)

type FactoryComponents struct{}

func NewFactoryComponents() *FactoryComponents {
	return &FactoryComponents{}
}

func (f *FactoryComponents) FactoryComponentTree(template template.Template, content content.Content) (core.Component, error) {
	return nil, nil
}
