package entity

import "github.com/johnfercher/maroto/v2/pkg/consts/protection"

type Protection struct {
	Type          protection.Type
	UserPassword  string
	OwnerPassword string
}
