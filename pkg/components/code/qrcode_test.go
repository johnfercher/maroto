package code_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/stretchr/testify/assert"
)

func TestNewQr(t *testing.T) {
	// Act
	sut := code.NewQr("code")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*code.qrCode", fmt.Sprintf("%T", sut))
}
