package core_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestNewPDF(t *testing.T) {
	// Act
	sut := core.NewPDF(nil, nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*core.pdf", fmt.Sprintf("%T", sut))
}
