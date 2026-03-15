package entity_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
)

func TestProtection_AppendMap(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := fixtureProtection()
	m := make(map[string]any)

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, sut.Type, m["config_protection_type"])
	assert.Equal(t, sut.UserPassword, m["config_user_password"])
	assert.Equal(t, sut.OwnerPassword, m["config_owner_password"])
}

func fixtureProtection() entity.Protection {
	return entity.Protection{
		Type:          protection.Print,
		OwnerPassword: "123456",
		UserPassword:  "654321",
	}
}
