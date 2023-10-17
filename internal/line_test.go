package internal_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLine(t *testing.T) {
	// Act
	sut := internal.NewLine(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*internal.line", fmt.Sprintf("%T", sut))
}
