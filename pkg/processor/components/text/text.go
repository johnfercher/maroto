package text

import "github.com/johnfercher/maroto/v2/pkg/processor/components/props"

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
