package mappers

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

// GenerateComponent defines the signature of a factory method, it is used
// to make it possible to send a factory method to another object
type GenerateComponent func(document interface{}, sourceKey string) (Componentmapper, error)

// The Component Mapper Interface defines the mapper component, the mapper component is responsible for
// transforming the structured document into the pdf components
type Componentmapper interface {
	Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error)
}

// The AbstractFactoryMaps interface defines the object responsible for wrapping the creation of components
// it is used to ensure decoupling between components
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
