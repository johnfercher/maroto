package mappers

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Componentmapper interface {
	Generate(content map[string]interface{}) (components.Component, error)
}
