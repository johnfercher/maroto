package line_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := line.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*line.line", fmt.Sprintf("%T", sut))
}
