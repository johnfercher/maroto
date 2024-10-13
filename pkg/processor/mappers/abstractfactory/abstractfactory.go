package abstractfactory

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/colmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/rowmapper"
)

// abstractFactoryMaps is responsible for providing a factory for all mapper components
type abstractFactoryMaps struct{}

func NewAbstractFactoryMaps() *abstractFactoryMaps {
	return &abstractFactoryMaps{}
}

// NewRow is responsible for wrapper the creation of a row
func (f *abstractFactoryMaps) NewRow(document interface{}, sourceKey string) (mappers.Componentmapper, error) {
	return rowmapper.NewRow(document, sourceKey, f)
}

// NewPage is responsible for wrapper the creation of a page
func (f *abstractFactoryMaps) NewPage(document interface{}, sourceKey string) (mappers.Componentmapper, error) {
	return pagemapper.NewPage(document, sourceKey, f)
}

// NewCol is responsible for wrapper the creation of a col
func (f *abstractFactoryMaps) NewCol(document interface{}) (mappers.Componentmapper, error) {
	return colmapper.NewCol(document)
}

// NewList is responsible for wrapper the creation of a list
func (f *abstractFactoryMaps) NewList(document interface{}, sourceKey string, generate mappers.GenerateComponent) (mappers.Componentmapper, error) {
	return listmapper.NewList(document, sourceKey, generate)
}
