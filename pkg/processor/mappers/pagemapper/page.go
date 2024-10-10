// pagemapper is the package responsible for mapping page settings
package pagemapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
)

type Page struct {
	Teste string
}

func NewPage(page interface{}, sourceKey string) (*Page, error) {
	return nil, nil
}

func (r *Page) Generate(content map[string]interface{}) (components.Component, error) {
	return nil, nil
}
