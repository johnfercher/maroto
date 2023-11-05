package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaderConfig_MakeValid(t *testing.T) {
	t.Run("when start page is less than 0, should set to 0", func(t *testing.T) {
		// Arrange
		prop := HeaderConfig{
			StartPage: -1,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.StartPage, 0)
	})
}

func TestHeaderConfig_AppendMap(t *testing.T) {
	// Arrange
	sut := fixtureHeaderConfig()
	m := make(map[string]interface{})

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, 1, m["config_header_start_page"])
}

func fixtureHeaderConfig() HeaderConfig {
	return HeaderConfig{
		StartPage: 1,
	}
}
