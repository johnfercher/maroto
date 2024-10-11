package text

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Text struct {
	Props props.TextProps
	Value string
}

func NewText(props props.TextProps, value string) *Text {
	return &Text{
		Props: props,
		Value: value,
	}
}

func (t *Text) Generate(provider processorprovider.ProcessorProvider) core.Component {
	return provider.CreateText(t.Value, t.Props)
}
