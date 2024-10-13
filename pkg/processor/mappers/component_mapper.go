package mappers

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type GenerateComponent func(document interface{}, sourceKey string) (Componentmapper, error)

type Componentmapper interface {
	Generate(content map[string]interface{}) (components.PdfComponent, error)
}

type AbstractFactoryMaps interface {
	NewRow(document interface{}, sourceKey string) (Componentmapper, error)
	NewPage(document interface{}, sourceKey string) (Componentmapper, error)
	NewCol(document interface{}) (Componentmapper, error)
	NewList(document interface{}, sourceKey string, generate GenerateComponent) (Componentmapper, error)
	NewBarcode(document interface{}) (Componentmapper, error)
	NewMatrixcode(document interface{}) (Componentmapper, error)
	NewQrcode(document interface{}) (Componentmapper, error)
	NewImage(document interface{}) (Componentmapper, error)
	NewLine(document interface{}) (Componentmapper, error)
	NewSignature(document interface{}) (Componentmapper, error)
	NewText(document interface{}) (Componentmapper, error)
}
