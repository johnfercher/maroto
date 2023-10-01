package text_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// Act
	sut := text.New("label")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*text.text", fmt.Sprintf("%T", sut))
}
