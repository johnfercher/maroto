package entity

import "github.com/johnfercher/maroto/v2/pkg/props"

// SubText represents part of a text, this structure allows different properties in the same text
type SubText struct {
	Value string
	Prop  props.SubText
}

func NewSubText(value string, ps ...props.SubText) *SubText {
	textProp := props.SubText{}
	if len(ps) > 0 {
		textProp = ps[0]
	}

	return &SubText{
		Value: value,
		Prop:  textProp,
	}
}
