package mappers

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type GenerateComponent func(document interface{}, sourceKey string) (Componentmapper, error)

type Componentmapper interface {
	Generate(content map[string]interface{}) (components.PdfComponent, error)
}

type AbstractFactoryMaps interface {
	NewRow(document interface{}, sourceKey string) (Componentmapper, error)
	NewPage(document interface{}, sourceKey string) (Componentmapper, error)
	NewList(document interface{}, sourceKey string, generate GenerateComponent) (Componentmapper, error)
}
