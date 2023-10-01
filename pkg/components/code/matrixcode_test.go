package code_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	// Act
	sut := code.NewMatrix("code")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*code.matrixCode", fmt.Sprintf("%T", sut))
}
