package row_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := row.New(10)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*row.row", fmt.Sprintf("%T", sut))
}
