package col

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Col struct {
	Props      props.ColProps
	Components []components.PdfComponent
}

func NewCol(props props.ColProps, components ...components.PdfComponent) *Col {
	return &Col{
		Props:      props,
		Components: components,
	}
}

func (c *Col) Generate(provider processorprovider.ProcessorProvider) core.Col {
	components := make([]core.Component, len(c.Components))

	for i, component := range c.Components {
		components[i] = component.Generate(provider)
	}
	return provider.CreateCol(c.Props.Size, components...)
}
