package abstractfactory

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/colmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/imagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/linemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/signaturemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/textmapper"
)

type abstractFactoryMaps struct {
	repository core.ProcessorRepository
}

// NewAbstractFactoryMaps is responsible for creating an object that encapsulates the creation of components
func NewAbstractFactoryMaps(repository core.ProcessorRepository) *abstractFactoryMaps {
	return &abstractFactoryMaps{repository: repository}
}

// NewRow is responsible for wrapper the creation of a row
func (f *abstractFactoryMaps) NewRow(document interface{}, sourceKey string) (mappers.OrderedComponents, error) {
	return rowmapper.NewRow(document, sourceKey, f)
}

// NewPage is responsible for wrapper the creation of a page
func (f *abstractFactoryMaps) NewPage(document interface{}, sourceKey string) (mappers.OrderedComponents, error) {
	return pagemapper.NewPage(document, sourceKey, f)
}

// NewCol is responsible for wrapper the creation of a col
func (f *abstractFactoryMaps) NewCol(document interface{}) (mappers.Componentmapper, error) {
	return colmapper.NewCol(document, f)
}

// NewList is responsible for wrapper the creation of a list
func (f *abstractFactoryMaps) NewList(document interface{}, sourceKey string,
	generate mappers.GenerateComponent,
) (mappers.OrderedComponents, error) {
	return listmapper.NewList(document, sourceKey, generate)
}

// NewBarcode is responsible for wrapper the creation of a barcode
func (f *abstractFactoryMaps) NewBarcode(document interface{}) (mappers.OrderedComponents, error) {
	return codemapper.NewBarcode(document)
}

// NewMatrixcode is responsible for wrapper the creation of a matrix code
func (f *abstractFactoryMaps) NewMatrixcode(document interface{}) (mappers.OrderedComponents, error) {
	return codemapper.NewMatrixcode(document)
}

// NewQrcode is responsible for wrapper the creation of a qrcode
func (f *abstractFactoryMaps) NewQrcode(document interface{}) (mappers.OrderedComponents, error) {
	return codemapper.NewQrcode(document)
}

// NewImage is responsible for wrapper the creation of a image
func (f *abstractFactoryMaps) NewImage(document interface{}) (mappers.OrderedComponents, error) {
	return imagemapper.NewImage(document)
}

// NewLine is responsible for wrapper the creation of a libe
func (f *abstractFactoryMaps) NewLine(document interface{}) (mappers.OrderedComponents, error) {
	return linemapper.NewLine(document)
}

// NewSignature is responsible for wrapper the creation of a signature
func (f *abstractFactoryMaps) NewSignature(document interface{}) (mappers.OrderedComponents, error) {
	return signaturemapper.NewSignature(document)
}

// NewText is responsible for wrapper the creation of a text
func (f *abstractFactoryMaps) NewText(document interface{}) (mappers.OrderedComponents, error) {
	return textmapper.NewText(document)
}
