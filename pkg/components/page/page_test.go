package page_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := page.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*page.Page", fmt.Sprintf("%T", sut))
}
