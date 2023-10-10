package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtection_AppendMap(t *testing.T) {
	// Arrange
	sut := fixtureProtection()
	m := make(map[string]interface{})

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, sut.Type, m["config_protection_type"])
	assert.Equal(t, sut.UserPassword, m["config_user_password"])
	assert.Equal(t, sut.OwnerPassword, m["config_owner_password"])
}

func fixtureProtection() Protection {
	return Protection{
		Type:          protection.Print,
		OwnerPassword: "123456",
		UserPassword:  "654321",
	}
}
