package entity

import "github.com/johnfercher/maroto/v2/pkg/consts/extension"

type Image struct {
	Bytes      []byte
	Extension  extension.Type
	Dimensions *Dimensions
}
