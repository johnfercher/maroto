package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
)

// Protection is the representation of a pdf protection.
type Protection struct {
	Type          protection.Type
	UserPassword  string
	OwnerPassword string
}

// AppendMap adds the Protection fields to the map.
func (p *Protection) AppendMap(m map[string]interface{}) map[string]interface{} {
	if p.Type != 0 {
		m["config_protection_type"] = p.Type
	}

	if p.UserPassword != "" {
		m["config_user_password"] = p.UserPassword
	}

	if p.OwnerPassword != "" {
		m["config_owner_password"] = p.OwnerPassword
	}

	return m
}
