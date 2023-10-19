package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewBorderLineStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewBorderLineStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderLineStyler", fmt.Sprintf("%T", sut))
}
