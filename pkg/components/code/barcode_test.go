package code_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBar(t *testing.T) {
	// Act
	sut := code.NewBar("code")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*code.barcode", fmt.Sprintf("%T", sut))
}
