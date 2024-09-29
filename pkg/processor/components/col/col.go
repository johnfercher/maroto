package col

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
)

type Col struct {
	Props      props.ColProps
	Components components.Component
}

func NewCol(props props.ColProps, components ...components.Component) *Col {
	return &Col{
		Props:      props,
		Components: components,
	}
}
