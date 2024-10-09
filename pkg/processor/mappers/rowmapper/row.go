// rowmapper is the package responsible for mapping row settings
package rowmapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Row struct {
	Test string
}

func NewRow(document interface{}, sourceKey string) (*Row, error) {
	return &Row{Test: "a"}, nil
}

func (r *Row) Generate(content map[string]interface{}) (components.Component, error) {
	return nil, nil
}
