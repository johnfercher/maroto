package signature_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := signature.New("label")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*signature.signature", fmt.Sprintf("%T", sut))
}
