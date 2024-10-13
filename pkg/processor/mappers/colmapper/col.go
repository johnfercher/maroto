package colmapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Col struct{}

func NewCol(interface{}) (*Col, error) {
	return nil, nil
}

func (c *Col) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
