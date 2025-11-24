package entity

import "github.com/johnfercher/maroto/v2/pkg/props"

// SubText represents part of a text, this structure allows different properties in the same text
type SubText struct {
	Value string
	Props props.SubText
}
