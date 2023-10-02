package maroto_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := maroto.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*maroto.maroto", fmt.Sprintf("%T", sut))
}
