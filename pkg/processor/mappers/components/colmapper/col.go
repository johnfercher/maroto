package colmapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type Col struct {
	Size       int
	Components mappers.Componentmapper
}

func NewCol(interface{}) (*Col, error) {
	return nil, nil
}

func (c *Col) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
