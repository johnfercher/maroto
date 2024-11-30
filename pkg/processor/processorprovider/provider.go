package processorprovider

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type ProviderComponent interface {
	core.Node
}

// ProcessorProvider provides an interface with all the methods that
// Maroto provides for pdf builder
type ProcessorProvider interface {
	Generate() (core.Document, error)
	GetStructure() *node.Node[core.Structure]
	AddPages(pages ...ProviderComponent) (ProcessorProvider, error)
	AddFooter(footer ...ProviderComponent) (ProcessorProvider, error)
	AddHeader(header ...ProviderComponent) (ProcessorProvider, error)
	CreatePage(components ...ProviderComponent) (ProviderComponent, error)
	CreateRow(height float64, components ...ProviderComponent) (ProviderComponent, error)
	CreateCol(size int, components ...ProviderComponent) (ProviderComponent, error)
	CreateText(value string, props ...*propsmapper.Text) ProviderComponent
	CreateSignature(value string, props ...*propsmapper.Signature) ProviderComponent
	CreateBarCode(value string, props ...*propsmapper.Barcode) ProviderComponent
	CreateMatrixCode(value string, props ...*propsmapper.Rect) ProviderComponent
	CreateQrCode(value string, props ...*propsmapper.Rect) ProviderComponent
	CreateImage(value []byte, extension string, props ...*propsmapper.Rect) ProviderComponent
	CreateLine(props ...*propsmapper.Line) ProviderComponent
}
